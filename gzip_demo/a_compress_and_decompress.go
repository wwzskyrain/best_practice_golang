package gzip_demo

import (
	"bytes"
	"compress/gzip"
	"io"
)

// CompressJSONToGzip 压缩
func CompressJSONToGzip(jsonData []byte) ([]byte, error) {
	var buf bytes.Buffer

	// 使用最佳压缩级别（可选）
	gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := gz.Write(jsonData); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecompressGzipToJSON 解压
func DecompressGzipToJSON(gzipData []byte) ([]byte, error) {
	// 1. 创建一个 bytes.Reader 读取压缩数据
	reader := bytes.NewReader(gzipData)
	// 2. 创建 Gzip 读取器
	gz, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	// 3. 读取解压后的数据
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, gz); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
