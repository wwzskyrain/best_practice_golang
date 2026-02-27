package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"code.byted.org/gopkg/logs"
)

// 解压内存中的 ZIP 文件到本地目录
func unzipFromMemory(zipData []byte, destDir string) error {
	// 创建一个 bytes.Reader 读取内存中的 zip 数据
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %v", err)
	}

	// 遍历 ZIP 文件中的所有文件
	for _, file := range zipReader.File {
		// 过滤 __MACOSX 目录和隐藏文件（._ 开头）
		if strings.HasPrefix(file.Name, "__MACOSX/") || strings.Contains(file.Name, "/._") {
			fmt.Println("Skipping:", file.Name) // 日志输出跳过的文件
			continue
		}
		logs.Info("file_name=%s", file.Name)
		destFilePath := filepath.Join(destDir, file.Name)
		// 检查并创建目录（如果是文件夹）
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(destFilePath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			continue
		}

		// 打开 ZIP 内的文件
		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip file: %v", err)
		}
		defer srcFile.Close()

		// 创建目标文件
		if err := os.MkdirAll(filepath.Dir(destFilePath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create parent directories: %v", err)
		}
		dstFile, err := os.Create(destFilePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer dstFile.Close()

		fmt.Println("Unzipping", file.Name, "to", destFilePath)
		// 复制文件内容到本地
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file content: %v", err)
		}
	}

	return nil
}

// 示例：解压一个从对象存储下载的 ZIP 文件
func main() {
	// 假设 zipData 是从对象存储获取的 ZIP 文件数据
	zipData, err := os.ReadFile("zip_file/test_zip_pdf.zip") // 这里模拟从磁盘读取
	// zipData, err := os.ReadFile("zip_file/13.3.1 (17D50).zip") // 这里模拟从磁盘读取
	if err != nil {
		fmt.Println("Failed to read zip file:", err)
		return
	}

	destDir := "./unzipped_files"
	if err := unzipFromMemory(zipData, destDir); err != nil {
		fmt.Println("Failed to unzip file:", err)
		return
	}

	fmt.Println("Successfully extracted files to:", destDir)
}
