package main

import (
	"cache/gredis"
	"cache/setting"
	"fmt"
)


func init(){
	setting.Setup()
	gredis.Setup()
}

func main(){
	//test1()
	test2()
}
// service test3 消息队列
func test1(){
	gredis.Lpush("list1","hhhahhhs",23)
}
func test2(){
	err :=gredis.SetWatch("jackers","jjjjjjsafsda",43);if err != nil {
		fmt.Println(err)
	}
}