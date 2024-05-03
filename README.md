# go-blog
## 项目目录结构
### api
    
接口的目录
1. settings_api     &emsp;负责配置文件的接口
    1. settings_info.go     &emsp;&emsp;实现了对配置文件信息的获取
    2. settings_update.go   &emsp;&emsp;实现了对配置文件信息的更新
2. images_api       &emsp;负责图片信息的接口
    1. image_list           &emsp;&emsp;获取图片信息的接口，可以分页返回
    2. image_remove         &emsp;&emsp;删除图片的接口
    3. iamge_update         &emsp;&emsp;更新图片名称的接口
    4. iamge_uoload         &emsp;&emsp;上传图片的接口，可以实现本地上传或者上传到七牛云
    5. image_name_list      &emsp;&emsp;返回图片名称列表
3. adverts_api      &emsp;负责广告的相关接口
    1. advert_create        &emsp;&emsp;创建广告
    2. advert_list          &emsp;&emsp;返回广告列表
    3. advert_remove        &emsp;&emsp;删除广告
    4. advert_update        &emsp;&emsp;更新广告信息
4. menu_api         &emsp;菜单相关的接口
    1. menu_create          &emsp;&emsp;创建菜单
    2. menu_detail          &emsp;&emsp;菜单的细节
    3. menu_list            &emsp;&emsp;返回菜单的列表
    4. menu_name_list       &emsp;&emsp;返回菜单名称列表
    5. menu_remove          &emsp;&emsp;删除菜单
    6. menu_update          &emsp;&emsp;更新菜单信息
5. tag_api          &emsp;文章标签相关的接口
    1. tag_create           &emsp;&emsp;创建标签
    2. tag_list             &emsp;&emsp;返回标签列表
    3. tag_remove           &emsp;&emsp;删除标签
    4. tag_update           &emsp;&emsp;更新标签信息
6. user_api         &emsp;用户相关的接口
    1. email_login          &emsp;&emsp;使用邮箱登录
    2. user_bind_email      &emsp;&emsp;绑定邮箱的操作，期间需要用到邮箱发送验证码，目前还没有实现
    3. user_create          &emsp;&emsp;创建用户
    4. user_list            &emsp;&emsp;返回用户列表
    5. user_logout          &emsp;&emsp;注销操作
    6. user_remove          &emsp;&emsp;删除用户
    7. user_update_password &emsp;&emsp;更新用户密码
    8. user_update_role     &emsp;&emsp;更新用户权限  
目前qq登录和email绑定还没有做，因为一些限制性的原因
### config

存放记录配置的结构体目录，主要是根据配置文件的结构信息，创建的结构体，直接用来序列化读取的，主要分为三个结构体：
1. logger： &emsp;&emsp;用来加载日志相关的参数来初始化log对象，包括一些日志的等级等信息
2. mysql：  &emsp;&emsp;是用来加载MySQL相关的信息来初始化mysql连接对象的这里是gorm，主要包括地址和连接数等信息
3. system： &emsp;&emsp;是用来加载系统相关的信息，比如服务监听的地址和端口
4. email：  &emsp;&emsp;是用来存储email相关信息的结构体
5. jwt：    &emsp;&emsp;token相关的东西
6. qiniu：  &emsp;&emsp;是用来存储七牛云相关信息的
7. qq：     &emsp;&emsp;是用来存储qq相关信息的
8. upload： &emsp;&emsp;用来存储上传图片的路径和大小限制的
9. redis：  &emsp;&emsp;用来存储连接redis的一些信息  
其中enter.go是包含了上述三个结构体的整体结构体。
### core

初始化操作
1. 配置文件信息
2. mysql连接
3. 日志的初始化
4. redis的初始化
这里有个顺序，依次是配置文件和日志，这两个需要先初始化
### docs

swag生成的api文档目录

已完成api接口文档：
1. 广告管理
2. 图片管理
3. 菜单管理
4. 配置管理
5. 用户管理
6. 标签管理
### flag

命令行相关的初始化
1. 数据库迁移   &emsp;&emsp;-db
2. 创建用户     &emsp;&emsp;-u [admin/user]
### global

全局变量的包，主要包括一些全局变量，包括mysql连接对象，日志对象，配置文件等，新添加了redis连接对象
1. Config   &emsp;&emsp;配置文件相关的
2. DB       &emsp;&emsp;数据库相关的
3. Log      &emsp;&emsp;日志相关的
4. MysqlLog &emsp;&emsp;mysql数据库相关的
5. Redis    &emsp;&emsp;redis相关的
### middleware

中间件
1. jwt_auth token验证的中间件，分为普通用户验证和管理员验证
### plugins

包含第三方的模块
1. qiniu &emsp;&emsp;是对接七牛云接口的代码
2. email &emsp;&emsp;目前是对接qq邮箱的代码，由于一些原因，目前还没有测试
### models

这里主要是mysql数据表的信息，还有一些类型，和返回响应的封装。这里目前有一个res文件夹，主要是封装向客户端回写数据的操作
### routers

gin路由的目录，主要是设置路由的信息。
1. advert_router    &emsp;&emsp;广告相关的路由
2. images_router    &emsp;&emsp;图片相关的路由
3. menu_router      &emsp;&emsp;菜单相关的路由
4. settings_router  &emsp;&emsp;配置相关的路由
5. tag_router       &emsp;&emsp;标签相关的路由
6. user_router      &emsp;&emsp;用户相关的路由
### service

项目与服务有关的项目
1. common       &emsp;&emsp;一些通用的接口
2. image_ser    &emsp;&emsp;图片相关的接口
3. redis_ser    &emsp;&emsp;redis相关的接口
4. user_ser     &emsp;&emsp;用户相关的接口
### testdata

测试文件的目录
### utils
    
常用的一些工具
1. desens   &emsp;&emsp;数据脱敏的处理
2. jwts     &emsp;&emsp;token相关的处理
3. pwd      &emsp;&emsp;密码相关的处理
4. random   &emsp;&emsp;随机数相关的处理
5. md5      &emsp;&emsp;md5相关的处理
6. utils    &emsp;&emsp;一些工具
7. valid    &emsp;&emsp;获取有效的信息
### main.go

程序的入口
### setting.yaml

配置文件

### air

修改项目文件后，项目自动重启工具 air