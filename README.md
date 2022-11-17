<h1 align="center">easygoing-project-demo</h1>

## 项目介绍
* 【导读】：`easygoing-project-demo v1.0` 
* 项目工程示例

## 功能列表
* 项目工程示例

## 服务拆分
* Api
    * http层api服务
* Mq
    * 消息系统
* Auth
    * 鉴权
* Gateway
    * 网关系统

## 项目结构
```shell
    ├── data                    # 数据挂载目录
    ├── deploy                  # 部署
    ├── docs                    # markdown
    ├── service
    │   ├── auth                # 鉴权服务
    │   ├── api                 # HTTP
    │   ├── mq                  # 消息
    │   ├── gateway             # 网关服务
    └── test
```
## 常见服务类型的目录结构
```shell
xxsrv
    ├── api         # http服务，业务需求实现
    ├── task        # 定时任务，处理数据更新业务等
    ├── model       # 数据操作
    ├── mq          # 消息系统
    ├── rpc         # rpc服务，给其他子系统提供基础数据访问
    └── script      # 脚本，处理一些临时运营需求，临时数据修复
```
## 服务内部分层
```shell
api/rpc/mq
    ├── etc             # 服务配置
    ├── internal
    │   ├── config
    │   ├── logic       # 业务逻辑
    │   ├── server
    │   └── svc         
    ├── transclient     # http和grpc实例的创建和配置
    └── pb              # pb文件
```
#### 项目参与者
- 维护者：董健 d.joeyana@gmail.com
- 整理日期 :  @date 2022年11月11日 星期五 17时17分23秒 CST
- 其他：略。