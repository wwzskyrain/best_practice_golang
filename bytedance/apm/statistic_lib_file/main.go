package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "/Users/bytedance/Downloads/14.2_18B92(arm64)/Symbols"
	doScan(root)
}

func doScan(root string) {
	fileMap := make(map[string][]string) // 文件名 -> 路径列表
	var fileCount, dirCount int64
	var totalSize int64
	var duplicatedNameCount int32

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil
		}

		if info.IsDir() {
			dirCount++
		} else {
			fileCount++
			totalSize += info.Size()
			fileMap[info.Name()] = append(fileMap[info.Name()], path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
		return
	}

	// 输出文件名 -> 路径映射
	fmt.Println("File Mapping:")
	for name, paths := range fileMap {
		fmt.Printf("file_name=[%s] paths=%s\n", name, strings.Join(paths, ","))
		if len(paths) > 2 {
			duplicatedNameCount++
		}
	}

	// 统计信息
	fmt.Println("\nStatistics:")
	fmt.Printf("Total files: %d\n", fileCount)
	fmt.Printf("Total directories: %d\n", dirCount)
	fmt.Printf("Total size: %.2f MB\n", float64(totalSize)/(1024*1024))
	fmt.Printf("Duplicated file names: %d\n", duplicatedNameCount)
}
