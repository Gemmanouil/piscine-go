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

// ListRemoveIf removes all nodes whose Data equals data_ref
func ListRemoveIf(l *List, data_ref interface{}) {
	// First, handle the case where the head nodes need to be removed
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If the list becomes empty, reset Tail
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Traverse the list with two pointers: prev and current
	prev := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Data == data_ref {
			// Skip the current node
			prev.Next = current.Next
			// If current was the tail, update Tail
			if current == l.Tail {
				l.Tail = prev
			}
		} else {
			// Move prev forward only if we didn't remove current
			prev = current
		}
		// Move current forward
		current = current.Next
	}
}
