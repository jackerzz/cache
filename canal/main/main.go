package main

import (
	"cache/lib/gredis"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"
)

func main() {

	connector := client.NewSimpleCanalConnector("192.168.225.131", 11111, "", "", "master", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = connector.Subscribe("go\\..*")
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
			fmt.Println(header.GetSchemaName())

			//================> binlog[mysql-bin.000005 : 17275],name[`HAH`,], eventType: QUERY
			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE{
					gredis.DelRedis(header,rowData.AfterColumns)
				} else if eventType == protocol.EventType_INSERT {
					gredis.SetRedis(header,rowData.AfterColumns)
				} else {
					fmt.Println("...........................")
					gredis.SetRedis(header,rowData.AfterColumns)
				}
			}
		}
	}
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}



