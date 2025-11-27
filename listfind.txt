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

// CompStr compares two interface{} values for equality
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind returns the address of the Data of the first node
// that matches ref according to the comparison function comp
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	// Start from the head of the list
	current := l.Head

	// Traverse the list
	for current != nil {
		// Use the comparison function to check equality
		if comp(current.Data, ref) {
			// Return the address of the Data field
			return &current.Data
		}
		// Move to the next node
		current = current.Next
	}

	// If no match found, return nil
	return nil
}
