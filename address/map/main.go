package main

import "fmt"

type Person struct {
	name string
	age  int
}

func testMapValueIsAddressable() {
	m := map[string]Person{
		"wang":  {name: "wang", age: 18},
		"Li":    {name: "Li", age: 18},
		"Zhang": {name: "Zhang", age: 18},
	}

	z := m["Zhang"]
	z.age = 28
	fmt.Printf("%v, %v", m, z) // 果然，map中的value只是一个值，是不可改变的
}

func testMapValuePointer() {

	m := map[string]*Person{}
	m["wang"] = &Person{name: "wang", age: 18}
	m["Li"] = &Person{name: "Li", age: 18}
	m["Zhang"] = &Person{name: "Zhang", age: 18}

	z := m["Zhang"]
	z.age = 28
	fmt.Printf("%v, %v", m, z) // 果然，map中的value只是一个值，是不可改变的

}

func main() {
	testMapValueIsAddressable()
	testMapValuePointer()
}
