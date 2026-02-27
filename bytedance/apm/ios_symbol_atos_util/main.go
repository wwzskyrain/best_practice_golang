package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	exeGoVersion()
	// exeGoatos()
}

func exeGoatos() {
	// 设置 atos 命令的参数
	uikitPath := "/Users/bytedance/Downloads/10.2.1 (14D27)/Symbols/System/Library/Frameworks/UIKit.framework/UIKit"
	loadAddress := "0x192c05000"
	targetAddress := "0x192e0a79c"

	// 构造 atos 命令
	cmd := exec.Command("atos", "-o", uikitPath, "-l", loadAddress, targetAddress)

	// 获取命令输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing atos:", err)
		return
	}

	// 打印解析结果
	fmt.Println("Symbolicated Address:", strings.TrimSpace(string(output)))
}

func exeGoVersion() {
	// 构造 atos 命令
	cmd := exec.Command("go", "version")

	// 获取命令输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing go:", err)
		return
	}

	// 打印解析结果
	fmt.Println("go", strings.TrimSpace(string(output)))
}
