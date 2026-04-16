package graphs

// DFS performs a depth-first search (DFS) on a graph represented by nodes
func DFS[T any](startNode *GraphNode[T]) []T {
	var traversalOrder []T
	visited := make(map[*GraphNode[T]]struct{})

	return dfsTraversal(startNode, visited, traversalOrder)
}

// dfsTraversal recursively traverses the graph in depth-first order
func dfsTraversal[T any](node *GraphNode[T], visited map[*GraphNode[T]]struct{}, traversalOrder []T) []T {
	visited[node] = struct{}{}
	traversalOrder = append(traversalOrder, node.Value)

	for _, neighbor := range node.Neighbors {
		if _, ok := visited[neighbor]; !ok {
			traversalOrder = dfsTraversal(neighbor, visited, traversalOrder)
		}
	}

	return traversalOrder
}
