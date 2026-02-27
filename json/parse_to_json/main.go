package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	// demoFromGPT4()

	// debugJsonGet()
	parseJsonStrToMap1()
	// parseJsonStrToMap2()
	// parseJsonStrToMap3()

}

func parseJsonStrToMap1() {
	jsonStr := `{
		"10028": -1,
		"10220": -6,
		"10064": -6
	}`

	aggregateConfMap := make(map[int64]int64)
	err := json.Unmarshal([]byte(jsonStr), &aggregateConfMap)
	if err != nil {
		fmt.Printf("err=%v", err)
	} else {
		fmt.Printf("aggregateConfMap = %v \n", aggregateConfMap)
	}
}

func parseJsonStrToMap2() {
	jsonStr := `{
    "-1":
    [
        10028,
        10100,
        10269
    ],
    "-2":
    [
        10065,
        10175
    ],
    "-6":
    [
        10063,
        10220,
        10064
    ],
    "-7":
    [
        10192,
        10205
    ]
}`

	spreadConfMap := make(map[int64][]int64)
	err := json.Unmarshal([]byte(jsonStr), &spreadConfMap)
	if err != nil {
		fmt.Printf("err=%v", err)
	} else {
		fmt.Printf("spreadConfMap = %v", spreadConfMap)
	}
}

func parseJsonStrToMap3() {
	jsonStr := `{
    "AggregationConf":
    {
        "10028": -1,
        "10100": -1,
        "10269": -1,
        "10065": -2,
        "10175": -2,
        "10192": -7,
        "10205": -7,
        "10020": -7,
        "10193": -7,
        "10194": -7,
        "10229": -7,
        "10226": -7,
        "10227": -7,
        "10021": -7,
        "10212": -7,
        "10213": -7,
        "10063": -6,
        "10220": -6,
        "10064": -6
    },
    "SpreadConf":
    {
        "-1":
        [
            10028,
            10100,
            10269
        ],
        "-2":
        [
            10065,
            10175
        ],
        "-6":
        [
            10063,
            10220,
            10064
        ],
        "-7":
        [
            10192,
            10205,
            10020,
            10193,
            10194,
            10229,
            10226,
            10227,
            10021,
            10212,
            10213
        ]
    }
}`
	conf := MatTagAggregation{}
	json.Unmarshal([]byte(jsonStr), &conf)
	fmt.Printf("conf=%v", conf)
}

type MatTagAggregation struct {
	AggregationConf map[int]int
	SpreadConf      map[int][]int
}

func debugJsonGet() {

	js := generateJsonObject()
	// 每个v底层都有一个type属性的
	if v, err := js.Get("name").String(); err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(v)
	}
}

func generateJsonObject() simplejson.Json {
	// 我们要先有一个 JSON 字符串
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "isMarried": true,
        "interests": ["Reading", "Cycling", "Hiking"]
    }`
	// 然后使用 simplejson.NewJson 方法解析 JSON 字符串
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err)
	}
	return *js
}

// 这是一个从gpt4直接copy出来的代码，真的很好用
func demoFromGPT4() {

	// 假设我们有一个 JSON 字符串
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "isMarried": true,
        "interests": ["Reading", "Cycling", "Hiking"]
    }`

	// 使用 simplejson.NewJson 方法解析 JSON 字符串
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err)
	}

	// 获取并打印 name 字段
	name, err := js.Get("name").String()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)

	// 获取并打印 age 字段
	age, err := js.Get("age").Int()
	if err != nil {
		panic(err)
	}
	fmt.Println("Age:", age)

	// 获取并迭代打印 interests 数组
	interests, err := js.Get("interests").Array()
	if err != nil {
		panic(err)
	}
	fmt.Println("Interests:")
	for i, val := range interests {
		fmt.Printf(" - Interest #%d: %s\n", i+1, val)
	}

	// 修改 age 字段
	js.Set("age", 31)

	// 添加一个新字段
	js.Set("city", "New York")

	// 将修改后的对象转换回 JSON 字符串
	newJsonStr, err := js.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Modified JSON:", string(newJsonStr))
}
