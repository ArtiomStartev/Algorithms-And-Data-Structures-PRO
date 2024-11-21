package main

import (
	"algorithms/graphs_algorithms/graphs"
	"fmt"
)

func main() {
	// Breadth-first search (BFS) algorithm
	println("Breadth-first search (BFS) algorithm")

	// Create a directed graph
	startNode := graphs.DirectedGraph1()
	// Perform a breadth-first search on the directed graph
	traversalOrder := graphs.BFS(startNode)
	// Print the traversal order
	graphs.PrintTraversalOrder(traversalOrder)

	// Calculate the sum of neighbors of the start node
	sumOfNeighbors := graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	// Create an undirected graph
	startNode = graphs.UndirectedGraph()
	// Perform a breadth-first search on the undirected graph
	traversalOrder = graphs.BFS(startNode)
	// Print the traversal order
	graphs.PrintTraversalOrder(traversalOrder)

	// Calculate the sum of neighbors of the start node
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	// Depth-first search (DFS) algorithm
	println("\nDepth-first search (DFS) algorithm")

	// Create a directed graph
	startNode = graphs.DirectedGraph1()
	// Perform a depth-first search on the directed graph
	traversalOrder = graphs.DFS(startNode)
	// Print the traversal order
	graphs.PrintTraversalOrder(traversalOrder)

	// Calculate the sum of neighbors of the start node
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)

	// Create an undirected graph
	startNode = graphs.UndirectedGraph()
	// Perform a depth-first search on the undirected graph
	traversalOrder = graphs.DFS(startNode)
	// Print the traversal order
	graphs.PrintTraversalOrder(traversalOrder)

	// Calculate the sum of neighbors of the start node
	sumOfNeighbors = graphs.SumOfNeighbors(startNode)
	fmt.Printf("Node value: %d, Sum of neighbors: %d\n", startNode.Value, sumOfNeighbors)
}
