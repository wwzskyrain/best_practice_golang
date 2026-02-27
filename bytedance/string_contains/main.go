package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "string"
	fmt.Printf("%x\n", s[:2])

	TestStringContain()
}

func TestStringContain() {
	bodyStrArr := []string{"DES_REQUEST_INVALID", "DES_IDL_SCHEMA_VALIDATE_ERROR_CONTROL_PLANE_DATA", "DES_GATEWAY_AGENT_RETURNS_ERROR"}
	for _, responseBodyStr := range bodyStrArr {
		if strings.Contains(responseBodyStr, "IDL_") || strings.Contains(responseBodyStr, "DES_") {
			fmt.Printf("%s : contains\n", responseBodyStr)
		} else {
			fmt.Printf("%s : not contains\n", responseBodyStr)
		}
	}
}
