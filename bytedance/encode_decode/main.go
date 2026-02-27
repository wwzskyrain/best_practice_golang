package main

import (
	bddecrypt "code.byted.org/dp/bytedance_encdec/decrypt"
	bdencrypt "code.byted.org/dp/bytedance_encdec/encrypt"
)

// 公司的加解密代码库： https://bits.bytedance.net/code/dp/bytedance_encdec
func main() {

	rawData := "我是一条raw数据"
	rawDataByteArr := []byte(rawData)
	encryptData, err := bdencrypt.Encrypt(rawDataByteArr)
	if err != nil {
		println("加密失败！！")
	}
	decryDataArr, err := bddecrypt.Decrypt(encryptData)
	if err != nil {
		println("解密失败！！")
	}
	decryData := string(decryDataArr)
	println(decryData)
}
