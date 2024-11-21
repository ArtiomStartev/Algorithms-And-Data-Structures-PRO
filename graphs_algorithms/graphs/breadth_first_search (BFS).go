package graphs

// BFS performs a breadth-first search on the graph
func BFS[T any](startNode *GraphNode[T]) []T {
	// Initialize an empty queue for BFS traversal
	var queue []*GraphNode[T]

	// Initialize a map to keep track of visited nodes
	visited := make(map[*GraphNode[T]]bool)

	// Initialize a slice to store the traversal order
	var traversalOrder []T

	// Add the start node to the queue and mark it as visited
	queue = append(queue, startNode)
	visited[startNode] = true

	// Perform BFS traversal
	for len(queue) > 0 {
		// Dequeue a node from the queue
		currentNode := queue[0]
		queue = queue[1:]

		// Add the node's value to the traversal order
		traversalOrder = append(traversalOrder, currentNode.Value)

		// Iterate over the neighbors of the current node
		for _, neighbor := range currentNode.Neighbors {
			// If the neighbor has not been visited, mark it as visited and enqueue it
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return traversalOrder
}
