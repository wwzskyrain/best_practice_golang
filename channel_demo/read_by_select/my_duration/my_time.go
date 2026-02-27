package main

type MyDuration int64

const (
	Nanosecond  MyDuration = 1
	Microsecond            = 1000 * Nanosecond
	Millisecond            = 1000 * Microsecond
	Second                 = 1000 * Millisecond
	Minute                 = 60 * Second
	Hour                   = 60 * Minute
)

func main() {
	// 自定义的Duration也能很好的支持加减乘除(虽然这里只测试了乘法)
	d := MyDuration(10) * Second
	d1 := 1000 * Millisecond
	println("%d, %d", d, d1) //占位符输出失败，哈哈
	println("%d", d-d1)
	println("%d", d-d1-d1)

}
