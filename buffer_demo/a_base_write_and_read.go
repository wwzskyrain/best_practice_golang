package buffer_demo

import (
	"bytes"
	"fmt"
)

func writeAndReadEnglish() {
	buf := bytes.NewBuffer(nil) // 初始化空缓冲区

	// 写入数据
	buf.Write([]byte("Hello")) // 写入字节切片
	buf.WriteString(" World!") // 直接写入字符串
	buf.WriteByte('!')         // 写入单个字节

	// 读取数据
	p := make([]byte, 2)
	n, _ := buf.Read(p) // 读取到字节切片（n是实际读取长度）
	fmt.Printf("实际读取长度 n=%d\n", n)
	fmt.Printf("读取内容=%s\n", string(p[:n]))
	str, _ := buf.ReadString(' ') // 读取直到分隔符（如空格）
	fmt.Printf("读取内容=%s\n", str)
	fmt.Printf("剩下内容=%s\n", string(buf.Bytes()))
}

func writeAndReadChinese() {
	buf := bytes.NewBuffer(nil) // 初始化空缓冲区

	// 写入数据
	buf.Write([]byte("欢迎你"))  // 写入字节切片
	buf.WriteString(" 我的朋友！") // 直接写入字符串
	buf.WriteByte('!')        // 写入单个字节

	// 读取数据
	p := make([]byte, 12)
	n, _ := buf.Read(p) // 读取到字节切片（n是实际读取长度）
	fmt.Printf("实际读取长度 n=%d\n", n)
	fmt.Printf("读取内容=%s\n", string(p[:n]))
	str, _ := buf.ReadString(' ') // 读取直到分隔符（如空格）
	fmt.Printf("读取内容=%s\n", str)
	fmt.Printf("剩下内容=%s\n", string(buf.Bytes()))
}
