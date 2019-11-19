package server

import (
	"cache/setting"
	"github.com/CanalClient/canal-go/client"
	"log"
)
var CanalClient *client.SimpleCanalConnector


// NewSimpleCanalConnector
func SetupCanal() error {
	CanalClient = &client.SimpleCanalConnector{
		Address:           setting.CanalSetting.Address,
		Port:              setting.CanalSetting.Port,
		UserName:          setting.CanalSetting.UserName,
		PassWord:          setting.CanalSetting.PassWord,
		SoTime:            setting.CanalSetting.SoTime,
		IdleTimeOut:       setting.CanalSetting.IdleTimeOut,
		RollbackOnConnect: true,
	}
	err := CanalClient.Connect();if err != nil {
		return err
	}
	return nil
}

// 添加匹配规则
func MatchingTable() {
	err := CanalClient.Subscribe(".*\\\\..*")
	if err !=nil {
		log.Println("Subscribe:",err)

	}
	err = CanalClient.UnSubscribe()
	if err != nil {
		log.Println("UnSubscribe",err)
	}

}

func init(){
	SetupCanal()
	MatchingTable()
}