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

// ListLast returns the Data of the last element in the list
func ListLast(l *List) interface{} {
	// If the list is empty, return nil
	if l.Head == nil {
		return nil
	}
	// Otherwise, return the Data of the Tail node
	return l.Tail.Data
}
