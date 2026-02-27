package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	query := sq.Select("id", "name").
		From("users").
		Where(sq.Eq{"status": "active", "role": "admin"}).
		Limit(10)

	sql, args, err := query.ToSql()
	if err != nil {
		log.Fatal("构建 SQL 出错:", err)
	}

	fmt.Println("SQL 模板:", sql) // SELECT id, name FROM users WHERE status = ? AND role = ? LIMIT 10
	fmt.Println("参数:", args)    // [active admin]

	// 构造 HTTP 请求参数
	params := url.Values{}
	params.Set("query", sql)

	// 给每个参数取一个名字，并传递到 URL 中
	for i, arg := range args {
		params.Set(fmt.Sprintf("param%d", i+1), fmt.Sprintf("%v", arg))
	}

	// 构造 ClickHouse HTTP 请求
	endpoint := "http://your-clickhouse-host:8123"
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(params.Encode()))
	if err != nil {
		log.Fatal("构造请求失败:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发起 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("响应状态码:", resp.StatusCode)
}
