## 书籍信息管理系统

- 采用grpc实现
- [data.json](server/data.json)实现数据存储

### Dependence

`golang 1.6+ or depoly on docker with Dockerfile`

### Server

```bash
go run server/server.go [-port=${port}]
```

### Client

```bash
go run client/client.go
```

1. 添加图书：	`add [id] [name]`
2. 按id查询：	`qbi [id]`
3. 按name查询：	`qbn [name]`
4. 删除图书：	`del [id]`
5. 展示所有图书 `show`
6. 退出        `quit`