package core

import (
	"cache/canal/cache/cache"
	"fmt"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"
	"os"
)

func PrintEntry(entrys []protocol.Entry) {

	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			//fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE{
					// 删除键
					cache.DelRedis(header,rowData.GetBeforeColumns())
				} else if eventType == protocol.EventType_INSERT {
					// 插入
					cache.SetRedis(header,rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					orther(header,rowData.GetBeforeColumns())
					cache.SetRedis(header,rowData.GetAfterColumns())
					//fmt.Println("-------> after")
					//orther(header,rowData.GetAfterColumns())
					//cache.SetRedis(header,rowData.GetAfterColumns())
				}
			}
		}
	}
}

func orther(header *protocol.Header,columns []*protocol.Column) {
	var key string
	for _, col := range columns {
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
			//fmt.Println("type",reflect.TypeOf(col.GetValue()))
		}

	}

	fmt.Println("other event:",key,header.GetEventType())


}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

