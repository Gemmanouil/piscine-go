package piscine

type TreeNode031 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeMin returns the node with the minimum value in the tree.
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
