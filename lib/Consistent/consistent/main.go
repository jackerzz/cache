package consistent

import (
	"fmt"
	"stathat.com/c/consistent"
)

func main() {
	cons := consistent.New()
	cons.Add("cacheA")
	cons.Add("cacheB")
	cons.Add("cacheC")

	server1, err := cons.Get("user_1")
	server2, err := cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC
	fmt.Println("server2:", server2) //输出 server2: cacheA

	fmt.Println()

	//user_1在cacheA上，把cacheA删掉后看下效果
	cons.Remove("cacheA")
	server1, err = cons.Get("user_1")
	server2, err = cons.Get("user_2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server1:", server1) //输出 server1: cacheC,和删除之前一样，在同一个server上
	fmt.Println("server2:", server2) //输出 server2: cacheB,换到另一个server了
}