package main

import (
	"cache/gredis"
	"cache/setting"
)


func init(){
	setting.Setup()
	gredis.Setup()
}

func main(){
	test1()
}
// service test3 消息队列
func test1(){
	gredis.Lpush("list1","hhhahhhs",23)
}
