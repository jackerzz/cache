package main

import (
	"cache/gredis"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"
)

func main() {

	connector := client.NewSimpleCanalConnector("192.168.225.129", 11111, "", "", "master", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = connector.Subscribe("game\\..*,test\\..*,page\\..*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		printEntry(message.Entries)

	}
}

func printEntry(entrys []protocol.Entry) {

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
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))


			//================> binlog[mysql-bin.000005 : 17275],name[`HAH`,], eventType: QUERY
			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE{
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == protocol.EventType_INSERT {
					printColumn(rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
}

func printColumn(columns []*protocol.Column) {
	var value []string
	var key string
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
		fmt.Println(col.GetName()+":"+col.GetValue())
		value = append(value,col.GetName()+":"+col.GetValue())
		if col.IsKey{
			key = "db:"+col.GetValue()
		}

	}
	fmt.Println(value,key)
	gredis.Set(key,value,6000)
	r,_ := gredis.Get(key)
	fmt.Println("r:",r)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
