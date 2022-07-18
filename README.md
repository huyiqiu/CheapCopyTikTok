# 极简抖音项目

## 项目简介
本项目为极简版抖音的后端项目

演示视频：
[演示视频地址](img/tiktokdemo.mp4)  
API文档：[抖音极简版](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)  
测试工具：[apk文件](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#quPkfu) + postman

web框架：gin  
数据库：MySQL  
中间件：Redis, 七牛云 


## 项目架构
本项目c采用`MVC三层架构`  
controller层负责响应接口  
service层负责处理业务逻辑  
repository层负责数据库交互   
![架构](img/%E9%A1%B9%E7%9B%AE%E6%9E%B6%E6%9E%84.jpg)

## 目录结构
```
|-- README.md
|-- controller
|   |-- commentController.go
|   |-- favoriteController.go
|   |-- feedController.go
|   |-- followController.go
|   |-- getUserInfo.go
|   |-- loginController.go
|   `-- publishVideo.go
|-- go.mod
|-- go.sum
|-- img
|   |-- controller.png
|   |-- rep.png
|   |-- service.png
|   `-- tiktokdemo.mp4
|-- main.go
|-- repository
|   |-- comment.go
|   |-- db_init.go
|   |-- favorite.go
|   |-- follow.go
|   |-- user.go
|   `-- video.go
|-- service
|   |-- commentStuff.go
|   |-- favoriteStuff.go
|   |-- followStuff.go
|   |-- publish.go
|   |-- queryFeedFlow.go
|   |-- registerAndLogin.go
|   `-- userInfo.go
`-- utils
    |-- logger.go
    `-- validate.go
```
