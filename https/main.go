package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"
)

// ---------- 工具函数部分 ----------

// 生成 RSA 密钥对
func generateServerKeyPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// RSA 加密
func encryptSymmetricKey(pub *rsa.PublicKey, key []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, key)
}

// RSA 解密
func decryptSymmetricKey(priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// AES 加密
func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	plaintext = pkcs7Pad(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, block.BlockSize())
	rand.Read(iv)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return append(iv, ciphertext...), nil
}

// AES 解密
func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	return pkcs7Unpad(plaintext)
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	padlen := int(data[length-1])
	return data[:length-padlen], nil
}

// ---------- 模拟 CA 签发 ----------

func createCACert() (*x509.Certificate, *rsa.PrivateKey, []byte) {
	caKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	caTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(2025),
		Subject:               pkix.Name{CommonName: "MyRootCA"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	caCertDER, _ := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caKey.PublicKey, caKey)
	return caTemplate, caKey, caCertDER
}

func createServerCert(caCert *x509.Certificate, caKey *rsa.PrivateKey, serverPub *rsa.PublicKey) ([]byte, error) {
	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2048),
		Subject:      pkix.Name{CommonName: "myserver.com"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
	}
	return x509.CreateCertificate(rand.Reader, serverTemplate, caCert, serverPub, caKey)
}

// 验证服务器证书（客户端做的）
func verifyServerCert(certDER []byte, caCertDER []byte) (*rsa.PublicKey, error) {
	// 解析服务器证书
	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, fmt.Errorf("解析服务器证书失败：%w", err)
	}

	// 解析 CA 证书
	caCert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		return nil, fmt.Errorf("解析 CA 证书失败：%w", err)
	}

	// 创建 CertPool，并加入 CA 证书
	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	// 验证证书链
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		return nil, fmt.Errorf("❌ 证书验证失败：%w", err)
	}

	// 提取服务器公钥
	pubKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("证书不是 RSA 公钥")
	}
	return pubKey, nil
}

// 模拟 HTTPS 核心流程
func main() {
	// 生成 CA 证书
	caCert, caKey, caCertDER := createCACert()

	// 生成服务器密钥对
	serverPrivKey, _ := generateServerKeyPair()

	// 签发服务器证书
	serverCertDER, _ := createServerCert(caCert, caKey, &serverPrivKey.PublicKey)

	// 客户端验证服务器证书
	serverPubKey, err := verifyServerCert(serverCertDER, caCertDER)
	if err != nil {
		panic(err)
	}

	// ------------------- 阶段1：身份认证 -------------------

	// 客户端生成对称密钥
	symmetricKey := make([]byte, 32)
	rand.Read(symmetricKey)

	// 客户端使用服务器的公钥加密客户端生成的对称密钥
	encryptedKey, _ := encryptSymmetricKey(serverPubKey, symmetricKey)

	// 服务器使用私钥解密，获取客户端生成的对称密钥
	decryptedKey, _ := decryptSymmetricKey(serverPrivKey, encryptedKey)

	// ------------------- 阶段2：对称加密传输 -------------------

	// 客户端传输加密数据给服务器
	plaintext := "我是一条 HTTPS 里的小秘密💬"
	ciphertext, _ := encryptAES(decryptedKey, []byte(plaintext))
	fmt.Println("加密后的内容：", ciphertext)

	// 服务器用解密并处理数据
	decryptedText, _ := decryptAES(decryptedKey, ciphertext)
	fmt.Println("解密后的内容：", string(decryptedText))
}
