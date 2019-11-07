package gredis

import (
	"github.com/gomodule/redigo/redis"
)


func Lpush(key string, data interface{}, time int) error{
	conn := RedisConn.Get()
	defer conn.Close()
	//v, err := json.Marshal(data); if err != nil {
	//	return err
	//}
	_, err := conn.Do("lpush",key,data);if err != nil{
		return err
	}
	_, err = conn.Do("expire",key,time);if err != nil {
		return err
	}
	return nil
}

func Brpop(key string,timeout int) (interface{},error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("brpop",key,timeout));if err != nil{
		return nil,err
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

// get all list info
func Lrange(key string)([]interface{}, error){
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err  := redis.Values(conn.Do("lrange",key,0,-1));if err != nil{
		return nil, err
	}
	return reply, nil
}
