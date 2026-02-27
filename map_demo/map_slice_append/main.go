package main

import "fmt"

func main() {

	m := make(map[int][]string)
	m[0] = append(m[0], "0")
	m[0] = append(m[0], "1")
	m[0] = append(m[0], "2")

	for k, v := range m {
		fmt.Printf("%d = %s\n", k, v)
	}
}
