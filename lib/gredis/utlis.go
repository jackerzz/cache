package gredis

import "github.com/json-iterator/go"

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
