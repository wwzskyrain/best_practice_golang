package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

const (
	key = 0x1a2b3c4d
)

var (
	// 128 IV
	InitialVector = []byte{
		0x20, 0x75, 0x50, 0x35, 0xf5, 0x8c, 0x50, 0xd6,
		0xef, 0x9e, 0x3e, 0x2a, 0xf6, 0xf5, 0xf0, 0xad,
	}
	// 256 密钥
	Passphrase = []byte{
		0x67, 0xff, 0xc3, 0xed, 0x68, 0xde, 0x5f, 0x89,
		0x10, 0xe3, 0xa5, 0x22, 0x73, 0xb4, 0x5e, 0xfb,
		0x99, 0x8b, 0x12, 0x08, 0x85, 0x32, 0xc7, 0x0a,
		0xac, 0xa4, 0x1a, 0xc3, 0xaf, 0x6c, 0x79, 0xe5,
	}

	ErrInvalidTokenSize = errors.New("invalid token size")
	ErrInvalidToken     = errors.New("invalid token")
)

// AesCbcEncrypt 包名加密
// aes 可以是128 192 256 三种密钥长度
// iv 和块大小是 128
func AesCbcEncrypt(src string) (string, error) {
	block, err := aes.NewCipher(Passphrase)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, InitialVector)
	content := []byte(src)
	println("block.BlockSize()=", block.BlockSize())
	content = PKCS5Padding_(content, block.BlockSize())
	println("content_len=", len(content))
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	fmt.Printf("crypted: %+v\n", crypted)
	encryptedString := base64.StdEncoding.EncodeToString(crypted)
	return encryptedString, nil
}

// AesCbcDecrypt 包名解密
func AesCbcDecrypt(cipherTextBase64 string) (string, error) {
	cipherText, _ := base64.StdEncoding.DecodeString(cipherTextBase64)
	block, err := aes.NewCipher(Passphrase)
	if err != nil {
		return "", err
	}
	if len(cipherText) == 0 {
		return "", errors.New("plain content empty")
	}
	if len(cipherText)%aes.BlockSize != 0 {
		return "", ErrInvalidTokenSize
	}
	ecb := cipher.NewCBCDecrypter(block, InitialVector)
	decrypted := make([]byte, len(cipherText))
	ecb.CryptBlocks(decrypted, cipherText)
	decryptedWithoutPadding, err := PKCS5Trimming(decrypted)
	if err != nil {
		return "", err
	}
	return string(decryptedWithoutPadding), nil
}

// PKCS5Padding_ 加密中使用，填充为整个的一块
func PKCS5Padding_(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming 解密中使用，解填充
func PKCS5Trimming(encrypt []byte) ([]byte, error) {
	padding := encrypt[len(encrypt)-1]
	highIndex := len(encrypt) - int(padding)
	if highIndex < 0 || highIndex > len(encrypt) {
		return nil, ErrInvalidToken
	}
	return encrypt[:highIndex], nil
}

func main() {
	a, _ := AesCbcEncrypt("com.lemon.lvoverseas")
	fmt.Println(a)
	// b, _ := AesCbcDecrypt(a)
	// fmt.Println(b)
	// //b, _ := AesCbcDecrypt("rz7gtVHaZRFt64hNWvnRxcsqZYWGoHu3MuQ2N+lUG/tqyHDHUiieq0Xz/gNXPe30")
	// fmt.Println(b)
	// fmt.Println(len(b))
}
