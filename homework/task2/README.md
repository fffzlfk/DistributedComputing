## 书籍信息管理系统

- 采用grpc实现

### Dependence

`golang 1.6+ or depoly on docker with [Dockerfile](./Dockerfile)`

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