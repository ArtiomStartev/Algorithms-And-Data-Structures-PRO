package search

type BinarySearchTreeSearcher struct{}

// Node represents a single node in the Binary Search Tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinarySearchTree represents the structure of a Binary Search Tree
type BinarySearchTree struct {
	Root *Node
}

// Search searches for the target in the Binary Search Tree and returns true if the target is found, otherwise false
func (BinarySearchTreeSearcher) Search(data any, target int) bool {
	// Type assertion to convert the interface{} type to a BinarySearchTree
	bst, ok := data.(*BinarySearchTree)
	if !ok {
		return false
	}

	return searchNode(bst.Root, target)
}

// searchNode is a helper function that recursively searches for a value in the Binary Search Tree
func searchNode(node *Node, target int) bool {
	// If the current node is nil, return false
	if node == nil {
		return false
	}

	// If the target is found in the current node, return true
	if target == node.Value {
		return true
	} else if target < node.Value { // If the target is less than the current node's value, search in the left subtree
		return searchNode(node.Left, target)
	} else { // Otherwise, search in the right subtree
		return searchNode(node.Right, target)
	}
}

// Insert adds a new value into the Binary Search Tree
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = insertNode(bst.Root, value)
}

// insertNode is a helper function that recursively inserts a value into the Binary Search Tree
func insertNode(node *Node, value int) *Node {
	// If the current node is nil, create a new node with the given value
	if node == nil {
		return &Node{Value: value}
	}

	// If the value is less than the current node's value, insert it into the left subtree
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else { // Otherwise, insert it into the right subtree
		node.Right = insertNode(node.Right, value)
	}

	return node // Return the current node
}

// MapSliceToTree converts a slice of integers into a Binary Search Tree
func MapSliceToTree(slice []int) *BinarySearchTree {
	bst := &BinarySearchTree{}
	for _, value := range slice {
		bst.Insert(value)
	}

	return bst
}
