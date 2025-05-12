package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "mTLS/mtls-grpc/proto/ping"
)

type pingServer struct {
	pb.UnimplementedPingServiceServer
}

func (s *pingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %s", req.Message)
	return &pb.PingResponse{Reply: "Pong back: " + req.Message}, nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("get exe path error: %v", err)
	}

	baseDir := filepath.Dir(exePath)
	caCertPath := filepath.Join(baseDir, "../certs/ca.pem")
	certPath := filepath.Join(baseDir, "../certs/server.pem")
	keyPath := filepath.Join(baseDir, "../certs/server-key.pem")

	// 加载服务端证书
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("load server cert error: %v", err)
	}

	// 加载 CA 验证客户端
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterPingServiceServer(server, &pingServer{})

	lis, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("🚀 gRPC server listening at :8443 with mTLS")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
