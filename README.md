## 分布式计算

### [Socket](./SocketTest)

### [homework](./homework)

#### [task1](./homework/task1)

- 题目1：将基于UDP协议的Client-Server通信程序示例的服务器端程序改造成多线程版。
- 题目2：将基于TCP协议的Client-Server通信程序示例的服务器端程序改造成线程池版。

#### [task2](./homework/task2)

利用RPC技术实现一个书籍信息管理系统，具体要求:
1. 客户端实现用户交互，服务器端实现书籍信息存储和管理。客户端与服务器端利用RPC机制进行通信。可以选择Java RMI、gRPC、Dubbo等任意RPC中间件。
2. 服务器端至少暴露如下RPC接口
    - bool add(Book b)   添加一个书籍对象。（注意Book对象序列化问题）
    - Book queryByID(int bookID) 查询指定ID号的书籍对象。
    - BookList queryByName(String name) 按书名查询书籍对象列表。
    - bool delete((int bookID) 删除指定ID号的书籍对象。


#### [task3](./homework/task3)

利用MOM消息队列技术实现一个分布式随机信号分析系统，具体要求：
- 随机信号产生器每隔10毫秒左右就产生一个正态分布的随机数字，并作为一个消息发布
- 多个随机信号分析模块订阅并接收该随机数字，然后对信号进行分析并实时显示分析结果。至少包含如下分析模块：
    1. 计算随机信号的均值；
    2. 计算过去N个随机信号的方差（N为常量，可设置）
    3. 实现基于正态分布的异常点检测
    4. 实时绘制过去一段时间内随机信号的折线图（选作

#### [task4](./homework/task4)

1. 自己查阅资料，说明三类分布式存储系统的区别：
    1. 块存储系统;
    2. 对象存储系统
    3. 文件存储系统。
2. 阅读论文《The Hadoop Distributed File System》并回答下面问题：
    - 客户端读取HDFS系统中指定文件指定偏移量处的数据时，工作流程是什么？
    - 客户端向HDFS系统中指定文件追加写入数据的工作流程是什么？
    - 新增加一个数据块时，HDFS如何选择存储该数据块的物理节点？
    - HDFS采用了哪些措施应对数据块损坏或丢失问题？
    - HDFS采用了什么措施应对主节点失效问题？
    - NameNode维护的“数据块—物理节点对应表”需不需要在硬盘中备份？为什么？

#### [task5](./homework/task5)
