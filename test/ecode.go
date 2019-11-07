package ecodes

import (
	"encoding/json"
	"fmt"
)

func Ecode(b []byte){
	var inter interface{}
	err := json.Unmarshal(b, &inter);if err != nil {
		fmt.Println("ecode:",err)
	}
	book, ok := inter.(map[string]interface{});if ok {
		for k, v := range book {
			switch vt := v.(type) {
			case float64:
				fmt.Println(k, " is float64 ", vt)
			case string:
				fmt.Println(k, " is string ", vt)
			case []interface{}:
				fmt.Println(k, " is an array:")
				for i, iv := range vt {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println("illegle type")
			}
		}
	}

}
