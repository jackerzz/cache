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




func main() {
	gredis.Set("jacker","arrerr",3600)
	r1,err :=gredis.Get("jacker"); if err != nil {
		fmt.Errorf("errr:",err)
	}
	fmt.Println(r1)

	gredis.Lpush("query","paker",36666)

	r2, err := gredis.Len("query"); if err !=nil {
		fmt.Errorf("eer",err)
	}
	fmt.Println("queue len",r2)

	r3,err := gredis.Lpop("query");if err != nil {
		fmt.Errorf("err",err)
	}
	fmt.Println("reply",r3)

}
