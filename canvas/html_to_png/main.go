package main

import (
	"fmt"
	"os"

	"github.com/pdfcrowd/pdfcrowd-go"
)

func main() {
	// create the API client instance
	client := pdfcrowd.NewHtmlToImageClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	// configure the conversion
	client.SetOutputFormat("png")

	// run the conversion and write the result to a file
	err := client.ConvertStringToFile("<html><body><h1>Hello World!</h1></body></html>", "HelloWorld2.png")

	// check for the conversion error
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		why, ok := err.(pdfcrowd.Error)
		if ok {
			os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
		} else {
			os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
		}

		panic(err.Error())
	}
}
