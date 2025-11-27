package piscine

// ListClear deletes all nodes from the linked list
func ListClear(l *List) {
	// Set Head to nil so the list no longer points to the first node
	l.Head = nil
	// Set Tail to nil so the list no longer points to the last node
	l.Tail = nil
	// At this point, the list is empty and all nodes will be garbage collected
}
