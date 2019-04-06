package ch9

import (
	// "gopl.io/ch5/links"
	// "html"
	"log"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
	"golang.org/x/net/html"
)

// Q6 测试
func Q6() {
	fmt.Println("Html Parse")
	fetch()
} 

const requstURL = "http://tup.com.cn/index.html"
func fetch() {
	resp, err := http.Get(requstURL)	
	if  err != nil  {
		log.Fatalf("fetch reading %s: %v \n", requstURL,err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("fetch reading %s: %v \n", requstURL,err)
	}
	fmt.Fprintf(os.Stdin, "%s",b)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findLinks: %v \n", err)
	}
	fmt.Println(doc)
	fmt.Println("The end")
	return
	for _, link := range visit(nil,doc) {
		fmt.Println(link)
	}
	fmt.Println("This is TTT")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
	return links
}