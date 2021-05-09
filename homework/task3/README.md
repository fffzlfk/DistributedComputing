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

### 设计思路

- 生产者向`topic`中采用异步方式发送随机信号
- 四个消费者订阅该topic
- 折线图绘制
    - 使用gin框架将数据以json格式推送到`http://localhost:8080`
    - 使用echarts从后端获取数据，并完成绘制

![](https://files.catbox.moe/zgwmov.png)

### Usage

#### 生产者
```
go run producer/producer.go
```

#### 计算均值

```
go run mean/mean.go #in other terminal
```

![](https://files.catbox.moe/nfeggr.png)

#### 计算方差

```
go run variance/variance.go #in other terminal
```

![](https://files.catbox.moe/uwaw72.png)

#### 异常点检测

```
go run outlier/outlier.go #in other terminal
```

![](https://files.catbox.moe/0ywaxi.png)

#### 图像绘制

```
go run chart/chart.go #in other terminal
```

打开`chart/index.html`可查看使用echarts绘制的折线图
![](https://files.catbox.moe/7o6ml6.gif)