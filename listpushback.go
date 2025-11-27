package piscine

// NodeL represents a single node in the linked list
type NodeL struct {
	Data interface{} // Data can hold any type of value
	Next *NodeL      // Pointer to the next node in the list
}

// List represents the linked list itself
type List struct {
	Head *NodeL // Pointer to the first node of the list
	Tail *NodeL // Pointer to the last node of the list
}

// ListPushBack inserts a new node at the end of the list
func ListPushBack(l *List, data interface{}) {
	// Create a new node with the given data
	newNode := &NodeL{Data: data}

	// If the list is empty (no Head yet)
	if l.Head == nil {
		// The new node becomes both Head and Tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If the list already has elements:
		// Link the current Tail to the new node
		l.Tail.Next = newNode
		// Update Tail to be the new node
		l.Tail = newNode
	}
}
