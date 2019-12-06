[kibana](https://www.elastic.co/guide/en/kibana/current/tutorial-build-dashboard.html)
---
    将数据集加载到Elasticsearch
    
    定义要连接到Elasticsearch的索引模式
    
    发现和探索数据
    
    可视化数据
    
    向仪表板添加可视化效果
---
* es ---> mysql
    --
        index ≈ table
        document ≈ record
        field ≈ column

[elasticsearch-head]()
---
    npm config set registry https://registry.npm.taobao.org
    git clone git://github.com/mobz/elasticsearch-head.git
    cd elasticsearch-head
    npm install
    npm run start
---
[es 旅程](http://you_ip:port/app/kibana#/dev_tools/console?_g=())

* 插入数据
    ---
        PUT /megacorp/employee/2
        {
          "first_name" :  "Jane",
            "last_name" :   "Smith",
            "age" :         32,
            "about" :       "I like to collect rock albums",
            "interests":  [ "music" ]
        }
  

* 查询
    ---
            GET /megacorp/employee/1                                        --> 根据_id查询
            GET /megacorp/employee/_search                                  --> all
            GET /megacorp/employee/_search?q=last_name:Smith                --> query last_name="Smith" 的字段
  


* dsl查询
    ---
        GET /megacorp/employee/_search
        {
            "query" : {
                "match" : {
                    "last_name" : "Smith"                                   --> 被查询的字段
                }
            }
        }
    
    
    
    
* filter（过滤器）
    ---
        GET /megacorp/employee/_search
        {
            "query" : {
                "filtered" : {
                    "filter" : {                                             --> 识别结构化搜索的限定条件
                        "range" : {
                            "age" : { "gt" : 30 } <1>                        --> 限定条件  注：gt 表示 大于
                        }
                    },
                    "query" : {
                        "match" : {
                            "last_name" : "Smith" <2>
                        }
                    }
                }
            }
        }
    


* 短语搜索+聚合
    --
        GET /megacorp/employee/_search
        {
            "query": {
                "match": {
                  "last_name": "smith"
                }
              },
            "aggs" : {
                "all_interests" : {
                    "terms" : { "field" : "interests" },
                    "aggs" : {
                        "avg_age" : {
                            "avg" : { "field" : "age" }
                        }
                    }
                }
            }
        }
    

* 多索引多类型
    ---
        URL	                            说明
        /_search	                    搜索所有的索引和类型
        /gb/_search	                    搜索索引gb中的所有类型
        /gb,us/_search	                搜索索引gb以及us中的所有类型
        /g*,u*/_search	                搜索所有以g或u开头的索引中的所有类型
        /gb/user/_search	            搜索索引gb中类型user内的所有文档
        /gb,us/user,tweet/_search	    搜索索引gb和 索引us中类型user以及类型tweet内的所有文档
        /_all/user,tweet/_search	    搜索所有索引中类型为user以及tweet内的所有文档
    
    
* 分页
    --
        参数	    说明
        size	每次返回多少个结果，默认值为10
        from	忽略最初的几条结果，默认值为0
    


* ik 
    --
        （1）先将其解压，将解压后的elasticsearch文件夹重命名文件夹为ik
        （2）将ik文件夹拷贝到elasticsearch/plugins 目录下。
        （3）重新启动，即可加载IK分词器













## [中文ik分词](https://github.com/medcl/elasticsearch-analysis-ik/releases)
    先将其解压，将解压后的elasticsearch文件夹重命名文件夹为ik
    将ik文件夹拷贝到elasticsearch/plugins 目录下。
    重新启动，即可加载IK分词器
## [Elasticsearch 权威指南](https://www.bookstack.cn/read/elasticsearch-definitive-guide-cn/README.md)
    Elasticsearch 是一个分布式的 RESTful 风格的检索和数据分析引擎，能够解决不同场景下的各种搜索、统计分析问题
## [数据表透视](https://github.com/flexmonster/pivot-kibana/)
## [kibana](https://www.elastic.co/cn/downloads/past-releases/kibana-6-8-3)
## [elasticsearch downloads](https://www.elastic.co/cn/downloads/past-releases/elasticsearch-6-8-3)
## [ik](https://github.com/medcl/elasticsearch-analysis-ik/releases)
## [elasticsearch-head](https://github.com/mobz/elasticsearch-head)

## [elastic-go](https://github.com/olivere/elastic)
## [ElasticHD](https://github.com/360EntSecGroup-Skylar/ElasticHD/releases/)