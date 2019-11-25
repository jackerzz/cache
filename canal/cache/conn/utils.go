package conn

import (
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/agrison/go-commons-lang/stringUtils"

)

// iso-8859-1 to utf-8
func Format(header *protocol.Header,columns []*protocol.Column) (string,error){

	for _,col := range columns{
		if (stringUtils.ContainsIgnoreCase(col.GetMysqlType(),"BLOB") ||
			stringUtils.ContainsIgnoreCase(col.GetMysqlType(),"BINARY")){
				// 类型转换
		}else {
			// 直接输出
		}
	}

}