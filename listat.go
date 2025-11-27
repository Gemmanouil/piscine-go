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

// ListAt returns the pointer to the node at position pos
func ListAt(l *List, pos int) *NodeL {
	// If the list is empty or pos is negative, return nil
	if l.Head == nil || pos < 0 {
		return nil
	}

	current := l.Head // Start from the head
	index := 0        // Position counter

	// Traverse the list until we reach the desired position
	for current != nil {
		if index == pos {
			// Found the node at the requested position
			return current
		}
		// Move to the next node and increase the counter
		current = current.Next
		index++
	}

	// If pos is out of range, return nil
	return nil
}
