# go-blog
## 项目目录结构
### api
    
接口的目录
1. settings_api 负责配置文件的接口
    1. settings_info.go 实现了对配置文件信息的获取
    2. settings_update.go 实现了对配置文件信息的更新
2. images_api 负责图片信息的接口
    1. image_list 获取图片信息的接口，可以分页返回
    2. image_remove 删除图片的接口
    3. iamge_update 更新图片名称的接口
    4. iamge_uoload 上传图片的接口，可以实现本地上传或者上传到七牛云
### config

存放记录配置的结构体目录，主要是根据配置文件的结构信息，创建的结构体，直接用来序列化读取的，主要分为三个结构体：
1. logger：用来加载日志相关的参数来初始化log对象，包括一些日志的等级等信息
2. mysql：是用来加载MySQL相关的信息来初始化mysql连接对象的这里是gorm，主要包括地址和连接数等信息
3. system：是用来加载系统相关的信息，比如服务监听的地址和端口
4. email：是用来存储email相关信息的结构体
5. jwt：
6. qiniu：是用来存储七牛云相关信息的
7. qq：是用来存储qq相关信息的
8. upload：用来存储上传图片的路径和大小限制的
其中enter.go是包含了上述三个结构体的整体结构体。
### core

初始化操作，主要是对配置文件信息、mysql连接和日志的初始化，这里有个顺序，依次是配置文件、日志和mysql连接，因为有依赖关系。
### docs

swag生成的api文档目录
### flag

命令行相关的初始化，目前有一个数据库迁移
### global

全局变量的包，主要包括一些全局变量，包括mysql连接对象，日志对象，配置文件等
### middleware

中间件
### models

这里主要是mysql数据表的信息，还有一些类型，和返回响应的封装。这里目前有一个res文件夹，主要是封装向客户端回写数据的操作
### routers

gin路由的目录，主要是设置路由的信息。
### service

项目与服务有关的项目
### testdata

测试文件的目录
### utils
    
常用的一些工具
### main.go

程序的入口
### setting.yaml

配置文件