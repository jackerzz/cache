package cache

import (
	"fmt"
	protocol "github.com/CanalClient/canal-go/protocol"
)

func SetRedis(header *protocol.Header,columns []*protocol.Column) {
	var value []string
	var key string
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
		fmt.Println(col.GetName()+":"+col.GetValue())
		value = append(value,col.GetName()+":"+col.GetValue())
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}

	}
	fmt.Println("value:",value)
	fmt.Println("key:",key)
	Set(key,value,6000)
	r,_ := Get(key)
	fmt.Println("r:",r)

}