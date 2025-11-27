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

// ListMerge places elements of l2 at the end of l1
func ListMerge(l1 *List, l2 *List) {
	// If l2 is empty, nothing to merge
	if l2.Head == nil {
		return
	}

	// If l1 is empty, l1 becomes l2
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Otherwise, connect l1's tail to l2's head
	l1.Tail.Next = l2.Head
	// Update l1's tail to be l2's tail
	l1.Tail = l2.Tail
}
