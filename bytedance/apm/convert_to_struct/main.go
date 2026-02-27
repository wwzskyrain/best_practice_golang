package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pratice/golang/util"
)

type StackEntry struct {
	LineNo   int    `json:"line_no"`
	LibName  string `json:"lib_name"`
	BaseAddr string `json:"base_add"`
	Offset   string `json:"offset"`
}

func parseStack(stack string) []StackEntry {
	var entries []StackEntry
	// 正则匹配每一行的格式
	re := regexp.MustCompile(`(?m)^\s*(\d+)\s+([\w\.]+)\s+0x([0-9a-fA-F]+)\s+0x([0-9a-fA-F]+) \+ \d+.*$`)
	matches := re.FindAllStringSubmatch(stack, -1)

	for _, match := range matches {
		if len(match) < 5 {
			continue
		}

		lineNo, _ := strconv.Atoi(match[1])
		entries = append(entries, StackEntry{
			LineNo:   lineNo,
			LibName:  match[2],
			BaseAddr: "0x" + match[4],
			Offset:   "0x" + match[3],
		})
	}

	return entries
}

func main() {
	iosStack := `
0   CoreFoundation                  0x000000019a5dd9d4 0x19a4c3000 + 1157588 ((null)) + 0)
1   libobjc.A.dylib                 0x00000001adf8eb54 0x1adf88000 + 27476 ((null)) + 0)
2   CoreFoundation                  0x000000019a4edbac 0x19a4c3000 + 175020 ((null)) + 0)
3   CoreFoundation                  0x000000019a5e0018 0x19a4c3000 + 1167384 ((null)) + 0)
4   CoreFoundation                  0x000000019a5e1f8c 0x19a4c3000 + 1175436 ((null)) + 0)
5   BUDemoDev                       0x00000001045eaa38 0x10404c000 + 5892664 ((null)) + 0)
6   BUDemoDev                       0x00000001045ea9e4 0x10404c000 + 5892580 ((null)) + 0)
7   BUDemoDev                       0x00000001045ea9b8 0x10404c000 + 5892536 ((null)) + 0)
8   BUDemoDev                       0x00000001045ea98c 0x10404c000 + 5892492 ((null)) + 0)
9   libdispatch.dylib               0x0000000105489528 0x105484000 + 21800 ((null)) + 0)
10  libdispatch.dylib               0x000000010548c1cc 0x105484000 + 33228 ((null)) + 0)
11  libdispatch.dylib               0x000000010549fa8c 0x105484000 + 113292 ((null)) + 0)
12  libdispatch.dylib               0x00000001054977fc 0x105484000 + 79868 ((null)) + 0)
13  CoreFoundation                  0x000000019a55d5d0 0x19a4c3000 + 632272 ((null)) + 0)
14  CoreFoundation                  0x000000019a557a78 0x19a4c3000 + 608888 ((null)) + 0)
15  CoreFoundation                  0x000000019a556b90 0x19a4c3000 + 605072 ((null)) + 0)
16  GraphicsServices                0x00000001b0879598 0x1b0876000 + 13720 ((null)) + 0)
17  UIKitCore                       0x000000019ce40638 0x19c314000 + 11716152 ((null)) + 0)
18  UIKitCore                       0x000000019ce45bb8 0x19c314000 + 11738040 ((null)) + 0)
19  BUDemoDev                       0x0000000104051864 0x10404c000 + 22628 ((null)) + 0)
20  libdyld.dylib                   0x000000019a235588 0x19a234000 + 5512 ((null)) + 0)`

	entries := parseStack(iosStack)
	fmt.Printf("look:\n%s", util.Struct2String(entries))
}
