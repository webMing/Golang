package ch9

import (
	// "gopl.io/ch5/links"
	// "html"
	"log"
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

// Q6 测试
func Q6() {
	fmt.Println("Html Parse")
	fetch()
} 

const requstURL = "https://golang.org"
func fetch() {
	resp, err := http.Get(requstURL)	
	if  err != nil  {
		log.Fatalf("fetch reading %s: %v \n", requstURL,err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("findLinks: %v \n", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		Log.Fatalf("findLinks: %v \n", err)
	}
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
		links = visit(links, c)
	}
	return links
}