---
title: Docker命令

date: 2021-11-01	

categories: Docker	

tags: [Docker]
---	

# 常用命令

## 新建+启动容器

docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

OPTIONS说明（常用）：有些是一个减号，有些是两个减号

--name="容器新名字"    为容器指定一个名称；

-d: 后台运行容器并返回容器ID，也即启动守护式容器(后台运行)；

-i：以交互模式运行容器，通常与 -t 同时使用；

-t：为容器重新分配一个伪输入终端，通常与 -i 同时使用；

-P: 随机端口映射，大写P

-p: 指定端口映射，小写p

![image-20230823095858023](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823095858023.png) 

使用镜像centos:latest以交互模式启动一个容器,在容器内执行/bin/bash命令。

docker run -it centos /bin/bash

## 容器管理

**退出容器：**

![image-20230823100125508](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823100125508.png) 

**启动停止的容器：**![image-20230823100154247](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823100154247.png) 

**重启：**restart

**停止：**stop

**强制停止：**kill

**删除已停止：**rm 容器ID (一次性删除多个：![image-20230823100515195](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823100515195.png))

**日志：**logs 容器ID

**容器呢i运行的进程：**top 容器ID

**容器内部细节：**inspect 容器ID

**拷贝文件到主机：**cp 容器ID:容器内路径 目的主机路径

**导入和导出容器：**![image-20230823101344450](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823101344450.png)

**提交容器副本使之成为一个新的镜像：**docker commit -m="提交的描述信息"-a="作者"容器ID 要创建的目标镜像名:[标签名]

## 进入正在运行的容器并以命令行交互

![image-20230823100942131](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823100942131.png)

## 容器卷

![image-20230823101940115](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823101940115.png)

![image-20230823102211961](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823102211961.png) 

# DockerFile

## Dokcer执行Dockerfile的大致过程

![image-20230823140711653](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823140711653.png) 

## 常用保留字指令

<img src="/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Docker/img/image-20230823140938785.png" alt="image-20230823140938785" style="zoom:150%;" />


构建：docker build -t 新镜像名字:TAG . (最后面有个点)

# Docker-compouse容器编排

Compose 是 Docker 公司推出的一个工具软件，可以管理多个 Docker 容器组成一个应用。你需要定义一个 YAML 格式的配置文件docker-compose.yml，写好多个容器之间的调用关系。然后，只要一个命令，就能同时启动/关闭这些容器

## Compose常用命令

docker-compose -h              # 查看帮助

docker-compose up              # 启动所有docker-compose服务

docker-compose up -d             # 启动所有docker-compose服务并后台运行

docker-compose down             # 停止并删除容器、网络、卷、镜像。

docker-compose exec  yml里面的服务id         # 进入容器实例内部 docker-compose exec docker-compose.yml文件中写的服务id /bin/bash

docker-compose ps            # 展示当前docker-compose编排过的运行的所有容器

docker-compose top           # 展示当前docker-compose编排过的容器进程

docker-compose logs  yml里面的服务id   # 查看容器输出日志

docker-compose config   # 检查配置

docker-compose config -q # 检查配置，有问题才有输出

docker-compose restart  # 重启服务

docker-compose start   # 启动服务

docker-compose stop    # 停止服务
