package config

import (
	"flag"
	"fmt"
	"gmimo/common/constant"
	"gmimo/common/log"
	"gmimo/common/util"
	"gopkg.in/ini.v1"
	"time"
)

var (
	filename string
	cfg      *ini.File

	// 全局变量
	AppFullName string
	// 配置项
	App    = ConfApp{}
	Rpc    = ConfRpc{}
	Logger = log.LogConfig{}
	Redis  = ConfRedis{}
)

type ConfApp struct {
	Name         string
	Addr         string
	Debug        bool
	JwtSecret    string
	JwtExpire    time.Duration
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

func InitConfig() (err error) {
	flag.StringVar(&filename, "config", "./config/config.ini", "default config file")
	cfg, err = ini.Load(filename)
	if err != nil {
		return
	}

	// load app Section
	if err = cfg.Section("app").MapTo(&App); err != nil {
		return
	}
	AppFullName = util.GetAppFullName(App.Name)
	App.JwtExpire = App.JwtExpire * time.Hour
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

	return nil

}

func String() string {
	return fmt.Sprintf("App:%+v, Rpc:%+v, Logger:%+v, Redis:%+v", App, Rpc, Logger, Redis)
}
