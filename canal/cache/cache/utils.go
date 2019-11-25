package cache

import (
	"cache/canal/cache/setting"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"time"
)



// Set a key/value
func Set(key string, data interface{}, expiration time.Duration) error {
	conn := RedisConn.Get()
	defer conn.Close()

	fmt.Println(data)
	value, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, expiration)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}


// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

//================================================================================================================

func Mget(){
	conn := RedisConn.Get()
	defer conn.Close()
}

func Hmset(){
	conn := RedisConn.Get()
	defer conn.Close()
}

//================================================================================================================

// Count key
func Incr(key string)(int, error){
	conn := RedisConn.Get()
	defer conn.Close()

	incr, err := redis.Int(conn.Do("INCR",key))
	if err != nil{
		return 0,err
	}
	return incr,nil
}

// Get key remaining time
func Ttl(key string) (int,error){
	conn := RedisConn.Get()
	defer conn.Close()

	times, err := redis.Int(conn.Do("ttl", key))
	if err != nil {
		return 0, err
	}
	return times, nil

}

//================================================================================================================
var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}


func init(){
	Setup()
}

