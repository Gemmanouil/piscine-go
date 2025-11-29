package piscine

type TreeNode2 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeApplyPreorder traverses the tree in preorder and applies f to each node's Data.
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Visit current node first
	f(root.Data)

	// Then traverse left subtree
	BTreeApplyPreorder(root.Left, f)

	// Finally traverse right subtree
	BTreeApplyPreorder(root.Right, f)
}
