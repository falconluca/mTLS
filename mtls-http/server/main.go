package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	// 加载服务端证书
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("load server cert error: %v", err)
	}

	// 加载 CA，用来验证客户端证书
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	// 设置 TLS 配置
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from secure server 🛡️")
	})

	log.Println("Secure server listening on https://localhost:8443")
	err = server.ListenAndServeTLS("", "") // cert+key 已经在 TLSConfig 中指定
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
