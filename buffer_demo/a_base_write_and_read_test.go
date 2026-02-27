package buffer_demo

import (
	"fmt"
	"testing"
)

func TestBaseWriteAndRead(t *testing.T) {
	writeAndReadEnglish()
	fmt.Println()
	// 会有有缺的事情发生.
	writeAndReadChinese()
}
