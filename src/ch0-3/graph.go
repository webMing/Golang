package main

import (
	"fmt"
	)

func main () {
	fmt.Println("=======================")
	
	var graph = make(map[string]map(string)bool)

	fmt.Println("=======================")
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}


func hasEdge(from, to stirng) bool {
	return graph[from][to]
}
