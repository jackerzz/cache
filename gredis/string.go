package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)


// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

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

//---------------------------------------------------------------------------------