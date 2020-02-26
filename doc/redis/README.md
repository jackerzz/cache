## redis 主目录
> [redis key 存储设计]()
## redis 项目实际应用
> redis 构建内容定向广告系统、系统分析平台、搜索
>> memcached 多线程 结合redis 单线程
>>>vm instances部署redis
  ---
    系统： ubuntu18.4 
    硬盘： 10g
    内存： 1.7g
    sudo apt update
    sudo apt install redis-server
>>>> redis 配置
   ---
     sudo vim /etc/redis/redis.conf
     重新加载配置
     sudo service redis restart    
     查看redis状态
     sudo systemct1 status redis
>>>>> redis 控制
    ---
      sudo service redis start
      sudo service redis stop
      sudo service redis restart
      客户端连接
      redis-cli
      远程连接
      sudo vim /etc/redis/redis-conf
           修改 bind 127.0.0.1:: --> bind 0.0.0.0
           sudo service redis restart
>>>>>> redis 设置密码
    ----
        sudo vim /etc/redis/redis.conf
            requirepass sdfghjkl;fdjkljhgf      注：尽可能的复杂
