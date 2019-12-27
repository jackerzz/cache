##[Avg](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_avg_test.go)
    POST /exams/_search?size=0
    {
        "aggs" : {
            "avg_grade" : { "avg" : { "field" : "grade" } }
        }
    }
    Response:
    {
        ...
        "aggregations": {
            "avg_grade": {
                "value": 75.0
            }
        }
    }
## [Cardinality](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_cardinality_test.go)
    计算与查询匹配的唯一,单值度量聚合，计算不同值的近似计数
    precision_threshold     允许以内存换取准确性          options       3000(默认)~~~~40000(max)
    
    POST /sales/_search?size=0
    {
        "aggs" : {
            "type_count" : {
                "cardinality" : {
                    "field" : "type",
                    "precision_threshold": 100 
                }
            }
        }
    }
    Response：
    {
        ...
        "aggregations" : {
            "type_count" : {
                "value" : 3
            }
        }
    }
## [Extended Stats]()
    统计信息聚合的扩展版本，在其中添加了其他指标，例如： 
        sum_of_squares, variance, std_deviation and std_deviation_bounds
    GET /exams/_search
    {
        "size": 0,
        "aggs" : {
            "grades_stats" : { "extended_stats" : { "field" : "grade" } }
        }
    }
    Response：
    {
        ...
    
        "aggregations": {
            "grades_stats": {
               "count": 2,
               "min": 50.0,
               "max": 100.0,
               "avg": 75.0,
               "sum": 150.0,
               "sum_of_squares": 12500.0,
               "variance": 625.0,
               "std_deviation": 25.0,
               "std_deviation_bounds": {
                "upper": 125.0,
                "lower": 25.0
               }
            }
        }
    }    
## [Geo Bounds](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_geo_bounds_test.go)
    PUT /museums
    {
        "mappings": {
            "properties": {
                "location": {
                    "type": "geo_point"   ------->度量聚合，计算包含一个字段(location)的所有geo_point值的边界框
                }
            }
        }
    }
    详:geo.text
## [Geo Centroid](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_geo_centroid_test.go)
    POST /museums/_search?size=0
    {
        "aggs" : {
            "centroid" : {
                "geo_centroid" : {           ----> 根据Geo-point字段的所有坐标值计算加权的度量聚合
                    "field" : "location"   
                }
            }
        }
    }

* [Stats](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_stats_test.go)
    ---
        Response： min, max, sum, count and avg
* [Sum](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_sum_test.go)
    --
        POST /sales/_search?size=0
        {
            "query" : {
                "constant_score" : {
                    "filter" : {
                        "match" : { "type" : "hat" }
                    }
                }
            },
            "aggs" : {
                "hat_prices" : { "sum" : { "field" : "price" } }
            }
        }
* [Top Hits]()
    -- 
        最佳命中(Top Hits)：

* [sort](https://sourcegraph.com/github.com/olivere/elastic/-/blob/sort_test.go)
    --  
        asc             Sort in ascending order
        desc            Sort in descending order
        min             Pick the lowest value
        sum             Use the sum of all values as sort value. Only applicable for number based array fields
        avg             Use the average of all values as sort value. Only applicable for number based array fields.
        median          Use the median of all values as sort value. Only applicable for number based array fields.
      
        详：geo.text/sort
* [Value Count](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_value_count_test.go)
* [Max](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_max_test.go)
* [Min](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_min_test.go)
* [Percentiles](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_percentiles.go)
* [Percentile Ranks](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_percentile_ranks.go)