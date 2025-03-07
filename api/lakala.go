package api

import (
	"bytes"
	"crypto"
	randc "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/sleep-go/lakala-pay/util"
	"github.com/tjfoc/gmsm/sm4"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	HOST_TEST = "https://test.wsmsd.cn/sit"
	HOST      = "https://s2.lakala.com"
	algorism  = "LKLAPI-SHA256withRSA"
)

type Client struct {
	algorism       string
	appid          string
	serialNo       string
	timestamp      int64
	privateKeyPath string
	publicCertPath string
	IsProd         bool
	Host           string
	Http           *http.Client
}

func NewClient(appid, serialNo, path, certPath string, prod bool) *Client {
	host := HOST_TEST
	if prod {
		host = HOST
	}
	return &Client{
		algorism:       algorism,
		appid:          appid,
		serialNo:       serialNo,
		timestamp:      time.Now().Unix(),
		privateKeyPath: path,
		publicCertPath: certPath,
		IsProd:         prod,
		Host:           host,
		Http:           http.DefaultClient,
	}
}

// GetAuthorization 生成签名
func (c *Client) GetAuthorization(body []byte) (string, error) {
	//body2, _ := base64.StdEncoding.DecodeString(string(body))
	//body = []byte(body2)

	//body2 := base64.StdEncoding.EncodeToString(body)
	//body = []byte(body2)

	nonceStr := util.RandStr(12)
	message := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", c.appid, c.serialNo, c.timestamp, nonceStr, body)
	privateKey, err := loadPrivateKey(c.privateKeyPath)
	if err != nil {
		log.Println("Failed to load private key:", err)
		return "", err
	}

	signature, err := signMessage(message, privateKey)
	if err != nil {
		log.Println("Failed to sign message:", err)
		return "", err
	}
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	//fmt.Println(signatureBase64)
	//
	//signature2, err := signMessage2(message, privateKey)
	//if err != nil {
	//	return "", err
	//}
	//signatureBase642 := base64.StdEncoding.EncodeToString(signature2)
	//fmt.Println(signatureBase642)

	sign := fmt.Sprintf(`%s appid="%s",serial_no="%s",timestamp="%d",nonce_str="%s",signature="%s"`, algorism, c.appid, c.serialNo, c.timestamp, nonceStr, signatureBase64)
	return sign, nil
}

// SignatureVerification 验签
func (c *Client) SignatureVerification(authorization, body string) bool {
	// 删除方案部分并替换逗号为&
	authorization = strings.Replace(authorization, algorism+" ", "", 1)
	authorization = strings.ReplaceAll(authorization, ",", "&")
	authorization = strings.ReplaceAll(authorization, `"`, "")
	// 将查询字符串解析为map
	authorizationMap, err := url.ParseQuery(authorization)
	if err != nil {
		log.Println("Failed to parse authorization:", err)
		return false
	}
	signStr := authorizationMap.Get("signature")
	signStr = strings.ReplaceAll(authorizationMap.Get("signature"), " ", "+")
	// 解码签名为字节切片
	signature, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		log.Println("Failed to decode signature:", err)
		return false
	}
	// 构造消息
	message := fmt.Sprintf("%s\n%s\n%s\n", authorizationMap.Get("timestamp"), authorizationMap.Get("nonce_str"), body)
	publicKey, err := loadPublicKey(c.publicCertPath)
	if err != nil {
		log.Println("loadPublicKey err:", err)
		return false
	}
	// 验证签名
	return verifyMessage(message, publicKey, signature)
}

// 验签，成功返回请求body
func (c *Client) VerifySign(r *http.Request) (string, error) {
	appid := r.Header.Get("Lklapi_Appid")
	serialNo := r.Header.Get("Lklapi_Serial")
	timestamp := r.Header.Get("Lklapi_Timestamp")
	nonce := r.Header.Get("Lklapi_Nonce")
	sign := r.Header.Get("Lklapi_Sign")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer r.Body.Close()
	signature, _ := base64.StdEncoding.DecodeString(sign)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", appid, serialNo, timestamp, nonce, body)
	// 解析公钥证书
	lklCertificate, err := loadPublicKeyNew(c.publicCertPath)
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = lklCertificate.CheckSignature(x509.SHA256WithRSA, []byte(message), signature)
	if err != nil {
		log.Printf("verify sign error: %v \n appid: %s, serialNo: %s, ts: %s, nonce: %s, sign: %s, body: %s",
			err, appid, serialNo, timestamp, nonce, sign, body)
		return "", err
	}
	return string(body), err
}

// 消息验签
func verifyMessage(message string, publicKey *rsa.PublicKey, signature []byte) bool {
	// 计算消息的 SHA-256 散列值
	hashed := sha256.Sum256([]byte(message))
	// 验证签名
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		log.Println("Signature verification failed:", err)
		return false
	}
	return true
}

// 消息签名
func signMessage(message string, privateKey *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256([]byte(message))
	return rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
}

func signMessage2(message string, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := crypto.SHA256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)
	return rsa.SignPKCS1v15(randc.Reader, privateKey, crypto.SHA256, hashed)
}

// 加载私钥
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM format")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key type")
	}
	return key, nil
}

// 加载公钥
func loadPublicKey(path string) (*rsa.PublicKey, error) {
	// 读取证书公钥
	certificateBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("Failed to read certificate:", err)
		return nil, err
	}
	// 解析公钥
	block, _ := pem.Decode(certificateBytes)
	if block == nil {
		return nil, err
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println("Failed to parse certificate:", err)
		return nil, err
	}
	// 获取 RSA 公钥
	publicKey, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		log.Println("Invalid public key type")
		return publicKey, errors.New("invalid public key type")
	}
	return publicKey, nil
}

func loadPublicKeyNew(path string) (*x509.Certificate, error) {
	// 读取证书公钥
	certificateBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("Failed to read certificate:", err)
		return nil, err
	}
	// 解析公钥
	block, _ := pem.Decode(certificateBytes)
	if block == nil {
		return nil, err
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	return certificate, nil
}

// ECB模式加密（PKCS5Padding）
func EncryptECB(key, plaintext []byte) (string, error) {
	// 检查密钥长度
	key2, _ := base64.StdEncoding.DecodeString(string(key))
	key = key2
	if len(key) != sm4.BlockSize {
		return "", fmt.Errorf("无效密钥长度:%d，需要16字节", len(key))
	}

	// 创建密码块
	cipher, err := sm4.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("创建密码块失败: %v", err)
	}

	// PKCS5填充
	paddedText := PKCS5Padding(plaintext, sm4.BlockSize)

	// ECB模式加密
	ciphertext := make([]byte, len(paddedText))
	for i := 0; i < len(paddedText); i += sm4.BlockSize {
		cipher.Encrypt(ciphertext[i:], paddedText[i:])
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// PKCS5填充
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
