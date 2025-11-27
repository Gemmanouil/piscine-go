package piscine

// NodeL represents a single node in the linked list
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents the linked list itself
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListAt returns the pointer to the node at position pos
func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}
	return nil
}
