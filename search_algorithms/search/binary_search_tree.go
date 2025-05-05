package search

type BinarySearchTreeSearcher struct{}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

// Search performs a binary search on a Binary Search Tree
func (BinarySearchTreeSearcher) Search(data any, x int) bool {
	bst, ok := data.(*BinarySearchTree)
	if !ok {
		return false
	}
	return searchNode(bst.Root, x)
}

// searchNode is a helper function to search for a value in the tree
func searchNode(node *Node, x int) bool {
	if node == nil {
		return false
	}

	if x == node.Value {
		return true
	} else if x < node.Value {
		return searchNode(node.Left, x)
	} else {
		return searchNode(node.Right, x)
	}
}

// Insert inserts a new value into the Binary Search Tree
func (bst *BinarySearchTree) Insert(val int) {
	bst.Root = insertNode(bst.Root, val)
}

// insertNode is a helper function to insert a new value into the tree
func insertNode(node *Node, val int) *Node {
	if node == nil {
		return &Node{Value: val}
	}

	if val < node.Value {
		node.Left = insertNode(node.Left, val)
	} else {
		node.Right = insertNode(node.Right, val)
	}

	return node
}

// MapSliceToTree creates a Binary Search Tree from an array of integers
func MapSliceToTree(arr []int) *BinarySearchTree {
	bst := &BinarySearchTree{}
	for _, val := range arr {
		bst.Insert(val)
	}
	return bst
}
