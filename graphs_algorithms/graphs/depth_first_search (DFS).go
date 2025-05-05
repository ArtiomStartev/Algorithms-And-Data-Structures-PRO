package graphs

// DFS performs a depth-first search on the graph
func DFS[T any](startNode *GraphNode[T]) []T {
	// Initialize a map to keep track of visited nodes
	visited := make(map[*GraphNode[T]]bool)

	// Initialize a slice to store the traversal order
	var traversalOrder []T

	// Perform DFS traversal starting from the start node
	return dfsTraversal[T](startNode, visited, traversalOrder)
}

func dfsTraversal[T any](startNode *GraphNode[T], visited map[*GraphNode[T]]bool, traversalOrder []T) []T {
	visited[startNode] = true                                // Mark the current node as visited
	traversalOrder = append(traversalOrder, startNode.Value) // Add the current node's value to the traversal order

	// Visit all adjacent nodes that have not been visited
	for _, neighbor := range startNode.Neighbors {
		if !visited[neighbor] {
			traversalOrder = dfsTraversal[T](neighbor, visited, traversalOrder)
		}
	}

	// Return the traversal order
	return traversalOrder
}
