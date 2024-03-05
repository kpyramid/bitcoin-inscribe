package types

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
)

type Manager struct {
	singleNodeClient *redis.Client
	clusterClient    *redis.ClusterClient
	Client           redis.Cmdable
}

var (
	ManagerClient *Manager
)

func GetClient() redis.Cmdable {
	return ManagerClient.Client
}

func InitRedis(cnf *Config) error {
	m := &Manager{}
	if cnf.RedisIsClusterMode {
		// 集群模式
		opt := &redis.ClusterOptions{
			Addrs:    strings.Split(cnf.RedisClusterAddresses, ","),
			Password: cnf.RedisPassword,
		}
		if cnf.TlsInsecureSkip {
			opt.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		cli := redis.NewClusterClient(opt)
		_, err := cli.Ping(context.Background()).Result()
		if err != nil {
			return err
		}
		m.clusterClient = cli
		m.Client = cli

	} else {

		// 单节点模式
		opt := &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cnf.RedisHost, cnf.RedisPort),
			Password: cnf.RedisPassword,
			DB:       cnf.RedisDB,
		}
		if cnf.TlsInsecureSkip {
			opt.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		cli := redis.NewClient(opt)
		_, err := cli.Ping(context.Background()).Result()
		if err != nil {
			return err
		}
		m.singleNodeClient = cli
		m.Client = cli
	}

	ManagerClient = m
	return nil
}
