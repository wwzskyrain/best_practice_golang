package str_cat_test

import (
	"strings"
	"testing"
)

// chapter8/sources/benchmark-compare/strcat_test.go

var sl = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func Strcat(sl []string) string {
	return concatStringByOperator(sl)
	return concatStringByJoin(sl)
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func BenchmarkStrcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Strcat(sl)
	}
}
