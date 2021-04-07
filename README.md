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

