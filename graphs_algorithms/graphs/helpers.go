package graphs

// GraphNode represents a single node in a graph
type GraphNode[T any] struct {
	Value     T               // The value of the node
	Neighbors []*GraphNode[T] // A dynamic buffer for neighboring nodes
}

// DirectedGraph1 initializes the first directed graph configuration
func DirectedGraph1() *GraphNode[int] {
	// Creating nodes
	A := &GraphNode[int]{Value: 1}
	B := &GraphNode[int]{Value: 2}
	C := &GraphNode[int]{Value: 3}
	D := &GraphNode[int]{Value: 4}

	// Setting up neighbors
	A.Neighbors = []*GraphNode[int]{B, C, D}
	B.Neighbors = []*GraphNode[int]{C}
	C.Neighbors = []*GraphNode[int]{D}
	D.Neighbors = []*GraphNode[int]{A}

	return A
}

// UndirectedGraph initializes the second undirected graph configuration
func UndirectedGraph() *GraphNode[int] {
	// Creating nodes
	A := &GraphNode[int]{Value: 1}
	B := &GraphNode[int]{Value: 2}
	C := &GraphNode[int]{Value: 3}
	D := &GraphNode[int]{Value: 4}

	// Setting up neighbors (bidirectional)
	A.Neighbors = []*GraphNode[int]{B, D}
	B.Neighbors = []*GraphNode[int]{A, C, D}
	C.Neighbors = []*GraphNode[int]{B, D}
	D.Neighbors = []*GraphNode[int]{A, B, C}

	return A
}

// SumOfNeighbors calculates the sum of values of neighboring nodes
func SumOfNeighbors(node *GraphNode[int]) int {
	var sum int
	for _, neighbor := range node.Neighbors {
		sum += neighbor.Value
	}
	return sum
}

// PrintTraversalOrder prints the traversal order of the graph nodes
func PrintTraversalOrder[T any](traversalOrder []T) {
	for i, value := range traversalOrder {
		print(value)
		if i < len(traversalOrder)-1 {
			print(" -> ")
		}
	}
	println()
}
