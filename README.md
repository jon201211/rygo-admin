## ruoyi go框架简介

本框架以ruoyi做前端，用golang Gin为web服务框架,xorm为数据库orm框架，继续沿用MIT开源协议，
架构思路沿袭着若依的以辅助生成重复代码为主，不过度封装，生成的代码可以快速修改适应不同的需求，适应每个开发者自己的习惯和风格。

> 1. 单体架构、前端为模板引擎，非vue, 适合后端开发人员。
> 2. 方便分布式部署,去掉了session组件,基于jwt token实现服务无状态化
> 3. 适合做为 k8s 微服务使用，每个微服务的前后端均在一个服务中，可通过一个portal聚合，没有vue单页面应用那种过度臃肿问题


## 核心技术及组件
> web服务框架    github.com/gin-gonic/gin v1.6.1
>
> ORM框架       github.com/go-xorm/xorm v0.7.9
>
>session       github.com/gorilla/sessions v1.2.0
>
>cache         github.com/patrickmn/go-cache v2.1.0+incompatible
>
> 配置文档       github.com/BurntSushi/toml v0.3.1
>
> 导出excel文件  tealeg/xlsx    v1.0.5   
>
> api文档生成    swaggo/swag    v1.6.5   
>
> 图形验证码     base64Captcha  v1.2.2  
>
> 服务器监控     gopsutil       v2.19.12+incompatible   
>
> 若依前端组件   RuoYi           v4.1.0


## 内置功能
1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  通知公告：系统通知公告信息发布维护。
9.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10.  登录日志：系统登录日志记录查询包含登录异常。
11.  在线用户：当前系统中活跃用户状态监控。
12.  定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13.  代码生成：前后端代码的生成（Go、html、json、sql） 。
14.  系统接口：根据业务代码自动生成相关的api接口文档。
15.  服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
16.  在线构建器：拖动表单元素生成相应的HTML代码。
17.  案例演示：常用的前端组件整合演示。


## 数据库
    可以选用mysql或者sqlite数据库，方便调试。
> mysql
   mysql -u root -p  < ./data/mysql.sql
> sqlite3
    .read ./data/sqlite3.sql    


## 编译测试
> go build & ./rygo
> 登录http://127.0.0.1:8080/
> 账号: admin / admin123 

## 开发计划
> 有时间持续更新基础功能，欢迎交流


## 感谢(排名不分先后)
> gin框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) 
>
> ORM框架  [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)      
>
>cache    [https://github.com/patrickmn/go-cache](https://github.com/patrickmn/go-cache)
>
> 配置文档   [https://github.com/BurntSushi/toml](https://github.com/BurntSushi/toml)
>
> RuoYi框架 [https://github.com/yangzongzhuan/RuoYi](https://github.com/yangzongzhuan/RuoYi)
>
> tealeg [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
>
> swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)
>
> 

## 交流
> mail:583173@qq.com