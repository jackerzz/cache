package cache

import (
	"cache/cache/gredis"
	"cache/cache/logic/model"
	"strconv"
	"time"
)
const (
	UserKey    = "user:"
	UserExpire = 2 * time.Hour
)
type userCache struct {}
var UserCache = new(userCache)

func (*userCache) Key(appId, userId int64) string {
	return UserKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(userId, 10)
}

func (c *userCache) Get(appid,userId int64) (*model.User,error) {
	var user model.User
	err := gredis.Gets(c.Key(appid,userId), &user)
	if err != nil{
		return nil,err
	}
	return &user, nil
}

func (c *userCache) set(user model.User) error {
	err := gredis.Set(c.Key(user.AppId, user.UserId), user, UserExpire)
	if err !=nil {
		return err
	}
	return nil
}

func (c *userCache)Del(appId, userId int64) error{
	_, err :=gredis.Delete(c.Key(appId,userId))
	if err !=nil{
		return err
	}
	return nil

}