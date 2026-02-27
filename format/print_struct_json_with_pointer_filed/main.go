package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Num struct {
	N int
}

type Packed struct {
	ID        int64
	PNum      *Num
	Name      string
	NameStar  *string
	ImageInfo *string
}

// PrintMap 打印key=int， value = interface{}
func PrintMap(m map[int]interface{}) string {
	jsonStrList := make([]string, 0)
	for id, materialInfo := range m {
		j, _ := json.Marshal(materialInfo)
		jsonStrList = append(jsonStrList, fmt.Sprintf("{\"%d\":%s}", id, string(j)))
	}
	return fmt.Sprintf("[%s]", strings.Join(jsonStrList[:], ","))
}

func generateMap() map[int]*Packed {
	num := &Num{N: 100}
	nameStar := "nameStar-hhhh"
	return map[int]*Packed{
		333: {
			ID:   int64(333),
			PNum: num, Name: "xx-packed-xy", NameStar: &nameStar, ImageInfo: &nameStar,
		},
		222: {
			ID:   int64(222),
			PNum: num, Name: "xx-packed-xy222", NameStar: &nameStar, ImageInfo: &nameStar,
		},
	}
}

func main() {
	// 	测试
	m := generateMap()
	mm := make(map[int]interface{})
	for i, p := range m {
		mm[i] = p
	}
	str := PrintMap(mm)
	println("str = ", str)
}
