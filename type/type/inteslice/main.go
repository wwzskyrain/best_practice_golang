package main

import (
	"fmt"
	"sort"
)

func main() {
	lst := []int{4, 5, 2, 8, 1, 9, 3}
	sort.Sort(sort.IntSlice(lst)) // sort.IntSlice 给基本类型[]int，添加了一些方法，让它有某些资格（方法集合就是资格），这里就是接口 sort.Interface
	sort.IntSlice(lst).Sort()     // 等同于上句话 sort.Sort(sort.IntSlice(lst))，因为sort.IntSlice给[]int挂靠接口sort.Interface时，顺带实现了Sort()方法
	fmt.Println(lst)
}
