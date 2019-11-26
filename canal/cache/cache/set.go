package cache

import (
	protocol "github.com/CanalClient/canal-go/protocol"
)

func SetRedis(header *protocol.Header,columns []*protocol.Column) {
	var value []string
	var key string
	for _, col := range columns {
		value = append(value,col.GetName()+":"+col.GetValue())
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}

	}
	Set(key,value,6000)
}

