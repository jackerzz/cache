package main

import (
	"fmt"
	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"log"
	"os"
	"time"
)
var (
	timeout 	int64
	units 		int32
)
func main(){
	timeout = 123
	units 	= 34
}
// 链接canal 获取增量数据
func connectGetData()  {
	connector := client.NewSimpleCanalConnector("192.168.225.129", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect();if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = connector.Subscribe(".*\\\\..*");if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {
		message, err := connector.Get(10,&timeout,&units);if err !=nil{
			log.Println("connector get err:",err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(3000 * time.Microsecond)
			fmt.Println("=========没有数据=================")
			continue
		}
		filerGetData(message.Entries)
	}

}
// 根据表筛选并处理增量数据
func filerGetData(entrys []protocol.Entry){
	for _,entrys := range entrys{
		if entrys.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN ||
			entrys.GetEntryType() == protocol.EntryType_TRANSACTIONEND{
				continue
		}
		rowChange := new(protocol.RowChange)
	}
}
// 依据策略处理增量数据
func ProcessingIncrementalData() {}
//更新缓存
func setCache(){}
//删除缓存
func delCache(){}
//格式转换
func dataFormate(){}