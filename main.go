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
	data,err :=gredis.Get("jacker"); if err != nil {
		fmt.Errorf("errr:",err)
	}
	fmt.Println(data)

}
