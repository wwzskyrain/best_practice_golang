package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("升序排列：")
	sortByAsc()

	fmt.Printf("升序排列：")
	sortByDesc()
}

func sortByAsc() {
	lst := []int{4, 5, 2, 8, 1, 9, 3}
	sort.Sort(sort.IntSlice(lst))
	fmt.Println(lst)
}

func sortByDesc() {
	lst := []int{4, 5, 2, 8, 1, 9, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(lst)))
	fmt.Println(lst)
}
