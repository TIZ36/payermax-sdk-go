package gateway

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/tiz36/payermax-sdk-go/config"
)

func Sign(config *config.PayMaxConfig, reqStr string) (string, error) {
	// 从某处获取现有的 RSA 私钥（这里简化，实际应用中需要从安全存储中加载私钥）
	privateKey, _ := ParsePKCS8PrivateKey([]byte(config.MerchantPrivateKey))

	// 要签名的数据
	data := []byte(reqStr)

	// 计算 SHA-256 散列
	hash := sha256.Sum256(data)

	// 对散列值进行签名
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Println("签名时发生错误:", err)
		return "", err
	}

	// 将签名结果转换为 base64 字符串
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("签名:", signatureBase64)

	return signatureBase64, nil
}

func Verify(data string, sign string, config *config.PayMaxConfig) (bool, error) {
	// 从某处获取现有的 RSA 公钥（这里简化，实际应用中需要从安全存储中加载公钥）
	publicKey, _ := ParsePublicKey([]byte(config.PayMaxPublicKey))

	// 要验证的数据
	dbin := []byte(data)

	// 对签名进行 base64 解码
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("签名解码时发生错误:", err)
		return false, err
	}

	// 计算 SHA-256 散列
	hash := sha256.Sum256([]byte(dbin))

	// 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("签名验证失败:", err)
		return false, err
	}

	fmt.Println("签名验证成功!")

	return true, nil
}

func ParsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("private key error")
	}

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, err
}

func ParsePublicKey(data []byte) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key error")
	}

	return key, err
}

func ParsePKCS8PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var priInterface interface{}
	priInterface, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := priInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key error")
	}

	return key, err
}
