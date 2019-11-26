package conn

import (
	"cache/canal/cache/core"
	"github.com/CanalClient/canal-go/client"
	"log"
	"os"
	"time"
)

func Core(){
	connector := client.NewSimpleCanalConnector("192.168.225.131", 11111, "", "", "master", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = connector.Subscribe("baby_api_admin\\..*,go\\..*")
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

		core.PrintEntry(message.Entries)

	}
}