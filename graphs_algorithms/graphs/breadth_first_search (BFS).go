package graphs

// BFS performs a breadth-first search (BFS) on a graph represented by nodes
func BFS[T any](startNode *GraphNode[T]) []T {
	var traversalOrder []T
	var queue []*GraphNode[T]
	visited := make(map[*GraphNode[T]]struct{})

	queue = append(queue, startNode)
	visited[startNode] = struct{}{}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		traversalOrder = append(traversalOrder, currNode.Value)

		for _, neighbor := range currNode.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = struct{}{}
			}
		}
	}

	return traversalOrder
}
