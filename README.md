# Simple mock server

## 构建

```bash
go build -tags release -a -installsuffix cgo -o mockserver ./cmd
```

## 运行

```bash
./mockserver
```

默认端口为80，可以通过`-port`参数指定端口：

```bash
./mockserver -port 8080
```