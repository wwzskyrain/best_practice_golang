package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Common struct {
	Name  string `validate:"not empty"`
	Class int    `validate:">1"`
}

func Validate(obj interface{}) error {
	st := reflect.TypeOf(obj)
	// 检测传入是否为结构体
	if st.Kind() != reflect.Struct {
		return errors.New("not struct")
	}
	v := reflect.ValueOf(obj)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		// 获取tag为alias的值
		if tagVal, ok := field.Tag.Lookup("validate"); ok {
			if tagVal == "" {
				continue
			}
			var val interface{}
			switch tagVal {
			case ">1":
				if v.Field(i).Kind() == reflect.Uint {
					if v.Field(i).Interface().(int) <= 0 {
						return failed(field.Name, val, tagVal)
					}
				}
			case "not empty":
				if v.Field(i).Kind() == reflect.String { // wwz-类型断言和转化可以一句话搞定吗？
					val := v.Field(i).Interface().(string)
					if val == "" {
						return failed(field.Name, val, tagVal)
					}
				}
			}
		}
	}
	return nil
}

func failed(str string, val interface{}, expect string) error {
	err := errors.New(fmt.Sprintf("invalid %s,expect %v,acutal %s,", str, val, expect))
	fmt.Println(err.Error())
	return err
}

func main() {
	a := []Common{
		{
			Name:  "1",
			Class: 0,
		},
		{
			Name:  "",
			Class: 1,
		},
		{
			Name:  "1",
			Class: 1,
		},
	}
	for i, v := range a {
		err := Validate(v)
		if err != nil {
			fmt.Printf("index:%d,validate failed:%s\n", i, err.Error())
			continue
		}
		fmt.Printf("index:%d validate success\n", i)
	}
	return
}
