package main

import (
	"fmt"

	toposort "github.com/Devang-25/go-topological-sort"
)

func main() {
	graph := toposort.NewGraph(8)
	graph.AddNodes("2", "3", "5", "7", "8", "9", "10", "11", "12")

	graph.AddEdge("7", "8")
	graph.AddEdge("7", "11")

	graph.AddEdge("5", "11")

	graph.AddEdge("3", "8")
	graph.AddEdge("3", "10")

	graph.AddEdge("11", "2")
	graph.AddEdge("11", "9")
	graph.AddEdge("11", "10")

	graph.AddEdge("8", "9")

	result, ok := graph.Toposort()
	if !ok {
		panic("cycle detected")
	}

	fmt.Println(result)
}
