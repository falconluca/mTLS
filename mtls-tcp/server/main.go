package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("get exe path error: %v", err)
	}

	baseDir := filepath.Dir(exePath)
	caCertPath := filepath.Join(baseDir, "../certs/ca.pem")
	certPath := filepath.Join(baseDir, "../certs/server.pem")
	keyPath := filepath.Join(baseDir, "../certs/server-key.pem")

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("load server cert error: %v", err)
	}

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caPool,
		MinVersion:   tls.VersionTLS12,
	}

	listener, err := tls.Listen("tcp", ":8443", tlsCfg)
	if err != nil {
		log.Fatalf("TLS listen error: %v", err)
	}
	defer listener.Close()
	log.Println("🔒 TCP server listening on :8443 with mTLS")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("Client connected: %v", conn.RemoteAddr())
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			log.Println("client closed connection")
			return
		}
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}
		cmd := string(buf[:n])
		log.Printf("Received: %s", cmd)

		var reply string
		switch cmd {
		case "PING":
			reply = "PONG\n"
		default:
			reply = "UNKNOWN COMMAND\n"
		}
		conn.Write([]byte(reply))
	}
}
