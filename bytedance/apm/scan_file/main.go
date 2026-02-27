package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 递归扫描目录，建立库名到库文件路径的映射
func scanSystemLibraries(rootDir string) error {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// 仅处理 .dylib 和 framework 二进制文件
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".dylib") || strings.HasSuffix(path, ".framework/"+info.Name())) {
			libName := info.Name()
			fmt.Printf("Indexed: %s -> %s\n", libName, path)
		}
		return nil
	})
	return nil
}

func main() {
	scanSystemLibraries("/mnt/symbols/14.2_18B92(Symbols)/System/Library")
}
