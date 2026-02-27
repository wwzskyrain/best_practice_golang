package main

import "fmt"

func main() {
	// testPrintSlice()
	// testAppend()

	for i := 0; i < 60; i++ {
		fmt.Printf("2^%d = %d\n", i, 1<<i)
	}

}

func testPrintSlice() {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = i * 10
	}
	fmt.Println(s)

	s = s[:5]
	fmt.Println(s)
}

func testAppend() {
	var s []int
	s1 := geneSlice(1)
	s = append(s, s1...)
	s2 := geneSlice(2)
	s = append(s, s2...)
	s3 := geneSlice(3)
	s = append(s, s3...)
	fmt.Printf("s = %+v", s)
}

func geneSlice(command int) []int {
	switch command {
	case 1:
		return nil
	case 2:
		return []int{1, 2, 3}
	default:
		return []int{1, 2, 3, 4, 5, 6}

	}
}
