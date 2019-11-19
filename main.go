package main

import (
	"cache/cache/gredis"
	"cache/cache/setting"
	"fmt"
)



func init(){
	setting.Setup()
	gredis.Setup()
}


func main() {
	test1()
}

func test1(){
	gredis.Set("jacker","arrerr",3600)
	r1,err :=gredis.Get("jacker"); if err != nil {
		fmt.Println("r1",err)
	}
	fmt.Println(r1)
	r5, err := gredis.Ttl("jacker")
	fmt.Println(r5,err)

	gredis.Lpush("query","jac",36666)

	r2, err := gredis.Len("query"); if err !=nil {
		fmt.Println("r2",err)
	}
	fmt.Println("r2",r2)

	reply,err := gredis.Lrange("query"); if err != nil {
		fmt.Println("reply",err)
	}
	fmt.Println(reply)


}

func test2(){
	_,err:=gredis.Brpop("list:1",0)
	fmt.Println(err)
	gredis.Lpush("list:1","hahahh",20)

}

func test3(){
	// 阻塞状态下
	r3,err := gredis.Brpop("list1",0);if err != nil {
		fmt.Println("r3",err)
	}
	fmt.Println("r3",r3)
}

func test4(){
	err :=gredis.SetWatch("jacker","jjjjjjsafsda",43);if err != nil {
		fmt.Println(err)
	}
}
func test5(){
	reply, err:=gredis.Lrange("data:student:1");if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}