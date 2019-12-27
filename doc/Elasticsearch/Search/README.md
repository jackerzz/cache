# 索引
> 查看集群的健康状况
 ---
     http://localhost:9200/_cat/health?v
     Green - everything is good (cluster is fully functional)，即最佳状态
     Yellow - all data is available but some replicas are not yet allocated (cluster is fully functional)，即数据和集群可用，但是集群的备份有的是坏的
     Red - some data is not available for whatever reason (cluster is partially functional)，即数据和集群都不可用   