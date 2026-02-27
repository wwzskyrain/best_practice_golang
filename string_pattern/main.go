package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

type MyNode struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
}

func main() {
	// zh := "如果您希望继续使用Pangle进行变现，请您请在14天内点击<a href=\"https://www.pangleglobal.com/zh/knowledge/28073\">此链接</a>进行整改，否则将停止填充。"
	// en := "If you’re keen on continuing to monetize with Pangle and wish to resolve this matter, Please click on this <a href=\"https://www.pangleglobal.com/zh/knowledge/28073\">link</a> to rectify within 14 days, otherwise the filling will be halted."
	// print(ParseTextIncludeSuperLink(zh))
	// print(ParseTextIncludeSuperLink(en))

	// test1()
	// test2()
	// test3()
	// myNodes := ParseTextIncludeSuperLink("aaabbbccc")
	// for _, node := range myNodes {
	// 	fmt.Printf("node=%+v\n", node)
	// }
	test05()
}

func test3() {
	htmlString := "如果您希望继续使用Pangle进行变现，请您请在14天内点击<a href=\"https://www.pangleglobal.com/zh/knowledge/28073\">此链接</a>进行整改，否则将停止填充。"

	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	var foundLink string
	var findALinks func(n *html.Node, level int)
	findALinks = func(n *html.Node, level int) {
		fmt.Printf("L[%d] n.Type=%d, n.Data=%s\n", level, n.Type, n.Data)
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					foundLink = a.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findALinks(c, level+1)
		}
	}

	findALinks(doc, 0)
	fmt.Println("Found Link:", foundLink)
}

func ParseTextIncludeSuperLink(text string) []MyNode {

	idx := strings.Index(text, "<a")
	if idx == -1 {
		return []MyNode{{Text: text}}
	}
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}
	var findBodyNode func(node *html.Node) (*html.Node, error)
	findBodyNode = func(node *html.Node) (*html.Node, error) {
		if node.Type == html.ElementNode && node.Data == "body" {
			return node, nil
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			bodyNode, err := findBodyNode(c)
			if err == nil {
				return bodyNode, nil
			}
		}
		return nil, fmt.Errorf("body node not found")
	}
	bodyNode, err := findBodyNode(doc)
	if err != nil {
		fmt.Println("Error finding body node:", err)
		return nil
	}
	result := make([]MyNode, 0)
	for c := bodyNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result = append(result, MyNode{Type: "text", Text: c.Data})
		} else if c.Type == html.ElementNode && c.Data == "a" {
			href := ""
			for _, attr := range c.Attr {
				if attr.Key == "href" {
					href = attr.Val
				}
			}
			result = append(result, MyNode{Type: "link", Text: c.FirstChild.Data, Url: href})
		}
	}
	return result
}

func test1() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func test2() {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func test05() {
	tableStr := "hahaha"
	fmt.Printf(`<p class="union-media-editor__paragraph">%s</p>`, tableStr)
	fmt.Printf("<p class=\"union-media-editor__paragraph\">%s</p>", tableStr)
}
