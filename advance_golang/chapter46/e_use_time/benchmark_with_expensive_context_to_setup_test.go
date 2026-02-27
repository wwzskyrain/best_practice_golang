package e_use_time

import (
	"strings"
	"testing"
	"time"
)

// chapter8/sources/benchmark_with_expensive_context_setup_test.go

var sl = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func expensiveTestContextSetup() {
	time.Sleep(200 * time.Millisecond)
}

func BenchmarkStrcatWithTestContextSetup(b *testing.B) {
	expensiveTestContextSetup()
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkStrcatWithTestContextSetupAndResetTimer(b *testing.B) {
	expensiveTestContextSetup()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkStrcatWithTestContextSetupAndRestartTimer(b *testing.B) {
	b.StopTimer()
	expensiveTestContextSetup()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkStrcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

// $GOROOT/src/runtime/map_test.go
func benchmarkMapDeleteInt32(b *testing.B, n int) {
	a := make(map[int32]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if len(a) == 0 { // 这个初始化，是没有计算在内的，注意这个初始化只会走一次。很奇怪，为什么不把if挪到for外面呢？
			b.StopTimer()
			for j := i; j < i+n; j++ {
				a[int32(j)] = j
			}
			b.StartTimer()
		}
		delete(a, int32(i))
	}
}

// $go test -bench . benchmark_with_expensive_context_setup_test.go
// goos: darwin
// goarch: amd64
// BenchmarkStrcatWithTestContextSetup-8                 16943037     65.9 ns/op
// BenchmarkStrcatWithTestContextSetupAndResetTimer-8    21700249     52.7 ns/op
// BenchmarkStrcatWithTestContextSetupAndRestartTimer-8  21628669     50.5 ns/op
// BenchmarkStrcat-8                                     22915291     50.7 ns/op
// PASS
// ok       command-line-arguments 9.838s
