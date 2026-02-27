package main

import (
	"context"
	"fmt"

	"github.com/pratice/golang/util/time_util"
)

// 验证ctx这棵树。构造了一个两层的完全三叉树，第一阶段取消了第二颗子树，第二阶段是取消了root节点。
func main() {

	rootCtxContext, rootCtxCancel := context.WithCancel(context.Background())
	TestCancelSubTree(rootCtxContext)

	rootCtxCancel()
	time_util.CountDownToZero("main", 3)
}

func TestCancelSubTree(rootCtx context.Context) {
	// 第一层的三个
	ctx_1, _ := context.WithCancel(rootCtx)
	ctx_2, ctx_2_cancel := context.WithCancel(rootCtx)
	ctx_3, _ := context.WithCancel(rootCtx)
	go DoTask(ctx_1, "ctx_1")
	go DoTask(ctx_2, "ctx_2")
	go DoTask(ctx_3, "ctx_3")

	// 第二层的第一个三个
	ctx_1_1, _ := context.WithCancel(ctx_1)
	ctx_1_2, _ := context.WithCancel(ctx_1)
	ctx_1_3, _ := context.WithCancel(ctx_1)

	go DoTask(ctx_1_1, "ctx_1_1")
	go DoTask(ctx_1_2, "ctx_1_2")
	go DoTask(ctx_1_3, "ctx_1_3")

	// 第二层的第二个三个
	ctx_2_1, _ := context.WithCancel(ctx_2)
	ctx_2_2, _ := context.WithCancel(ctx_2)
	ctx_2_3, _ := context.WithCancel(ctx_2)

	go DoTask(ctx_2_1, "ctx_2_1")
	go DoTask(ctx_2_2, "ctx_2_2")
	go DoTask(ctx_2_3, "ctx_2_3")

	// 第二层的第三个三个
	ctx_3_1, _ := context.WithCancel(ctx_3)
	ctx_3_2, _ := context.WithCancel(ctx_3)
	ctx_3_3, _ := context.WithCancel(ctx_3)

	go DoTask(ctx_3_1, "ctx_3_1")
	go DoTask(ctx_3_2, "ctx_3_2")
	go DoTask(ctx_3_3, "ctx_3_3")

	println("开启了【3 + 9】个task")
	time_util.CountDownToZero("见证奇迹", 3)
	ctx_2_cancel()
	time_util.CountDownToZero("等3秒", 3)
}

func DoTask(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Printf("context %s done with %+v\n", name, ctx.Err())
	}
}
