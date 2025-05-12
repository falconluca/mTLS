package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var (
		caFile   = flag.String("cacert", "../certs/ca.pem", "CA cert")
		certFile = flag.String("cert", "../certs/client.pem", "client cert")
		keyFile  = flag.String("key", "../certs/client-key.pem", "client key")
		server   = flag.String("server", "localhost:8443", "server address")
		msg      = flag.String("msg", "PING", "message to send")
	)
	flag.Parse()

	// 加载客户端证书
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("load client cert error: %v", err)
	}

	caCert, err := os.ReadFile(*caFile)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	tlsCfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: false,
	}

	conn, err := tls.Dial("tcp", *server, tlsCfg)
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(*msg))
	if err != nil {
		log.Fatalf("write error: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err == io.EOF {
		log.Println("server closed connection")
	} else if err != nil {
		log.Fatalf("read error: %v", err)
	}
	log.Printf("Received: %s", string(buf[:n]))
}
