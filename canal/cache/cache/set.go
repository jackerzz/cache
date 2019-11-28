package cache

import (
	protocol "github.com/CanalClient/canal-go/protocol"
)

func SetRedis(header *protocol.Header,column []*protocol.Column) {
	var value []string
	var key string
	for _, col := range column {
		value = append(value,col.GetName()+":"+col.GetValue())
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}

	}
	Set(key,value,6000)
}

func HsetRedis(header *protocol.Header,column []*protocol.Column){
	var key string
	value := make(map[string]string)
	for _, col := range column {
		value[col.GetName()]=col.GetValue()
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}
		Hset(key,value)
	}
}