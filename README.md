# Microservice
一个练习微服务的项目，从0搭建一个实际应用的多服务架构

## 项目架构

    客户端(前期暂时由curl)
            |
           Nginx
            |
         backend - 业务网关层（BFF）
            |
           / \          \
        auth  service(1)...service(N) - 微服务层

1. 业务网关不一定只有一个，这里暂且只有一个，如果需要对接多个不同的业务前端或客户端平台，可以通过这个聚合微服务提供的功能
2. 微服务之间可能存在调用关系

## 框架

使用了 [jotnar](https://github.com/paulyung541/jotnar) 快速搭建项目