## mysql 数据刷新到redis 
    
    https://www.cnblogs.com/duanxz/p/5062833.html
    
    https://github.com/alibaba/canal/releases
    备选：
        https://github.com/siddontang/go-mysql/blob/master/README.md
        git clone https://github.com/withlin/canal-go.git
    重点：(基础配置)
        https://github.com/alibaba/canal/wiki/AdminGuide
## mysql
    docker run --name mysql -di  \
    -v /home/cananl/mysql/conf:/etc/mysql/conf.d \
    -v /home/cananl/mysql/data:/var/lib/mysql \
    -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:5.7
    
    在 /home/cananl/mysql/conf 下新增mysql配置文件 来开启 mysql binlog模式
   
## canal 部署（容器版）
    mkdir -p /tmp/canal
    mkdir -p /tmp/canal/canal-deployer
    mkdir -p /tmp/canal/canal-admin
    tar zxvf canal.deployer-$version.tar.gz  -C /tmp/canal
    https://my.oschina.net/teddyIH/blog/3102095
    
## 配置修改

## 启动&关闭 (/tmp/canal)
    sh bin/startup.sh
    sh bin/startup.sh local
    sh bin/stop.sh
## log
    查看 server 日志
    vi logs/canal/canal.log</pre>
    
    查看 instance 的日志
    vi logs/example/example.log
    
    
## 对于开启mysql binlong 的日志激增
    [mysqld]  
    log-bin=mysql-bin #添加这一行就ok  
    binlog-format=ROW #选择row模式  
    server_id=1 #配置mysql replaction需要定义，不能和canal的slaveId重复
    #日志超过7天自动过期
    expire_logs_days = 7
    ====================================================================================================================
    MySQL的四种BLOB类型:
    类型         大小(单位：字节)
    　　TinyBlob     最大255
    　　Blob         最大65K
    　　MediumBlob   最大16M
    　　LongBlob     最大4G
    linux通过etc/my.cnf
    [mysqld]
        max_allowed_packet = 10M 
## 对于canal 启动的一些问题
    1.
        问题：
            server 启动显示正常，但日志有报错
            instance 启动显示正常，但没有任何日志信息 
        解决：https://github.com/alibaba/canal/issues/2129
            sh startup.sh local
        =================================================================================================================
    2.
        instance管理中
            instance 名称 master 对应client.NewSimpleCanalConnector()中destination 参数默认的为example
        server管理
            canal.properties  默认开启（canal instance user/passwd）认证
            并且添加
                canal.destinations = master
        https://github.com/alibaba/canal/issues/1214
## connector.Subscribe（）
    过滤数据库和表的:
        设置： instance.properties文件的 
                canal.instance.filter.regex=
                或者直接写在connector.Subscribe（）
        
        1. 所有表：.* or .*\\..*
        
        2. canal schema下所有表： canal\\..* （canal为数据库）
        
        3. canal下的以canal打头的表：canal\\.canal.*
        
        4. canal schema下的一张表：canal.test1
        
        5. 多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)
    =================================================================================================================
    发往topic:
        instance.properties文件的 canal.mq.dynamicTopic=
        
        canal 1.1.3版本之后, 支持配置格式：schema 或 schema.table，多个配置之间使用逗号或分号分隔
        
        例子1：test\\.test 指定匹配的单表，发送到以test_test为名字的topic上
        例子2：.*\\..* 匹配所有表，则每个表都会发送到各自表名的topic上
        例子3：test 指定匹配对应的库，一个库的所有表都会发送到库名的topic上
        例子4：test\\.* 指定匹配的表达式，针对匹配的表会发送到各自表名的topic上
        例子5：test,test1\\.test1，指定多个表达式，会将test库的表都发送到test的topic上，
              test1\\.test1的表发送到对应的test1_test1 topic上，其余的表发送到默认的canal.mq.topic值
    =================================================================================================================
    灵活性：
        允许对匹配条件的规则指定发送的topic名字，配置格式：topicName:schema 或 topicName:schema.table
        例子1: test:test\\.test 指定匹配的单表，发送到以test为名字的topic上
        例子2: test:.*\\..* 匹配所有表，因为有指定topic，则每个表都会发送到test的topic下
        例子3: test:test 指定匹配对应的库，一个库的所有表都会发送到test的topic下
        例子4：testA:test\\.* 指定匹配的表达式，针对匹配的表会发送到testA的topic下
        例子5：test0:test,test1:test1\\.test1，指定多个表达式，会将test库的表都发送到test0的topic下，
                   test1\\.test1的表发送到对应的test1的topic下，其余的表发送到默认的canal.mq.topic值
    =================================================================================================================
    重要说明
    canal.mq.dynamicTopic	    canal.instance.filter.regex	            发送topic	                    描述
    .\..*	                        d1.t,d2.t,d3.t	                        d1,d2,d3	                同一个数据库发送到以数据库为名字的topic
    .\...*	                        d1.t,d2.t,d3.t	                        d1,d2,d3	                同一个数据库发送到以数据库为名字的topic
    .*\\..*                 	d1.t,d2.t,d3.t	                        d1_t,d2_t,d3_t	                单表单topic，数据库和表名为topic
    t:d.*\\.t	                d1.t,d2.t,d3.t	                        t	                        不同库同表，以表名为topic
    =================================================================================================================
    支持指定topic名称匹配, 配置格式：topicName:schema 或 schema.table，多个配置之间使用逗号分隔, 多组之间使用 ; 分隔
    =================================================================================================================
    
    https://blog.csdn.net/weixin_40126236/article/details/92654961
    https://gitee.com/starcwang/canal_mysql_elasticsearch_sync
## Example
    CREATE TABLE `us` (
       `id` int(12) NOT NULL,
       `name` char(23) DEFAULT NULL,
       `age` int(23) DEFAULT NULL,
       `phone` int(23) DEFAULT NULL,
       PRIMARY KEY (`id`)
     ) ENGINE=InnoDB DEFAULT CHARSET=latin1
     =================================================================================================================
     Run 结果
     ================> binlog[mysql-bin.000010 : 56734],name[game,us], eventType: UPDATE
     -------> before
     id : 3  update= false
     name : fsdjakfj  update= false
     age : 23232  update= false
     phone : 423  update= false
     -------> after
     id : 3  update= false
     name : fsdjakfjfdfdsfds  update= true
     age : 23232  update= false
     phone : 423  update= false
     ================> binlog[mysql-bin.000010 : 56462],name[game,us], eventType: INSERT
     id : 5  update= true
     name : asfds  update= true
     age : 45  update= true
     phone : 343  update= true
     ================> binlog[mysql-bin.000010 : 57040],name[game,us], eventType: DELETE
     id : 5  update= false
     name : asfds  update= false
     age : 45  update= false
     phone : 343  update= false