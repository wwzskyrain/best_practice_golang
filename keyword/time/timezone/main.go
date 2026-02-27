package main

import (
	"fmt"
	"time"
)

const Layout = "2006/01/02 03:04:05PM (UTC-07)"

func formatDemo() {
	n := time.Now()
	// 两个格式化
	fmt.Println(n.Format(Layout))

	// fmt.Println(n.Format(time.DateOnly))
}

func zoneDemo() {
	n := time.Now()
	fmt.Println(n.Location())
	fmt.Println(n.Zone())
}

func main() {
	formatDemo()
	// zoneDemo()
}
