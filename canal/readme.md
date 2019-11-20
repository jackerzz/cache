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
   
## canal 部署
    mkdir -p /tmp/canal
    mkdir -p /tmp/canal/canal-deployer
    mkdir -p /tmp/canal/canal-admin
    tar zxvf canal.deployer-$version.tar.gz  -C /tmp/canal
    https://my.oschina.net/teddyIH/blog/3102095
    
## 配置修改

## 启动&关闭 (/tmp/canal)
    sh bin/startup.sh
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