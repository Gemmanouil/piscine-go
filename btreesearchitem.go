package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeSearchItem searches for a node with Data == elem.
// Returns the node if found, otherwise returns nil.
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	// Compare elem with current node's Data
	if elem == root.Data {
		return root
	} else if elem < root.Data {
		// Search left subtree
		return BTreeSearchItem(root.Left, elem)
	} else {
		// Search right subtree
		return BTreeSearchItem(root.Right, elem)
	}
}
