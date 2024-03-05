package types

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Address               string `env:"ADDRESS"  envDefault:"0.0.0.0" json:"address,omitempty"`
	GrpcPort              int64  `env:"GRPC_PORT"  envDefault:"9000" json:"grpc_port,omitempty"`
	HttpPort              int64  `env:"HTTP_PORT"  envDefault:"80" json:"http_port,omitempty"`
	MySQLDataSource       string `env:"MYSQL_DATA_SOURCE" envDefault:""`
	MySQLMaxIdleConns     int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns     int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime  int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RedisHost             string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort             string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword         string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB               int    `env:"REDIS_DB" envDefault:"0"`
	TlsInsecureSkip       bool   `env:"TLS_INSECURE_SKIP" envDefault:"false"`
	RedisIsClusterMode    bool   `env:"REDIS_IS_CLUSTER_MODE" envDefault:"false"`
	RedisClusterAddresses string `env:"REDIS_CLUSTER_ADDRESSES" envDefault:""`
	UseCoroutine          bool   `env:"USE_COROUTINE" envDefault:"true"`
	Network               string `env:"NETWORK"  envDefault:"testnet" json:"network,omitempty"`
	RpcHost               string `env:"RPC_HOST"  envDefault:"" json:"rpc_host,omitempty"`
	RpcUser               string `env:"RPC_USER"  envDefault:"" json:"rpc_user,omitempty"`
	RpcPass               string `env:"RPC_PASS"  envDefault:"" json:"rpc_pass,omitempty"`
	MinFeeRate            int64  `env:"MIN_FEE_RATE" envDefault:"1"`
	OrderWalletMnemonic   string `env:"ORDER_WALLET_MNEMONIC"  envDefault:"" json:"order_wallet_mnemonic,omitempty"`
	MempoolAddress        string `env:"MEMPOOL_ADDRESS" envDefault:"https://blockstream.info/testnet/api" json:"mempool_address"`
	UnisatApiKey          string `env:"UNISAT_API_KEY" envDefault:"" json:"unisat_api_key"`
	MinimumReceiptAmount  int64  `env:"MINIMUM_RECEIPT_AMOUNT" envDefault:"250000" json:"minimum_receipt_amount"`
}

func GetConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.WithField("error", err).Panic("load config failed")
		return nil
	}

	return cfg
}
