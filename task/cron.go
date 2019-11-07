package task

import (
	"github.com/robfig/cron"
	"time"
)

// to done : 注册到 main.go
func Setup() error {
	c := cron.New()
	c.AddFunc("*********缓存到redis********", func() {
		// 添加需要执行的方法
	})
	c.AddFunc("*********持久化到mysql******************", func() {
		// 添加需要执行的方法
	})






	c.Start()


	t1 := time.NewTimer(time.Second * 3600)

	for {
		select {
		case <- t1.C:
			t1.Reset(time.Second * 20)
		}
	}
}
