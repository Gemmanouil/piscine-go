package piscine

type Treenode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// helper function to check BST properties with bounds
func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}

	// Check current node against bounds
	if min != nil && node.Data <= *min {
		return false
	}
	if max != nil && node.Data > *max {
		return false
	}

	// Left subtree: max bound is current node's Data
	if !isBST(node.Left, min, &node.Data) {
		return false
	}
	// Right subtree: min bound is current node's Data
	if !isBST(node.Right, &node.Data, max) {
		return false
	}

	return true
}

// BTreeIsBinary returns true if the tree follows BST properties
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}
