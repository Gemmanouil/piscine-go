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

// ListSize returns the number of elements in the list
func ListSize(l *List) int {
	count := 0           // Initialize counter
	current := l.Head    // Start from the head of the list
	for current != nil { // Traverse until the end of the list
		count++                // Increase counter for each node
		current = current.Next // Move to the next node
	}
	return count // Return the total number of nodes
}
