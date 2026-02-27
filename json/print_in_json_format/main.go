package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pratice/golang/util/pointer_yy"
)

func main() {
	TestPrintJsonMap()
}

type Creative struct {
	Id        *int64  `json:"id"`
	ImageInfo *string `json:"image_info"`
}

func TestPrintJsonMap() {
	structJsonSlice := make([]string, 0)

	m := map[int64]*Creative{
		int64(2222): {
			Id:        pointer_yy.Int64Ptr(2222),
			ImageInfo: pointer_yy.StringPtr("image_info_ddd"),
		},
		int64(3333): {
			Id:        pointer_yy.Int64Ptr(3333),
			ImageInfo: pointer_yy.StringPtr("image_info_ddd_eeee"),
		},
	}

	for id, materialInfo := range m {
		j, _ := json.Marshal(materialInfo)
		structJsonSlice = append(structJsonSlice, fmt.Sprintf("{\"%d\":%s}", id, string(j)))
	}
	// 直接打印mapJson这个slice是不行的，还需要用printf的形式
	println("structJsonSlice=", structJsonSlice)
	fmt.Printf("structJsonSlice=%+v\n", structJsonSlice)
	println(strings.Join(structJsonSlice, ",")) // 这样打印出来的是一个字符串，而不是一个slice

}
