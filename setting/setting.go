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
type Canal struct {
	Address           string
	Port              int
	UserName          string
	PassWord          string
	SoTime            int32
	IdleTimeOut       int32
	Connected         bool
	Running           bool
	Filter            string
	RollbackOnConnect bool
	LazyParseEntry    bool
}
var DatabaseSetting = &Database{}
var RedisSetting = &Redis{}
var CanalSetting = &Canal{}

var cfg  *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("src/cache/conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'src/cache/conf/app.ini': %v", err)
	}
	mapTo("redis", RedisSetting)
	mapTo("database", DatabaseSetting)
	mapTo("canal",CanalSetting)
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}