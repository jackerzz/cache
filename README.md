## 分布式缓存
    memberlist ----> github.com/hashicorp/memberlist
         用来管理分布式集群内节点发现、 节点失效探测、节点列表的软件包。
    consistent hash(一致性哈希算法) --->stathat.com/c/consistent
    redis 分布式存储     
                
## 读写分离
    mysql binlog 将数据刷新到redis 并设置合理的有效时间，redis 只提供数据查询
    优点：增量式的导入保证业务上的all 查询都可miss
    缺点：由于增量式，会导致有很多无用的key以及内存紧张
    
       
## [Elasticsearch](https://github.com/jackerzz/cache/tree/master/doc/es.md)
    接收来自 canal 的数据增量写入， 并且创建对应的索引
    通过对应的查询方式取出字段，在Kibana/go-echarts 中进行数据可视化
    

## doc 
    所有文档以及笔记
## lib 
    一些杂物存放地方
## 用到的package
> [ini中文](https://ini.unknwon.io/docs/intro/getting_started)
>> [task任务调度](github.com/robfig/cron)
>>> [redis命令](http://doc.redisfans.com/)
>>>> [redis 设计](https://github.com/jackerzz/cache/tree/master/doc/redis.md)
