package main

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Book 定义图书结构体
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// 模拟数据库
var books = []Book{
	{ID: 1, Title: "《论语》", Author: "孔子弟子及再传弟子"},
	{ID: 2, Title: "《孟子》", Author: "孟子及其弟子万章、公孙丑等"},
	{ID: 3, Title: "《诗经》", Author: "佚名，传为尹吉甫采集、孔子编订"},
	{ID: 4, Title: "《史记》", Author: "司马迁"},
	{ID: 5, Title: "《资治通鉴》", Author: "司马光"},
	{ID: 6, Title: "《红楼梦》", Author: "曹雪芹"},
	{ID: 7, Title: "《西游记》", Author: "吴承恩"},
	{ID: 8, Title: "《水浒传》", Author: "施耐庵"},
	{ID: 9, Title: "《三国演义》", Author: "罗贯中"},
	{ID: 10, Title: "《西厢记》", Author: "西厢记作者"},
}

func main() {
	h := server.New(server.WithHostPorts(":6789"))

	// 注册路由
	h.GET("/books", getBooks)
	h.POST("/books", saveBook)
	h.POST("/books/search", searchBooks)

	// 启动服务器
	h.Spin()
}

// getBooks 查询图书
func getBooks(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{
		"books": books,
	})
}

// saveBook 保存图书
func saveBook(c context.Context, ctx *app.RequestContext) {
	var newBook Book
	if err := ctx.BindAndValidate(&newBook); err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"error": err.Error(),
		})
		return
	}

	// 模拟生成图书 ID
	newBook.ID = len(books) + 1
	books = append(books, newBook)

	ctx.JSON(consts.StatusCreated, utils.H{
		"message": "Book saved successfully",
		"book":    newBook,
	})
}

// searchBooks 检索图书
func searchBooks(c context.Context, ctx *app.RequestContext) {
	var query Book
	if err := ctx.BindAndValidate(&query); err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"error": err.Error(),
		})
		return
	}

	var result []Book
	for _, book := range books {
		if (query.Title == "" || strings.Contains(book.Title, query.Title)) &&
			(query.Author == "" || strings.Contains(book.Author, query.Author)) {
			result = append(result, book)
		}
	}

	ctx.JSON(consts.StatusOK, utils.H{
		"books": result,
	})
}
