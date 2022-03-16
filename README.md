# GoFrame Template For SingleRepo

### CentOS 8 yum 相关命令
```
安装软件包：yum install <packageName> -y
卸载软件包：yum remove <packageName> -y
查找软件包：yum search <packageName>
列出所有可安装的软件包：yum list
列出所有可更新的软件包：yum list updates
列出所有已安装的软件包：yum list installed
列出所指定软件包：yum list <packageName>
升级指定软件包：yum update <packageName>
清除缓存目录(/var/cache/yum)下的软件包：yum clean packages
清除所有缓存：yum clean all
```

### CentOS 8 上下载并安装Go
```
1、wget https://dl.google.com/go/go1.18.linux-amd64.tar.gz
2、sudo tar -C /usr/local -xf go1.18.linux-amd64.tar.gz
3、vim ~/.bash_profile
4、export PATH=$PATH:/usr/local/go/bin
5、source ~/.bash_profile
6、go env -w GOPROXY="https://goproxy.cn,direct"
```

### CentOS 8 安装docker
```
1、卸载旧版本（如果有需要的话）
yum remove docker \
docker-client \
docker-client-latest \
docker-common \
docker-latest \
docker-latest-logrotate \
docker-logrotate \
docker-engine
2、设置仓库
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
3、安装 Docker Engine-Community
sudo yum install docker-ce docker-ce-cli containerd.io
4、如果提示接受 GPG 密钥，请选是
5、镜像加速，对于使用 systemd 的系统，请在 /etc/docker/daemon.json 中写入如下内容（如果文件不存在请新建该文件）：
{
  "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn","http://hub-mirror.c.163.com"]
}
6、重新启动服务
sudo systemctl daemon-reload
sudo systemctl restart docker
```

### docker 常用命令总结
```
启动docker：systemctl start docker
停止docker：systemctl stop docker
重启docker：systemctl restart docker
查看docker状态：systemctl status docker
开机启动：systemctl enable docker
查看docker概要信息：docker info
查看docker帮助文档：docker --help
docker创建镜像：docker build -t <name>:<tag> <Dockerfile文件所在目录>
启动docker容器：docker run -d -p <hostPort>:<containerPort> <imageID>
进入docker容器：docker exec -it <containerName>/<containerID> /bin/sh
查看当前所有镜像：docker images
删除镜像，通过镜像的id来指定删除谁：docker rmi <imageID>
删除镜像id为<None>的镜像：docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
删除全部镜像：docker rmi $(docker images -q)
查看所有容器：docker ps
查看所有容器，包括关闭的：docker ps -a
停止容器：docker stop <containerID>
启动容器：docker start <containerID>
重启容器：docker restart <containerID>
停止所有的容器：docker stop $(docker ps -a -q)
删除容器，通过容器的id来指定删除谁：docker rm <containerID>
删除所有的容器：docker rm $(docker ps -a -q)
获取容器的日志：docker logs -f <containerName>/<containerID>
文件从容器中拷贝到宿主机中：docker cp <containerID>:SRC_PATH DEST_PATH
文件从宿主机中拷贝到容器中：docker cp SRC_PATH <containerID>:DEST_PATH
查看容器中运行的进程信息：docker top <containerName>/<containerID>
列出指定的容器的端口映射：docker port <containerName>/<containerID>
```

### GoFrame gf 相关命令
```
安装gf工具：
wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
gf创建项目：
gf init projectName
go clean -modcache
rm go.sum
go mod tidy
gf编译：gf build -o main
gf资源打包：gf pack resource,manifest/config internal/packed/data.go -n=packed
```
