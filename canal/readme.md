## mysql 数据刷新到redis 
    
    https://www.cnblogs.com/duanxz/p/5062833.html
    
    https://github.com/alibaba/canal/releases
    备选：
        https://github.com/siddontang/go-mysql/blob/master/README.md
    git clone https://github.com/withlin/canal-go.git
    
## mysql
    docker run --name mysql -di  \
    -v /home/cananl/mysql/conf:/etc/mysql/conf.d \
    -v /home/cananl/mysql/data:/var/lib/mysql \
    -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:5.7
    
    在 /home/cananl/mysql/conf 下新增mysql配置文件 来开启 mysql binlog模式
   
## canal 部署
    mkdir -p /tmp/canal
    mkdir -p /tmp/canal/canal-deployer
    mkdir -p /tmp/canal/canal-admin
    tar zxvf canal.deployer-$version.tar.gz  -C /tmp/canal
    https://my.oschina.net/teddyIH/blog/3102095
    
## 配置修改
    vi conf/example/instance.properties
    
    ## mysql serverId
    canal.instance.mysql.slaveId = 1234
    #position info，需要改成自己的数据库信息
    canal.instance.master.address = 127.0.0.1:3306 
    canal.instance.master.journal.name = 
    canal.instance.master.position = 
    canal.instance.master.timestamp = 
    #canal.instance.standby.address = 
    #canal.instance.standby.journal.name =
    #canal.instance.standby.position = 
    #canal.instance.standby.timestamp = 
    #username/password，需要改成自己的数据库信息
    canal.instance.dbUsername = canal  
    canal.instance.dbPassword = canal
    canal.instance.defaultDatabaseName =
    canal.instance.connectionCharset = UTF-8
    #table regex
    canal.instance.filter.regex = .\*\\\\..\*
## 启动&关闭 (/tmp/canal)
    sh bin/startup.sh
    sh bin/stop.sh
## log
    查看 server 日志
    vi logs/canal/canal.log</pre>
    
    查看 instance 的日志
    vi logs/example/example.log