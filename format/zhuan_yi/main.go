package main

import "fmt"

func main() {
	tableStr := fmt.Sprintf("<table width = '100%%' border = \"1\" cellspacing=\"0\" cellpadding=\"10px\">%s</table>", "tableBodyStr")
	fmt.Println(tableStr)
}
