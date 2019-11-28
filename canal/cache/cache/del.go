package cache

import (
	protocol "github.com/CanalClient/canal-go/protocol"
)

func DelRedis(header *protocol.Header,columns []*protocol.Column,deltype string) {
	var key string
	for _, col := range columns {
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}
	}
	if deltype == "string"{
		Delete(key)
	}else if deltype == "hset"{
		Hdel(key)
	}

}
