package ch9

import (
	// "gopl.io/ch5/links"
	// "html"
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

const requstURL = "http://www.baidu.com"
func fetch() {
	resp, err := http.Get(requstURL)	
	if  err != nil  {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n",requstURL,err)	
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch reading %s: %v \n", requstURL,err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdin, "%s",b)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLins: %v \n", err)
	}
	for _, link := range visit(nil,doc) {
		fmt.Println(link)
	}
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