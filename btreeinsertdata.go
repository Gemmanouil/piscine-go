package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node with the given data into the BST.
// If root is nil, it creates and returns a new root node.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// If the tree is empty, create the root node
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Start from the root and traverse down
	current := root
	for {
		if data < current.Data {
			// Go left
			if current.Left == nil {
				// Insert new node here
				current.Left = &TreeNode{Data: data, Parent: current}
				break
			}
			current = current.Left
		} else {
			// Go right (equal values also go right)
			if current.Right == nil {
				// Insert new node here
				current.Right = &TreeNode{Data: data, Parent: current}
				break
			}
			current = current.Right
		}
	}

	return root
}
