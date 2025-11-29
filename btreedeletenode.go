package piscine

type TreeNode88 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeDeleteNode deletes 'node' from the tree rooted at 'root'.
// Returns the (possibly new) root of the tree.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1: no left child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		// Case 2: no right child
		root = BTreeTransplant(root, node, node.Left)
	} else {
		// Case 3: two children
		successor := BTreeMin(node.Right)
		if successor.Parent != node {
			// Replace successor with its right child
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}
		// Replace node with successor
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}

	return root
}
