package main

import (
	"cache/gredis"
	"cache/setting"
	"fmt"
	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"

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
	connectGetData()
}

// 链接canal 获取增量数据
func connectGetData()  {
	connector := client.NewSimpleCanalConnector("192.168.225.129", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect();if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = connector.Subscribe(setting.SubscribesSetting.Subscribe);if err != nil {
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
			continue
		}
		filerGetData(message.Entries)
	}

}


// 根据表筛选并处理增量数据
func filerGetData(entrys []protocol.Entry){
	for _,entry := range entrys{
		//如果是事务开始或、事务结束或者查询，则跳过本次循环
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN ||
			entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND{
				continue
		}
		//定义数据变化并序列化
		rowChange := new(protocol.RowChange)
		err := proto.Unmarshal(entry.GetStoreValue(),rowChange)
		log.Println("proto.Unmarshal err:",err)
		// 获取事件类型
		eventype := rowChange.GetEventType()
		// 检测 dbname 是否为需要写入redis的库
		processingIncrementalData(entry,eventype,rowChange)

	}
}



// 依据策略处理增量数据
func processingIncrementalData(Entry protocol.Entry, eventType protocol.EventType,Rowchange *protocol.RowChange ){
	header := Entry.GetHeader()
	fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))
	for _,rowdata := range Rowchange.GetRowDatas(){
		if eventType == protocol.EventType_INSERT {
			delCache(Entry.GetHeader(),rowdata.GetAfterColumns())
		}else if eventType == protocol.EventType_DELETE {
			setCache(Entry.GetHeader(),rowdata.GetBeforeColumns())
		}else {

		}
	}
}
//更新缓存
func setCache(Header *protocol.Header,columns []*protocol.Column){
	for _, col := range columns{
		value := col.GetName()+":"+col.GetValue()
		if col.IsKey{
			key := Header.GetSchemaName()+":"+Header.GetTableName()+":"+col.GetName()
			gredis.Set(key,value,600000)
		}

	}
}
//删除缓存
func delCache(Header *protocol.Header,columns []*protocol.Column){}
//组装
