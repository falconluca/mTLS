package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// CLI 解析（模拟 etcdctl 的风格）
	caPath := flag.String("cacert", "../certs/ca.pem", "CA cert path")
	certPath := flag.String("cert", "../certs/client.pem", "client cert path")
	keyPath := flag.String("key", "../certs/client-key.pem", "client key path")
	server := flag.String("server", "https://localhost:8443/ping", "server endpoint")
	flag.Parse()

	// 加载客户端证书
	cert, err := tls.LoadX509KeyPair(*certPath, *keyPath)
	if err != nil {
		log.Fatalf("load client cert error: %v", err)
	}

	// 加载 CA
	caCert, err := ioutil.ReadFile(*caPath)
	if err != nil {
		log.Fatalf("read ca.pem error: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	// 创建 HTTPS 客户端
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: false,
	}

	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}

	resp, err := client.Get(*server)
	if err != nil {
		log.Fatalf("GET request error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response from server: %s", string(body))
}
