package buffer_demo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/pratice/golang/gzip_demo"
)

func asIoWriter() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 【主角】初始化 Buffer（实现了 io.Writer）
	var buf bytes.Buffer

	// 将 JSON 数据写入 Buffer
	user := User{Name: "Alice", Age: 25}
	enc := json.NewEncoder(&buf) // json.Encoder 需要 io.Writer
	if err := enc.Encode(user); err != nil {
		panic(err)
	}

	// 输出缓冲区内容
	fmt.Printf("JSON Data: %s", buf.String())

	// 当然这里json的功能就是一个序列化Marshal，如下
	marshal, _ := json.Marshal(user)
	fmt.Printf("marshal=%s\n", string(marshal))
}

func asIoReader() {
	// 模拟一个 Gzip 压缩的数据
	compressedData, _ := gzip_demo.CompressJSONToGzip([]byte("你好，我的宝宝！"))

	// 【主角】将压缩数据放入 Buffer（实现了 io.Reader）
	buf := bytes.NewBuffer(compressedData)

	// 从 Buffer 读取并解压数据
	gzReader, err := gzip.NewReader(buf) // gzip_demo.NewReader 需要 io.Reader
	if err != nil {
		panic(err)
	}
	defer gzReader.Close()

	// 读取解压后的数据
	decompressed, err := io.ReadAll(gzReader)
	if err != nil {
		fmt.Printf("error=%s", err.Error())
		panic(err)
	}

	fmt.Printf("Decompressed: %s\n", decompressed) // 输出: "hello"
}

// 一个作为reader一个作为writer，使用io.Copy方法来读和写.
func asReaderAndWriter() {
	source := bytes.NewBuffer([]byte("hello world"))
	var destination bytes.Buffer
	_, err := io.Copy(&destination, source)
	if err == nil {
		fmt.Println(destination.String())
	}

	fmt.Printf("---自己实现copy哈---\n")
	// 当然你去看Copy的实现，有意思呢
	r := bytes.NewBuffer([]byte("hello world"))
	var w bytes.Buffer
	n, err := r.WriteTo(&w)
	if err != nil {
		panic(err)
	}
	fmt.Printf("复制了多少字节：%d\n", n)
	fmt.Printf("w的内容：%s\n", w.String())
}

// 读文件，读到buffer中
func asReadFrom() {
	var buf bytes.Buffer

	// 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将文件内容全部读取到 Buffer（利用 ReadFrom 接口）
	n, err := buf.ReadFrom(file) // 内部会自动扩容
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf.String())
	fmt.Printf("第二次读取（即调用String()： %d bytes: %s\n", n, buf.String())
}

// 写到文件中.
func asFileWriter() {
	// 初始化一个带数据的 Buffer
	buf := bytes.NewBufferString("哈喽，云妹，你就成全师哥吧！")

	// 创建文件
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将 Buffer 数据写入文件（利用 WriteTo 接口）
	n, err := buf.WriteTo(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", n)

	// 如果再写一遍呢？将 Buffer 数据写入文件（利用 WriteTo 接口）
	n, err = buf.WriteTo(file)
	if err != nil {
		panic(err)
	}
	// 第二次WriteTo就没东西写了，就只有0个字节了.
	fmt.Printf("Wrote %d bytes\n", n)
}
