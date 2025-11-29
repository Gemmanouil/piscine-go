package piscine

type Treenode001 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeApplyByLevel traverses the tree level by level (breadth-first)
// and applies f to each node's Data.
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// simple queue implemented with a slice
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// take the first element
		current := queue[0]
		queue = queue[1:]

		// apply function to current node's Data
		f(current.Data)

		// enqueue children
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
