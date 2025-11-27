package piscine

type TreeNode1 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeApplyPostorder traverses the tree in postorder and applies f to each node's Data.
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Traverse left subtree
	BTreeApplyPostorder(root.Left, f)

	// Traverse right subtree
	BTreeApplyPostorder(root.Right, f)

	// Apply function to current node's Data
	f(root.Data)
}
