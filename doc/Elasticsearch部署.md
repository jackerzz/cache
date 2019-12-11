## 谷歌云部署
>  创建项目 
* 项目名称 my-app
> > 修改防火墙规则或者创建防火墙入站规则
  ---
    目标 
        目标标记 http-server 
    来源过滤条件 IP 地址范围 
        0.0.0.0/0 
    协议和端口
        tcp:9200
    
>>> 拉取es文件
 ---
    curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.4.2-linux-x86_64.tar.gz
    tar -zxf elasticsearch-7.4.2-linux-x86_64.tar.gz
    
    
## Elasticsearch 配置

> conf/Elasticsearch.yml
 
    cluster.name: my_app
    node.name: node-1
    network.host: 0.0.0.0
    http.port: 9200
    discover.seed_hosts: ["127.0.0.1","::1"]
    cluster.initial_master_nodes: ["node-1"]

>> 异常处理
    
    编辑 /etc/security/limits.conf，追加以下内容；
    soft nofile 65536
    hard nofile 65536
    编辑 /etc/sysctl.conf，追加以下内容：
    vm.max_map_count=655360
    保存后，执行：
        sysctl -p
    重新启动，成功

>>> 启动elasticsearch
    
     bash bin/elasticsearch &
     
     注： elasticsearch 7*的包中自带jdk11无需安装，但是系统默认装jdk8，会在打印日志异常，由于jdk向下兼容