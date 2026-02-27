package a_benchmark_first

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// chapter8/sources/benchmark-impl/sequential_test.go
var (
	m            = make(map[int64]struct{}, 10)
	mu           sync.Mutex
	round        int64 = 1
	goroutineNum       = uint32(1)
)

func BenchmarkSequential(b *testing.B) {

	fmt.Printf("\nenter BenchmarkSequential: round[%d], b.N[%d]\n", atomic.LoadInt64(&round), b.N)

	defer func() {
		atomic.AddInt64(&round, 1)
	}()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		_, ok := m[round]
		if !ok {
			m[round] = struct{}{}
			fmt.Printf("enter loop in BenchmarkSequential: round[%d], b.N[%d]\n", atomic.LoadInt64(&round), b.N)
		}
		mu.Unlock()
	}
	fmt.Printf("exit BenchmarkSequential: round[%d], b.N[%d]\n", atomic.LoadInt64(&round), b.N)
}
