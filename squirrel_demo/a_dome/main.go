package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	query := sq.Select("id", "name").
		From("users").
		Where(sq.Eq{"status": "active", "role": "admin"}).
		Limit(10)

	sql, args, err := query.ToSql()
	if err != nil {
		fmt.Println("构建 SQL 出错:", err)
		return
	}

	fmt.Println("生成的 SQL:", sql)
	fmt.Println("参数:", args)
}
