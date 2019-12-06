## conf
   [ini中文](https://ini.unknwon.io/docs/intro/getting_started)
## task
   [task任务调度](github.com/robfig/cron)
## gredis<done>
   [redis package](github.com/gomodule/redigo/redis)
   [redis命令](http://doc.redisfans.com/)

## doc 
    所有文档以及笔记
  
## 分布式缓存
         采用：
         memberlist ----> github.com/hashicorp/memberlist
                用来管理分布式集群内节点发现、 节点失效探测、节点列表的软件包。
         consistent hash(一致性哈希算法) --->stathat.com/c/consistent
                https://studygolang.com/articles/13997
                
## 读写分离
    mysql binlog 将数据刷新到redis 并设置合理的有效时间
    redis 只提供数据查询
       
## Elasticsearch
    接收来自 canal 的数据写入， 并且创建对应的索引
    通过对应的查询方式取出字段，在Kibana/go-echarts 中进行数据可视化
    
## lib 
    一些杂物存放地方