## 设计说明

### 消息中间件

[NSQ](https://nsq.io/)

NSQ是一个基于Go语言的分布式实时消息平台。

#### NSQ工具组件

- nsqd：一个负责接收、排队、转发消息到客户端的守护进程
- nsqlookupd：管理拓扑信息并提供最终一致性的发现服务的守护进程
- nsqadmin：一套Web用户界面，可实时查看集群的统计数据和执行各种各样的管理任务

#### topic和channel

![](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

1. 生产者往topic发消息时，topic会把消息copy到每个channel
2. 在channel中的每条消息会被放进队列中, 直到消息被consumer所消费掉, 如果队列占用的内存超出限制, 消息会被写进硬盘

### 依赖

- golang1.16+
- NSQ

### Usage

```
go run consumer/consumer.go

#in other terminal
go run producer/producer.go
```