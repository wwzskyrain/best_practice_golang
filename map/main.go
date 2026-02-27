package main

import "fmt"

func testAppendValueToSliceWhenSliceIsMapValue() {

	m := map[string]int{
		"wang":  18,
		"zhang": 18,
		"li":    16,
	}

	fmt.Printf("%v \n", m)

	age2NameSliceMap := make(map[int][]string)

	for name, age := range m {
		names, ok := age2NameSliceMap[age]
		if !ok {
			names = make([]string, 0)
		}
		names = append(names, name)
		age2NameSliceMap[age] = names
	}

	fmt.Printf("age2NameSliceMap = %v \n", age2NameSliceMap)

}

func main() {
	// testAppendValueToSliceWhenSliceIsMapValue()
	// Test_Npe()
	TestMapGetAndCast()
}

func Test_Npe() {
	names := []string{"1", "2"}
	var siteIdMap = make(map[string]bool)
	for _, name := range names {
		siteIdMap[name] = true
	}
}

func TestMapGetAndCast() {
	// 当直接用 m[key].(type)时，是分不清楚是因为没有key还是类型不对的.
	m := map[string]interface{}{
		"wang":  18,
		"zhang": 18,
		"li":    16,
	}
	w := m["wang"].(int)
	fmt.Printf("w = %d\n", w)

	if zhao, ok := m["zhao"].(int); ok {
		fmt.Printf("zhao = %d\n", zhao)
	} else {
		fmt.Printf("没有【zhao】\n")
	}

	if zhao, ok := m["zhao"].(string); ok {
		fmt.Printf("zhao = %s\n", zhao)
	} else {
		fmt.Printf("【zhao】 cast失败\n")
	}

}
