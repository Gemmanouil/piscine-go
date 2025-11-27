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

// ListReverse reverses the order of the elements in the list
func ListReverse(l *List) {
	// If the list is empty or has only one element, nothing to do
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev *NodeL   // Previous node (initially nil)
	current := l.Head // Start from the head
	l.Tail = l.Head   // After reversal, the old head becomes the new tail

	for current != nil {
		next := current.Next // Save the next node
		current.Next = prev  // Reverse the link
		prev = current       // Move prev forward
		current = next       // Move current forward
	}

	// At the end, prev points to the new head
	l.Head = prev
}
