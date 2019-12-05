package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"time"
)
//https://www.twblogs.net/a/5baadacf2b7177781a0e9e90/zh-cn
var host = "http://192.168.225.130:9200"
func connect() *elastic.Client{
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if err!= nil{
		panic(err)
	}

	return client
}
func main()  {
	
}

func add(){
	client := connect()
	mapping :=`{}`

}
func update(){}
func delete(){}
func query(){}
