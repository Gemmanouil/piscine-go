package piscine

// NodeL represents a single node in the linked list
type NodeL struct {
	Data any    // Data can hold any type
	Next *NodeL // Pointer to the next node
}

// ListAt returns the pointer to the node at position pos
func ListAt(l *NodeL, pos int) *NodeL {
	// If the list is empty or pos is negative, return nil
	if l == nil || pos < 0 {
		return nil
	}

	current := l // Start from the given head node
	index := 0   // Position counter

	// Traverse until we reach the desired position
	for current != nil {
		if index == pos {
			return current // Found the node at position pos
		}
		current = current.Next // Move to next node
		index++
	}

	// If pos is out of range, return nil
	return nil
}
