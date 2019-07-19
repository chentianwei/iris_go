# iris_go
go Iris 框架的demo

项目下目录结构：
 - controllers
    - controller.go
    - otherController.go
 - models
    - mysql_con.go
    - otherModel.go
 - routers
    - web.go
 - util
    - log.go
    - readEnv.go
 - main.go
 - .env
 
 还有一个.env文件和main.go同级
 
 
 controllers目录下是控制层，包含基础基本返回函数
 models目录是数据库的model层，包含数据库连接
 util目录是工具层，包含读取.env文件配置和日志输出文件
 routers是单独的路由文件，编辑路由和对应的func
 
