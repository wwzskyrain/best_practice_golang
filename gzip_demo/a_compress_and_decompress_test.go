package gzip_demo

import (
	"fmt"
	"testing"
)

var originJsonData = `
{
    "apm_id": "20000001",
    "header":
    {
        "custom":
        {
            "sdk_version": "6.4.0.1",
            "host_app_id": "8025677"
        },
        "os": "Android",
        "os_version": "12",
        "device_model": "SM-A217F",
        "device_brand": "samsung",
        "sdk_version_name": "0.0.5",
        "aid": "10000000",
        "update_version_code": 6401,
        "bd_did": "63186367869893"
    },
    "local_time": 1735628334587,
    "launch":
    [
        {
            "local_time_ms": 1735628334587
        }
    ]
}
`

func TestCompressAndDecompress(t *testing.T) {
	// 压缩
	originJsonDataByte := []byte(originJsonData)
	compressedData, err := CompressJSONToGzip(originJsonDataByte)
	if err != nil {
		panic(err)
	}

	// 压缩前后的len对比
	fmt.Printf("Original: %d bytes\n", len(originJsonData))
	fmt.Printf("Compressed: %d bytes\n", len(compressedData))

	// 解压
	decompressedBytes, err := DecompressGzipToJSON(compressedData)
	if err != nil {
		panic(err)
	}
	// 解压数据的长度和内容.
	fmt.Printf("decompressed Data: %d bytes\n", len(decompressedBytes))
	fmt.Printf("Original Date : %s", string(decompressedBytes))
}
