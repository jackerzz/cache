package gredis

import (
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/json-iterator/go"
)
func Sort(key string, data interface{})(){

}
// get cache Unserialization
func Gets(key string,data interface{}) error {
	bytes, err := Get(key);if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(bytes, data); if err != nil {
		return err
	}
	return nil
}

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

func DelRedis(header *protocol.Header,columns []*protocol.Column) {
	var key string
	for _, col := range columns {
		if col.IsKey{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}else if col.GetName()== "id"{
			key = header.GetSchemaName()+":"+header.GetTableName()+":"+col.GetValue()
		}
	}
	Delete(key)
}

