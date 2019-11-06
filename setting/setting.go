package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var DatabaseSetting = &Database{}
var RedisSetting = &Redis{}

var cfg  *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("cache/conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'cache/conf/app.ini': %v", err)
	}
	mapTo("redis", RedisSetting)
	mapTo("database", DatabaseSetting)
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}