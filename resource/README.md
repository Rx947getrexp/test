代码结构
``` lua
web
├── api/v1 -- 主要API
|   ├── sys_initdb.go -- ico
|   └── sys_user.go --  
├── cmd -- 可执行文件
├── config -- 配置文件 设定操作的结构体
├── global -- global
├── initialize -- 初始化global包的工具
├── middleware -- 中间件
├── model -- global
│   ├── request  -- 所有请求model结构体
|   |   ├── common.go 
|   |   ├── ...
|   ├── response  -- 返回数据
|   |   ├── common.go 
|   |   ├── ...
├── router -- 路由
├── service -- service层
├── utils
├── config.yaml  -- 
├── go.mod    -- mod 配置
├── go.sum -- sum
├── goxorm.yaml -- 数据库表生成struct配置
└── main.go  -- main函数
```

生成数据库model：

1.首次使用要安装xorm.io/reverse

2.运行reverse -f goxorm.yaml


