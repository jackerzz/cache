package setting

import "log"

// canal 专用添加需要处理的数据库表
const (
	// game 库下的 user 表
	game_user 	=1
	game_app    =2
)
var messagedb map[int]string

func GetMessagedb(code int) string {
	msg, err := messagedb[code];if err {
		log.Println("get message db table err",err)
		return ""
	}
	return msg
}


func init(){
	messagedb = make(map[int]string)
	messagedb[game_user] = "user"
	messagedb[game_app]  =	"app"
}
