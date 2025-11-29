package piscine

type TreeNode09 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeTransplant replaces the subtree rooted at 'node' with 'rplc'.
// Returns the (possibly new) root of the tree.
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// If node is the root
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// node is a left child
		node.Parent.Left = rplc
	} else {
		// node is a right child
		node.Parent.Right = rplc
	}

	// Update parent pointer of rplc
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
