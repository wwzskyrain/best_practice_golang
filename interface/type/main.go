package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	age int
}

func main() {
	var p interface{}
	p = Person{age: 12}
	fmt.Printf("%+v ,%T \n", p, p)
	p = "hello world"
	// 打印变量的类型
	fmt.Printf("%+v ,type= %T, type= %s, value=%v \n", p, p, reflect.TypeOf(p), p)

	switch p.(type) { // 断言变量类型
	case interface{}: // 把interface{}放在前面，则走该分支，把string放到前面，则走string分支
		fmt.Printf("p is interface{} ! \n")
	case string:
		fmt.Printf("p is tring ! \n")
	default:
		fmt.Printf("unkown type of p, print it %T \n", p)
	}

	// 	断言变量类型
	if str, ok := p.(string); ok {
		fmt.Printf("断言成功，string类型啦，v=%s \n", str)
		fmt.Printf("len(str)= %d \n", len(str))
	}

	// 	断言变量类型
	if per, ok := p.(Person); !ok {
		// 断言失败，其变量per也是Person，不过是一个零值对象
		fmt.Printf("断言失败嘞，不是Person类型；那么这时候per变量是什么呢？ v=%+v , reflect.TypeOf() = %s", per, reflect.TypeOf(per))
	}

}
