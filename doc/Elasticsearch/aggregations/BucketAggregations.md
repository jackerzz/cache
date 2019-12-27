* [Adjacency Matrix](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_bucket_adjacency_matrix.go)
    ---
        交叉存储，这个特性适合做多表关联
* [Weighted Avg Aggregation](https://sourcegraph.com/github.com/olivere/elastic/-/blob/search_aggs_metrics_weighted_avg_test.go)
    --
       ∑(value * weight) / ∑(weight)
       
       Name	            Description	                                                            Required Value
       value            The configuration for the field or script that provides the values      Required
       
       weight           The configuration for the field or script that provides the weights     Required
       
       format           The numeric response formatter                                          Optional
       
       value_type       A hint about the values for pure scripts or unmapped fields             Optional
       
       type MultiValuesSourceFieldConfig struct {
       	FieldName string
       	Missing   interface{}
       	Script    *Script
       	TimeZone  string
       }
       type WeightedAvgAggregation struct {
       	fields          map[string]*MultiValuesSourceFieldConfig
       	valueType       string
       	format          string
       	value           *MultiValuesSourceFieldConfig
       	weight          *MultiValuesSourceFieldConfig
       	subAggregations map[string]Aggregation
       	meta            map[string]interface{}
       }
   
* undone
 ---
    [Children]()
    [Auto-interval Date Histogram]()
    [Date Histogram]()
    [Date Range]()
    [Diversified Sampler]()
    [Filter]()
    [Filters]()
    [Geo Distance]()
    [Global]()
    [Histogram]()
    [IP Range]()
    [Missing]()
    [Nested]()
    [Range]()
    [Reverse Nested]()
    [Sampler]()
    [Significant Terms]()
    [Significant Text]()
    [Terms]()
    [Composite]()   
   
   
       