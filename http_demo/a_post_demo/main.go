package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Book 定义图书结构体，与服务端保持一致
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// 定义搜索条件
	query := Book{
		Title:  "三国",
		Author: "",
	}

	// 将搜索条件转换为 JSON 字节切片
	queryJSON, err := json.Marshal(query)
	if err != nil {
		fmt.Println("Error marshaling query:", err)
		return
	}

	// 创建 HTTP POST 请求
	resp, err := http.Post("http://localhost:6789/books/search", "application/json", bytes.NewBuffer(queryJSON))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 打印响应状态码和响应体
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
