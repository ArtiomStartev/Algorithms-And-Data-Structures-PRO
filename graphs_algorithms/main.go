package main

import (
	"algorithms/graphs_algorithms/graphs"
	"fmt"
)

func main() {
	println("Breadth-first search (BFS) algorithm")

	startNode := graphs.DirectedGraph1()
	traversalOrder := graphs.BFS(startNode)
	graphs.PrintTraversalOrder(traversalOrder)
	sumOfNeighbors := graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	startNode = graphs.UndirectedGraph()
	traversalOrder = graphs.BFS(startNode)
	graphs.PrintTraversalOrder(traversalOrder)
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	println("\nDepth-first search (DFS) algorithm")

	startNode = graphs.DirectedGraph1()
	traversalOrder = graphs.DFS(startNode)
	graphs.PrintTraversalOrder(traversalOrder)
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	startNode = graphs.UndirectedGraph()
	traversalOrder = graphs.DFS(startNode)
	graphs.PrintTraversalOrder(traversalOrder)
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)
}
