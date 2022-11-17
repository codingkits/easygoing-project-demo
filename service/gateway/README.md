<h1 align="center">KrakenD Api Gateway</h1>

#### 项目介绍
* 【导读】：`KrakenD Api Gateway` Api网关系统 
* 其他：略。

#### 软件版本号
* [![](https://img.shields.io/badge/golang-1.19-blue)](https://golang.org)

#### 构建平台
* MacOs 12.5

####  测试地址

#### 插件构建脚本
```shell
go build  -buildmode=plugin -o krakend-auth-plugin.so .
```

#### KrakenD 运行脚本
```shell
krakend run -dc ./deploy/krakend/krakend.json
```

#### 核心依赖
|    服务    |              技术方案              |
| :--------: | :--------------------------------: |
|  注册中心  |                 //                 |
|  配置中心  |                 //                 |
|  消息队列  | // |
|  灰度分流  |                 //                 |
|  动态网关  |                 KrakenD                 |
|  授权认证  |                 自定义签名                |
|  服务容错  |                 //                 |
|  服务调用  |                 //                 |
|  任务调度  |                 //                 |
|   数据库   |               //                |
|  配置加密  |                 //                 |
|    缓存    |               //                |
| 操作数据库 |                 //                 |
|  字段映射  |                 //                 |

#### 注意事项
```shell
# 1.插件二进制文件不是跨平台兼容的，必须使用相同架构/平台来编译插件。
# 2.插件库要和主工程依赖的库版本一致
```

#### 记录
```shell
# 1.
```

#### 项目参与者
- 维护者：董健 d.joeyana@gmail.com
- 整理日期 :  @date 2022年10月19日20:43:27
- 其他：略。

