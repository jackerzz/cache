package cache

import (
	"cache/gredis"
	"cache/logic/model"
	"strconv"
)

const (
	DeviceIPKey = "device_ip:"
)

type deviceIPCache struct{}

var DeviceIPCache = new(deviceIPCache)

func (c *deviceIPCache) Key(deviceId int64) string {
	return DeviceIPKey + strconv.FormatInt(deviceId, 10)
}

func (c *deviceIPCache) Get(deviceId int64) (*model.Device,error) {
	var device model.Device
	err := gredis.Gets(c.Key(deviceId),&device)
	if err != nil {
		return nil, err
	}
	return &device,nil
}

func (c *deviceIPCache) Set(deviceId int64, ip string) error {
	err:=gredis.Set(c.Key(deviceId),ip,0);if err !=nil {
		return err
	}
	return nil
}

func (c *deviceIPCache) Del(deviceId int64) error {
	_,err := gredis.Delete(c.Key(deviceId))
	if err != nil {
		return err
	}
	return nil
}
