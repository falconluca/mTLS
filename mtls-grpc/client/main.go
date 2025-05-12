package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "mTLS/mtls-grpc/proto/ping"
)

func main() {
	var (
		caFile   = flag.String("cacert", "../certs/ca.pem", "CA cert")
		certFile = flag.String("cert", "../certs/client.pem", "client cert")
		keyFile  = flag.String("key", "../certs/client-key.pem", "client key")
		server   = flag.String("server", "localhost:8443", "gRPC server address")
		msg      = flag.String("msg", "hello from client", "message to send")
	)
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("load client cert error: %v", err)
	}

	// Load CA
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: false,
	})

	conn, err := grpc.Dial(*server, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Ping(ctx, &pb.PingRequest{Message: *msg})
	if err != nil {
		log.Fatalf("error calling Ping: %v", err)
	}
	log.Printf("Response: %s", res.Reply)
}
