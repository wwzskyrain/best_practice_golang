package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func main() {
	diffJson()
}
func diffJson() {
	dmp := diffmatchpatch.New()
	source := `{
    "A": "bar",
    "B": "baz",
    "C": "foo"
}`
	target := `{
    "A": "rab",
    "B": "baz",
    "D": "foo"
}`
	diffs := dmp.DiffMain(source, target, false)

	h := dmp.DiffPrettyHtml(diffs)
	fmt.Println(h)
}

func diffTextInLine() {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	r := dmp.DiffPrettyText(diffs)
	fmt.Println(r)

	h := dmp.DiffPrettyHtml(diffs)
	fmt.Println(h)
}
