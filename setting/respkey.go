package setting


// 需要新增 类型账本key
const (
	Account_book_string		=     1
	Account_book_hash		=  	  2
)
var Message map[int]string
func init() {
	Message= make(map[int]string)
	Message[Account_book_string] = "Account:book:string"
	Message[Account_book_hash] = "Account:book:hash"

}

func GetMessage(code int) string {
	msg, ok := Message[code]
	if ok {
		return msg
	}
	return ""
}