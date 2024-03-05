package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/pkg/btcapi/mempool"
	"github.com/kpyramid/bitcoin-inscribe/types/service"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"sync"
	"time"
)

var _service *ServiceContext = nil
var once sync.Once

const FeeRateCacheKey = "btc_fee_rate_estimate"
const FeeRateCacheKeyTTL = time.Minute * 5

type ServiceContext struct {
	NetParams    *chaincfg.Params
	Client       *rpcclient.Client
	BtcApiClient *mempool.MempoolClient
	UnisatClient *service.UnisatClient
	Config       *Config
	Wallet       *HDWallet
	Redis        redis.Cmdable
	Db           *gorm.DB
	QuitMutex    *sync.Mutex
}

func GetServiceContext() *ServiceContext {
	once.Do(func() {
		cfg := GetConfig()
		netParams := &chaincfg.TestNet3Params
		if cfg.Network == "mainnet" {
			netParams = &chaincfg.MainNetParams
		}

		getUrl, err := url.Parse(cfg.RpcHost)
		if err != nil {
			log.Fatal(err)
		}

		isTls := true
		host := fmt.Sprintf("%s%s", getUrl.Host, getUrl.Path)
		if getUrl.Scheme == "https" {
			isTls = false
		}
		client, err := rpcclient.New(&rpcclient.ConnConfig{
			Host:         host,
			User:         cfg.RpcUser,
			Pass:         cfg.RpcPass,
			HTTPPostMode: true,
			DisableTLS:   isTls,
		}, nil)
		if err != nil {
			log.Fatal(err)
		}

		// from order wallet
		wallet := &HDWallet{}
		if err := wallet.Init(netParams, cfg.OrderWalletMnemonic); err != nil {
			log.Fatal(err)
		}

		// DB
		Db, err := gorm.Open(mysql.Open(cfg.MySQLDataSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		sqlDB, err := Db.DB()
		if err != nil {
			panic(err)
		}
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.MySQLConnMaxLifetime) * time.Second)

		// redis
		if err := InitRedis(cfg); err != nil {
			log.Fatal(err)
		}

		// unisat
		unisatClient := service.NewUnisatClient(netParams, cfg.UnisatApiKey)
		_service = &ServiceContext{
			NetParams:    netParams,
			Client:       client,
			BtcApiClient: mempool.NewClient(cfg.MempoolAddress, netParams),
			Wallet:       wallet,
			Config:       cfg,
			Redis:        GetClient(),
			UnisatClient: unisatClient,
			Db:           Db,
			QuitMutex:    &sync.Mutex{},
		}
	})

	return _service
}

func (s *ServiceContext) GetEstimateFeeRate() (*mempool.FeeRateEstimate, error) {
	result, err := s.Redis.Get(context.TODO(), FeeRateCacheKey).Result()
	if err != nil {
		// if key not found
		if err == redis.Nil {
			// get estimate from mempool
			feeRateEstimate, err := s.BtcApiClient.GetFeeRateEstimate()
			if err != nil {
				return nil, err
			}
			feeRateEstimateBz, err := json.Marshal(feeRateEstimate)
			if err != nil {
				return nil, err
			}
			if _, err := s.Redis.Set(context.TODO(), FeeRateCacheKey, feeRateEstimateBz, FeeRateCacheKeyTTL).Result(); err != nil {
				return nil, err
			}
			return feeRateEstimate, nil
		}
		return nil, err
	}

	// parse from cache
	feeRateEstimate := mempool.FeeRateEstimate{}
	if err := json.Unmarshal([]byte(result), &feeRateEstimate); err != nil {
		return nil, err
	}
	return &feeRateEstimate, nil
}
