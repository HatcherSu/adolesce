# Go 回调工具

## 目录结构

```
./
├── /api                             # http,grpc目录的api目录
│   ├── /public                          # 存放页面文件
│   ├── api_router.go                   # 配置路由
├── /bin                             # 编译的二进制文件存放位置
├── /cmd                             # 应用入口目录
│   ├── wire.go                         # 生成依赖注入
│   └── main.go                         # main函数
├── /internal                        # 封装内部实现
├── /scripts                         # 脚本
│   ├── /biz                             # usecase
│   ├── /conf                            # 读取配置目录
│   ├── /data                            # 仓库
│   ├── /pkg                             # 封装工具包
│   ├── /server                          # http、grpc、定时器等服务实例
│   └── /service                         # service
├── /storage                         # 运行数据存储目录
│   └── logs                            # 日志存储目录
├── /third_party                     # 第三方工具
├── go.mod                          # go mod
├── go.sum                          # go sum
├── Makefile                        # 编译脚本
└── README.md                       # 文档说明
```

## 使用说明

* 前提，需要配置好go环境

1. 把相关项目`clone`到自己目录下
```shell
git clone 
```

2. 同步包到本地
```shell
go mod tidy
```

3. 设置全局环境 `.env`
 - 主要需要设置的配置如下

```
# 数据库
DB_HOST=127.0.0.1       # 数据库host
DB_PORT=3306            # 数据库端口
DB_USERNAME=root        # 用户名
DB_PASSWORD=root1234    # 密码
DB_DATABASE=db_ex_user  # 需要连接的数据库名

# HTTP配置
HTTP_PORT=8081          # http端口
HTTP_IP_ADDR=localhost  # IP地址
```

4. 修改html中相关url, 修改`localhost:8081`，修改成配置相应路由`ip:port`
5. 执行sql文件，创建相应的表`/cloud_callback/internal/data/schema.sql`
6. `make install` -- 安装相关的包以及编译
7. `make run`或者`make all` -- 运行
7. 浏览url`ip:port/web/index.html`
