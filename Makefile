BINS := mtls-http-server mtls-http-client mtls-tcp-server mtls-tcp-client mtls-grpc-server mtls-grpc-client

.PHONY: all build clean certs run-http run-tcp run-grpc gen-protoc

all: build

certs:
	@echo "🔐 Generating certificates using cfssl..."
	@cfssl gencert -initca certs/ca-csr.json | cfssljson -bare certs/ca
	@cfssl gencert \
		-ca=certs/ca.pem \
		-ca-key=certs/ca-key.pem \
		-config=certs/ca-config.json \
		-profile=server \
		certs/server-csr.json | cfssljson -bare certs/server
	@cfssl gencert \
		-ca=certs/ca.pem \
		-ca-key=certs/ca-key.pem \
		-config=certs/ca-config.json \
		-profile=client \
		certs/client-csr.json | cfssljson -bare certs/client
	@echo "✅ Certificates generated in certs/"

build:
	@echo "📦 Building all modules..."
	GOOS=darwin GOARCH=arm64 go build -o bin/mtls-http-server ./mtls-http/server/main.go
	go build -o bin/mtls-http-client ./mtls-http/client/main.go
	go build -o bin/mtls-grpc-server ./mtls-grpc/server/main.go
	go build -o bin/mtls-grpc-client ./mtls-grpc/client/main.go
	go build -o bin/mtls-tcp-server ./mtls-tcp/server/main.go
	go build -o bin/mtls-tcp-client ./mtls-tcp/client/main.go

run-http:
	@echo "🚀 Running HTTP server with mTLS..."
	./bin/mtls-http-server

run-http-client:
	./bin/mtls-http-client --cacert=certs/ca.pem --cert=certs/client.pem --key=certs/client-key.pem

gen-protoc:
	protoc --go_out=. --go-grpc_out=. ./mtls-grpc/proto/ping.proto

run-grpc:
	@echo "🚀 Running gRPC server with mTLS..."
	./bin/mtls-grpc-server

run-grpc-client:
	./bin/mtls-grpc-client --cacert=certs/ca.pem --cert=certs/client.pem --key=certs/client-key.pem --msg="hi grpc"

run-tcp:
	@echo "🚀 Running TCP server with mTLS..."
	./bin/mtls-tcp-server

run-tcp-client:
	./bin/mtls-tcp-client --cacert=certs/ca.pem --cert=certs/client.pem --key=certs/client-key.pem --msg="PING"

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf bin/
	@rm -f certs/*.pem certs/*.csr certs/*.key
