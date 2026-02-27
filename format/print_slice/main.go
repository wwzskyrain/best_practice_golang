package main

import (
	"fmt"
)

type Person struct {
	age  int    `json:"age"`
	name string `json:"name"`
}

func (p *Person) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdMaterial(%+v)", *p)
}

func main() {
	s := []int64{1, 2, 3}
	fmt.Printf("s = %v\n", s)

	persons := []*Person{
		{
			1, "name1",
		},
		{
			2, "name2",
		},
		{
			3, "name3",
		},
	}

	fmt.Printf("persons = %v\n", persons)
	for _, person := range persons {
		fmt.Printf("person= %v\n", person)
	}

	m := make(map[int]*Person)
	for _, person := range persons {
		m[person.age] = person
	}

	for k, person := range m {
		fmt.Printf("%d = %v \n", k, person)
	}

	fmt.Printf("m = %v \n", m)

}
