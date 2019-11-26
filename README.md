# cache
    client              客户端测试
    conf                ini 配置
    scheduling          检测模块调度中心
    gdb                 mysql持久化
    gredis              redis缓存
    task                任务中心
    test                单元测试    
    utlis               工具包
    main.go             启动器
    
## conf
   [ini中文](https://ini.unknwon.io/docs/intro/getting_started)
## scheduling
## gdb
   [mysql gorm]()
## task
   [task任务调度](github.com/robfig/cron)
## gredis<done>
   [redis package](github.com/gomodule/redigo/redis)
   [redis命令](http://doc.redisfans.com/)
   写入支持乐观锁


## main.go
   [数据处理](https://www.processon.com/view/link/5daea29be4b0893e999df7b1)
    
        
## 新版 分布式缓存（Memcached 内存式实现） undone
         采用：
         memberlist ----> github.com/hashicorp/memberlist
                用来管理分布式集群内节点发现、 节点失效探测、节点列表的软件包。
         consistent hash(一致性哈希算法) --->stathat.com/c/consistent
                https://studygolang.com/articles/13997
                
## 读写分离
    mysql binlog 将数据刷新到redis 并设置合理的有效时间
    redis 只提供数据查询
    
                    
                
## undone
    sentinel----> 实现redis 读写分离
    Memcached
    RabbitMQ
    es