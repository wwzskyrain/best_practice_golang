package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

// 压缩文本
func compressText(text string) ([]byte, error) {
	var buf bytes.Buffer
	// 创建一个新的 gzip.Writer，它会将压缩后的数据写入 buf
	gz := gzip.NewWriter(&buf)
	// 将文本写入 gzip.Writer
	_, err := gz.Write([]byte(text))
	if err != nil {
		return nil, err
	}
	// 关闭 gz 以确保数据写入完毕
	err = gz.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解压文本
func decompressText(compressed []byte) (string, error) {
	// 创建一个读取 gzip 数据的 Reader
	gz, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return "", err
	}
	// 使用缓冲区来存储解压后的内容
	var result bytes.Buffer
	// 从 gz 读取解压数据并写入 result
	_, err = io.Copy(&result, gz)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func main() {
	// 原始文本
	text := "hello world golang"

	// 压缩
	compressed, err := compressText(text)
	if err != nil {
		log.Fatalf("压缩失败: %v", err)
	}
	fmt.Println("压缩后的数据：", compressed)

	// 解压
	decompressed, err := decompressText(compressed)
	if err != nil {
		log.Fatalf("解压失败: %v", err)
	}
	fmt.Println("解压后的数据：", decompressed)
}
