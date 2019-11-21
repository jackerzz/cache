package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/memberlist"
	"os"
	"stathat.com/c/consistent"
	"strconv"
	"time"
)

var bindPort = flag.Int("port", 8001, "gossip port")

func main()  {
	flag.Parse()
	hostname, _ := os.Hostname()
	config := memberlist.DefaultLocalConfig()  // 返回一个 struct

	config.Name = hostname + "-" + strconv.Itoa(*bindPort)
	config.BindPort = *bindPort
	config.AdvertisePort = *bindPort

	fmt.Println("config.DisableTcpPings -->禁用TCP Ping：", config.DisableTcpPings)
	fmt.Println("config.IndirectChecks -->间接检查", config.IndirectChecks)
	fmt.Println("config.RetransmitMult -->重发多个", config.RetransmitMult)
	fmt.Println("config.PushPullInterval -->推拉间隔", config.PushPullInterval)
	fmt.Println("config.ProbeInterval--> 探测间隔", config.ProbeInterval)
	fmt.Println("config.GossipInterval->发送需要的消息之间的间隔", config.GossipInterval)
	fmt.Println("config.GossipNodes-->向其发送gossip消息的随机节点数 ", config.GossipNodes)
	fmt.Println("config.BindPort--> 绑定端口", config.BindPort)

	fmt.Println("绑定后的config info:",config)
	list, _ := memberlist.Create(config)
	list.Join([]string{"127.0.0.1:8001", "127.0.0.1:8002"})
	fmt.Println(list.Members())

	cons := consistent.New()
	cons.NumberOfReplicas = 256
	go func(){
		for {
			m := list.Members()
			nodes := make([]string, len(m))
			for i, n := range m{
				nodes[i] = n.Name
			}
			cons.Set(nodes)
			time.Sleep(time.Second)
		}
	}()

	for {
		fmt.Println("---------------start----------------")
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
		}
		fmt.Println("---------------end----------------")
		time.Sleep(time.Second * 3)
	}

}