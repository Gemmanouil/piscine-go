package piscine

// NodeL represents a single node in the linked list
type NodeL struct {
	Data interface{} // Data can hold any type
	Next *NodeL      // Pointer to the next node
}

// List represents the linked list itself
type List struct {
	Head *NodeL // Pointer to the first node
	Tail *NodeL // Pointer to the last node
}

// IsPositiveNode returns true if the node's Data is a positive number
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return node.Data.(int) > 0
	case float32:
		return node.Data.(float32) > 0
	case float64:
		return node.Data.(float64) > 0
	case byte:
		return node.Data.(byte) > 0
	default:
		return false
	}
}

// IsAlNode returns true if the node's Data is not a number (i.e., is alphanumeric/string)
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// ListForEachIf applies function f to nodes that satisfy condition cond
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	// Start from the head of the list
	current := l.Head
	// Traverse until the end
	for current != nil {
		// Apply f only if cond returns true
		if cond(current) {
			f(current)
		}
		// Move to the next node
		current = current.Next
	}
}
