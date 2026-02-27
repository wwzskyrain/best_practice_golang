package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 确保 goatos 可执行文件存在
	goatosDir := "/Users/bytedance/work_dir/goatos-macos/goatos_darwin_arm64"
	goatosPath := "/Users/bytedance/work_dir/goatos-macos/goatos_darwin_arm64/goatos"
	if _, err := os.Stat(goatosPath); os.IsNotExist(err) {
		fmt.Sprintf("goatos tool not found at: %s", goatosPath)
	}
	libPath := "/Users/bytedance/Downloads/14.2_18B92(arm64)/Symbols/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation"
	// 构造命令
	cmd := exec.Command(goatosPath, "-o", libPath, "-l", "0x19a4c3000", "0x000000019a5dd9d4")

	// 设置工作目录（切换到 goatos 目录）
	cmd.Dir = goatosDir

	// 捕获命令的标准输出和错误输出
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Sprintf("failed to execute goatos: %v, stderr: %s", err, errBuffer.String())
	}

	// 解析输出结果
	output := strings.TrimSpace(outBuffer.String())
	if output == "" {
		fmt.Printf("output = empty")
	}
	fmt.Printf("output = %s\n", output)
}
