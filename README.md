# mTLS

演示如何通过 mTLS 构建安全通信，包括：

- [x] 使用 [cfssl](https://github.com/cloudflare/cfssl) 管理证书
- [x] HTTP + mTLS
- [x] gRPC + mTLS
- [x] TCP + mTLS

## 项目结构

```bash
.
├── bin/               # 编译产物目录
├── certs/             # cfssl 生成的证书文件
├── mtls-http/         # HTTP 版本（client/server）
├── mtls-grpc/         # gRPC 版本（client/server + proto）
├── mtls-tcp/          # TCP 原生版本（client/server）
└── Makefile           # 一键构建与清理脚本
````

## 快速开始

1. 安装依赖

```bash
brew install cfssl  # macOS
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. 生成 mTLS 证书

```bash
make certs
```

3. 构建所有二进制

```bash 
make build
```

4. HTTP + mTLS 示例


```bash
make run-http-server
make run-http-client
```

5. gRPC + mTLS 示例

```bash
make run-grpc-server
make run-grpc-client
```

6. TCP + mTLS 示例

```bash
make run-tcp-server
make run-tcp-client
```

7. 清理构建产物

```bash
make clean         
```

## 参考

* `mTLS`：双向认证，客户端和服务端都要提供证书，适合高安全场景（如 etcd、K8s apiserver）
* `cfssl`：Cloudflare 出品的轻量 CA 工具，适合本地快速签发测试证书
* gRPC 通信：通过 `.proto` 定义接口 + TLS 安全通信
* TCP mTLS：原生 net.Conn 封装 `tls.Conn`，低层实现更贴近原理
* [Configuring Your Go Server for Mutual TLS](https://smallstep.com/hello-mtls/doc/server/go)
* [cfssl 官方文档](https://github.com/cloudflare/cfssl)
* [gRPC With TLS Example](https://grpc.io/docs/guides/auth/)
