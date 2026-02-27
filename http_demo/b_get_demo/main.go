package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler() {
	// 发起 HTTP 请求
	resp, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取响应体数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}
	for k, v := range resp.Header {
		fmt.Printf("header= [%s:%s]\n", k, v)
	}

	// 打印读取的内容
	fmt.Println("Fetched data from HTTP response:")
	fmt.Println(string(body))

}

func main() {
	handler()
}
