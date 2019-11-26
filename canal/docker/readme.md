## 安装rancher（v2.x）
    docker pull rancher/rancher
    查看rancher/rancher info
    docker inspect rancher/rancher:latest
        
## 新建宿主挂载
    mkdir -p /home/jacker/rancher_home/rancher
    mkdir -p /home/jacker/rancher_home/auditlog
## 启动rancher/rancher 容器
    Rancher v2.x （配置要求过高）   https://www.bookstack.cn/books/rancher-v2.x
    docker run -d --restart=unless-stopped -p 80:80 -p 443:443 -v /home/jacker/rancher_home/rancher:/var/lib/rancher -v /home/jacker/rancher_home/auditlog:/var/log/auditlog --name rancher rancher/rancher
    查看run
        docker container ls
    Rancher v1.x    https://www.bookstack.cn/read/rancher-v1.x/30f7507a551232e8.md
        docker run -d --name=rancher --restart=unless-stopped  -p 80:80 -p 9090:8080 -v /home/rancher_home/rancher:/var/lib/rancher -v /home/rancher_home/auditlog:/var/log/auditlog rancher/server
        mkdir -p /home/jacker/rancher_home/rancher
        mkdir -p /home/jacker/rancher_home/auditlog
        添加主机
        sudo docker run --rm --privileged -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/rancher:/var/lib/rancher rancher/agent:v1.2.11 https://192.168.225.129:8080/v1/scripts/3665C1981D5CD3BF32C7:1546214400000:tc3W1eiEtdckXFtIfxO3gBU
        添加
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