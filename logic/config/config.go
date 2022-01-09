package config

import (
	"flag"
	"fmt"
	"gmimo/common/constant"
	"gmimo/common/log"
	"gopkg.in/ini.v1"
	"time"
)

var (
	filename string
	cfg      *ini.File
	// 配置项
	App    = ConfApp{}
	Rpc    = ConfRpc{}
	Logger = log.LogConfig{}
	Redis  = ConfRedis{}
	Kafka  = ConfKafka{}
)

type ConfApp struct {
	Name  string
	Addr  string
	Debug bool

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxKB        int
}

type ConfRpc struct {
	CometSrv  string
	CometPort int
	LogicSrv  string
	LogicPort int
	PushSrv   string
	PushPort  int
	Timeout   time.Duration
}

type ConfRedis struct {
	HostPort    string
	Password    string
	Db          int
	PoolSize    int
	MaxRetries  int
	IdleTimeout time.Duration //秒
}

type ConfKafka struct {
	Brokers         []string
	Topic           string
	Partitions      int32
	MaxProduceRetry int
	ReturnSuccesses bool
}

func InitConfig() (err error) {
	//flag.StringVar(&Filename, "config", "./config/config.ini", "default config file")
	flag.StringVar(&filename, "config", "C:/Data/GoProject/gmimo/logic/config/config.ini", "default config file")
	cfg, err = ini.Load(filename)
	if err != nil {
		return
	}

	// load app Section
	if err = cfg.Section("app").MapTo(&App); err != nil {
		return
	}
	App.ReadTimeout = App.ReadTimeout * time.Millisecond
	App.WriteTimeout = App.WriteTimeout * time.Millisecond
	App.MaxKB = App.MaxKB * constant.KB

	// load logger Section
	if err = cfg.Section("logger").MapTo(&Logger); err != nil {
		return
	}

	// load rpc Section
	if err = cfg.Section("rpc").MapTo(&Rpc); err != nil {
		return
	}
	Rpc.Timeout = Rpc.Timeout * time.Millisecond

	// load redis Section
	if err = cfg.Section("redis").MapTo(&Redis); err != nil {
		return
	}
	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

	// load kafka Section
	if err = cfg.Section("kafka").MapTo(&Kafka); err != nil {
		return
	}

	return nil

}

func String() string {
	return fmt.Sprintf("App:%+v, Rpc:%+v Logger:%+v, Redis:%+v, Kafka:%+v", App, Rpc, Logger, Redis, Kafka)
}
