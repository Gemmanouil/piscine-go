package piscine

// TreeNode definition
type TreeNode0 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeApplyInorder traverses the tree in-order and applies f to each node's Data.
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Traverse left subtree
	BTreeApplyInorder(root.Left, f)

	// Apply function to current node's Data
	f(root.Data)

	// Traverse right subtree
	BTreeApplyInorder(root.Right, f)
}
