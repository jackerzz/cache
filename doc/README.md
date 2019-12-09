## CANAL ubuntu 18.4 server 部署
*  [下载canal](https://github.com/alibaba/canal/releases)
*  [官方文档](https://github.com/alibaba/canal/wiki)

## 安装rancher（v2.x）
    docker pull rancher/rancher
    查看rancher/rancher info
    docker inspect rancher/rancher:latest
        
## 新建宿主挂载
    mkdir -p /home/jacker/rancher_home/rancher
    mkdir -p /home/jacker/rancher_home/auditlog
## 启动rancher/rancher 容器
* [Rancher v2.x](https://www.bookstack.cn/books/rancher-v2.x)
    > 这个版本本身对配置要求较高，但是提供更多的服务
* [Rancher v1.x](https://www.bookstack.cn/read/rancher-v1.x/30f7507a551232e8.md)
    > docker run -d --restart=unless-stopped -p 80:80 -p 443:443 -v /home/jacker/rancher_home/rancher:/var/lib/rancher -v /home/jacker/rancher_home/auditlog:/var/log/auditlog --name rancher rancher/rancher
    >> 查看run
    >>> docker container ls
    >>>> Rancher v1.x    
    >>>>> mkdir -p /home/jacker/rancher_home/rancher
    >>>>>> mkdir -p /home/jacker/rancher_home/auditlog
    >>>>>>> docker run -d --name=rancher --restart=unless-stopped  -p 80:80 -p 9090:8080 -v /home/rancher_home/rancher:/var/lib/rancher -v /home/rancher_home/auditlog:/var/log/auditlog rancher/server
    >>>>>>>> 添加主机
    >>>>>>>>> .......
  
## 容器批量操作
    查看所有容器
        docker ps -a   
    进入容器
        其中字符串为容器ID:
        docker exec -it d27bd3008ad9 /bin/bash
    停用全部运行中的容器:
        docker stop $(docker ps -q)
    删除全部容器：
        docker rm $(docker ps -aq)
    一条命令实现停用并删除容器：
        docker stop $(docker ps -q) & docker rm $(docker ps -aq)
    删除所有未运行 Docker 容器
        docker rm $(docker ps -a -q)
    删除所有镜像
        docker rmi $(docker images -q)
    根据格式删除所有镜像
        docker rm $(docker ps -qf status=exited)