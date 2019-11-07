package gredis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Lpush(key string, data interface{}, time int) error{
	conn := RedisConn.Get()
	defer conn.Close()
	v, err := json.Marshal(data); if err != nil {
		return err
	}
	fmt.Println("v:",v)
	_, err = conn.Do("lpush",key,v);if err != nil{
		return err
	}
	_, err = conn.Do("expire",key,time);if err != nil {
		return err
	}
	return nil
}
func Lrange(key string, data interface{}){}

func Lpop(key string) (string,error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("lpop",key));if err != nil{
		return "",err
	}

	return reply,nil
}

func Len(key string) (int, error){
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("llen",key));if err != nil {
		return 0,err
	}

	return reply, nil
}
//---------------------------------------------------------------------------------