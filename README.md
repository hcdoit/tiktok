# 微抖音

## 简介

一个用`kitex`和`hertz`构建，根据不同场景采用`MySQL`、`Minio`、`MongoDB`、`Redis`存储，实现了上述方案说明的基础、互动、社交接口的，划分为网关、用户、视频、社交、交互五个微服务的服务端项目。

| 服务名 | 功能                | 框架   | 协议 | 路径     | IDL             |
|--------------|----------------------|-------------|----------|----------|-----------------|
| 网关      | 接收HTTP请求       | kitex/hertz | http     | cmd/api  | idl/api.thrift  |
| 用户     | 注册、登录、用户信息 | kitex/gorm  | thrift   | cmd/user | idl/user.thrift |
| 视频     | 上传视频、视频信息、视频流 | kitex/gorm  | thrift   | cmd/video | idl/video.thrift |
| 交互     | 点赞、评论 | kitex/gorm  | thrift   | cmd/interact | idl/interact.thrift |
| 社交     | 关注、聊天 | kitex/gorm  | thrift   | cmd/social | idl/social.thrift |

## 详细文档地址
[9015-手眼通天队 青训营后端结业项目答辩汇报文档](https://is8ljs2kps.feishu.cn/docx/IubpdY3uEoltfUxhoBec8ltSndc?from=from_copylink)

## 运行指南
### 环境

```shell
docker-compose up
```

### 运行用户服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
### 运行视频服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
### 运行社交服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
### 运行交互服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
### 运行用户服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 运行网关服务

```shell
cd cmd/api
sh run.sh
```

注意：服务地址、端口需在`pkg/const`及`docker-compose.yaml`中视情况修改配置

## 待优化事项
* 引入消息队列中间件：以异步的方式实现对视频文件处理和保存以及保证缓存与数据库的一致性
* 优化微服务拆分：避免强耦合和循环依赖