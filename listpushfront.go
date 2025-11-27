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

// ListPushFront inserts a new node at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	// Create a new node with the given data
	newNode := &NodeL{Data: data}

	// If the list is empty, both Head and Tail point to the new node
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, link the new node to the current Head
		newNode.Next = l.Head
		// Update Head to be the new node
		l.Head = newNode
	}
}
