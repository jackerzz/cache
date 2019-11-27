# [redis](http://doc.redisfans.com/)
    redis 查询，更新，插入结构设计
    Bitmaps：
        setbit key offset value
        getbit key offset
        bitcount [start] [end]
        bitop op destkey key[key....]  Bitmaps间的运算
        bitpos key targetBit [start] [end] 计算Bitmaps中第一个值为targetBit的偏移量
    HyperLogLog:
        pfadd key element [element …]                   添加
        pfcount key [key …]                             计算独立用户数
        pfmerge destkey sourcekey [sourcekey ...]       合并
    publish    
        publish channel message
        subscribe channel [channel ...]
        unsubscribe [channel [channel ...]]
        按照模式订阅和取消订阅
            psubscribe pattern [pattern...]
            punsubscribe [pattern [pattern ...]]
        pubsub channels [pattern]                  查看活跃的频道
        pubsub numsub [channel ...]                查看频道订阅数
        pubsub numpat                              查看模式订阅数 
    set
        sunionstore destination  key [key]        
## [结构设计](https://blog.csdn.net/w13528476101/article/details/70146064)
    全局索引：
        key值构建：idx:[表名]
        value值构建：主键set集合
        sadd idx:user:1....n
        sort idx:user get go:user:*
            缺陷：在redis分片集群中，如果data:student:x[1,2,3,4,5]与idx:student不完全在同一个集群，则不支持sort get 命令组合
            lua && Piplining
    分页查询：
        sort idx:user get go:user:* limit 0 4                  ===> sort key get key:* limit offset count
    条件查询：
        key值构建：idx:[表名]:[字段名]:[字段值]
        value值构建：主键set集合
        设置查询key：sadd idx:user:add 1 2 3 4 5 6 7 8 9        ===> sadd key member
                查询：
                    sort idx:user:addr get go:user:*           ===> sort key get key:*
                    
        设置查询key：sunionstore destination  key [key]         ---> select * from user where addr in("gaodi","shanghai") 
                查询：
                    sort key get key:*
                    
        构建区间条件索引（field >= ? and field < ?）
            key值构建：idx:[表名]:[字段名]
            value值构建：主键zset集合([value:[主键],score:[字段值]])
                idx:student:age -> [{value:4,score:15},{value:2,score:16},{value:1,score:17},{value:3,score:17},{value:5,score:17}]
        设置查询key：
                查询：
        构建索引记录
            key值构建：record:[表名]:[主键]
            value值构建：索引set集合
                srem key  member [member.....]
                zrem key member [memeber....]
## key 过期策略
    1.使用zset把主键，过期参照字段存储起来
    2.使用redis键空间通知（keyspace notification）
    
## [新增压缩策略](https://github.com/golang/snappy)
* [Example](https://github.com/jackerzz/cache/blob/master/canal/cache/test/main.go)
