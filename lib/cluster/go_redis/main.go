package main


import (
	"fmt"
	"github.com/chasex/redis-go-cluster"
	"log"
	"time"
)
func main(){
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes: []string{"127.17.0.4:6381","127.17.0.3:6380"},
			ConnTimeout: 50 * time.Microsecond,
			ReadTimeout: 50 * time.Microsecond,
			WriteTimeout: 50 * time.Microsecond,
			KeepAlive: 16,
			AliveTime:60 * time.Microsecond,
		})
	if err !=nil {
		log.Fatal("redis new error: %s",err.Error())
	}
	_, err = cluster.Do("set", "{user000}.name", "Joel")
	_, err = cluster.Do("set", "{user000}.age", "26")
	_, err = cluster.Do("set", "{user000}.country", "China")

	name, err := redis.String(cluster.Do("get", "{user000}.name"))
	if err != nil {
		log.Fatal(err)
	}
	age, err := redis.Int(cluster.Do("get", "{user000}.age"))
	if err != nil {
		log.Fatal(err)
	}
	country, err := redis.String(cluster.Do("get", "{user000}.country"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s, age: %d, country: %s\n", name, age, country)

	cluster.Close()
	_, err = cluster.Do("set", "foo", "bar")
	if err == nil {
		log.Fatal("expect a none nil error")
	}
	log.Println(err)

}