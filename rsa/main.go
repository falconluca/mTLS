package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// 服务端生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 显示公钥和私钥（PEM 格式）
	fmt.Println("🔐 私钥:")
	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	fmt.Println(string(privPem))

	fmt.Println("🔓 公钥:")
	pubASN1, _ := x509.MarshalPKIXPublicKey(publicKey)
	pubPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})
	fmt.Println(string(pubPem))

	// 客户端用公钥加密
	message := []byte("今晚星星很亮，我想你会来。🌙")
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("🔒 加密后的密文: %x\n", cipherText)

	// 服务端用私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		panic(err)
	}
	fmt.Printf("💌 解密后的明文: %s\n", string(plainText))
}
