package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "2.0.54"
	output := strings.ReplaceAll(input, ".", "_")
	fmt.Println(output) // 输出: a_b_c_d
}
