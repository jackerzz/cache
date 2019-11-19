package server

import (
	"fmt"
	"log"
	"os"
	"time"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"
)

func ReadToCanal() {
	for {
		message, err := CanalClient.Get(100,nil,nil)
		if err !=nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(600 *time.Microsecond)
			fmt.Println("==========正在等待数据============")
			continue
		}
		// 消息处理列
		printEntry(message.Entries)
	}
}
func printEntry(entrys []protocol.Entry){
	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}

		rowChange := new(protocol.RowChange)
		err := proto.Unmarshal(entry.GetStoreValue(),rowChange);
		checkError(err)

		if rowChange != nil {
			eventype := rowChange.GetEventType()
			header := entry.Header
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))


			for _, rowData := range rowChange.GetRowDatas() {
				if eventype == protocol.EventType_DELETE{
					// event type == delete
					printColumn(rowData.GetAfterColumns())
				}else if eventype == protocol.EventType_INSERT {
					// evnet type == insert
					printColumn(rowData.GetAfterColumns())
				}else if eventype == protocol.EventType_UPDATE{
					printColumn(rowData.GetAfterColumns())
				}else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}


	}
}

func printColumn(columns []*protocol.Column)  {
	// 解析数据
}



func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}