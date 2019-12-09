# [aggregations Catalo]()
> [聚合分析](https://www.google.com/search?q=elasticsearch+%E8%81%9A%E5%90%88%E4%BD%BF%E7%94%A8&oq=elasticsearch+%E8%81%9A%E5%90%88%E4%BD%BF%E7%94%A8&aqs=chrome..69i57.19542j0j1&sourceid=chrome&ie=UTF-8)
 ---
    "aggregations" : {
        "<aggregation_name>" : { <!--聚合的名字 -->
            "<aggregation_type>" : { <!--聚合的类型 -->
                <aggregation_body> <!--聚合体：对哪些字段进行聚合 -->
            }
            [,"meta" : {  [<meta_data_body>] } ]? <!--元 -->
            [,"aggregations" : { [<sub_aggregation>]+ } ]? <!--在聚合里面在定义子聚合 -->
        }
        [,"<aggregation_name_2>" : { ... } ]*<!--聚合的名字 -->
    }
    注：
        aggregations 也可简写为 aggs
   
>> 指标聚合
  ---
    1. max min sum avg
    2. 文档计数 count
    3. Value count 统计某字段有值的文档数
    4. cardinality  值去重计数
    5. stats 统计 count max min avg sum 5个值
    6. Extended stats 
        高级统计，比stats多4个统计结果： 平方和、方差、标准差、平均值加/减两个标准差的区间
    7. Percentiles 占比百分位对应的值统计
        对指定字段（脚本）的值按从小到大累计每个值对应的文档数的占比（占所有命中文档数的百分比），
        返回指定占比比例对应的值。默认返回[ 1, 5, 25, 50, 75, 95, 99 ]分位上的值。如下中间的结果，
        可以理解为：占比为50%的文档的age值 <= 31，或反过来：age<=31的文档数占总命中文档数的50%
    8. Percentiles rank 统计值小于等于指定值的文档占比
    9. Geo Bounds aggregation 求文档集中的地理位置坐标点的范围
    10. Geo Centroid aggregation  求地理位置中心点坐标值
>>> 桶聚合
   ---
     1. Terms Aggregation  根据字段值项分组聚合
        请求：
            size 指定返回多少个分组
            shard_size 的默认值为：
            索引只有一个分片：= size
            多分片：= size * 1.5 + 10
            order  指定分组的排序
        响应：
            "doc_count_error_upper_bound": 0：文档计数的最大偏差值
            "sum_other_doc_count": 463：未返回的其他项的文档数 
            
     2.  filter Aggregation  对满足过滤查询的文档进行聚合计算
      在查询命中的文档中选取符合过滤条件的文档进行聚合，先过滤再聚合
      
     3. Filters Aggregation  多个过滤组聚合计算
     
     4. Range Aggregation 范围分组聚合
     
     5. Date Range Aggregation  时间范围分组聚合
     
     6. Date Histogram Aggregation  时间直方图（柱状）聚合
         就是按天、月、年等进行聚合统计。
         可按 year (1y), quarter (1q), month (1M), week (1w), day (1d), hour (1h), minute (1m), second (1s) 
         间隔聚合或指定的时间间隔聚合。
         
     7. Missing Aggregation  缺失值的桶聚合
     
     8. Geo Distance Aggregation  地理距离分区聚合        
# [路人笔记](https://www.cnblogs.com/leeSmall/category/1210814.html)
> [elasticsearch系列一：elasticsearch（ES简介、安装&配置、集成Ikanalyzer）](https://www.cnblogs.com/leeSmall/p/9189078.html)
>>[elasticsearch系列二：索引详解（快速入门、索引管理、映射详解、索引别名）](https://www.cnblogs.com/leeSmall/p/9193476.html)
>>>[elasticsearch系列三：索引详解（分词器、文档管理、路由详解（集群））](https://www.cnblogs.com/leeSmall/p/9195782.html)
>>>>[elasticsearch系列四：搜索详解（搜索API、Query DSL）](https://www.cnblogs.com/leeSmall/p/9206641.html)
>>>>>[elasticsearch系列五：搜索详解（查询建议介绍、Suggester 介绍）](https://www.cnblogs.com/leeSmall/p/9206646.html)
>>>>>>[elasticsearch系列六：聚合分析（聚合分析简介、指标聚合、桶聚合）](https://www.cnblogs.com/leeSmall/p/9215909.html)
>>>>>>>[elasticsearch系列七：ES Java客户端](https://www.cnblogs.com/leeSmall/p/9218779.html)
>>>>>>>>[elasticsearch系列八：ES 集群管理（集群规划、集群搭建、集群管理）](https://www.cnblogs.com/leeSmall/p/9220535.html)

