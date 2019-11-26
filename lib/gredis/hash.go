package gredis

import "github.com/gomodule/redigo/redis"

// hash

func Hset(key string, data interface{}) error{
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := redis.Bytes(conn.Do("hset",key,data)); if err != nil {
		return err
	}
	return nil
}

func Hget(key string)(string,error){
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("hget",key)); if err != nil {
		return "",err
	}
	return reply,nil
}
//---------------------------------------------------------------------------------
func Hmset()(){}
func Hmget()(){}

// --------------------------------------------------------------------------------
