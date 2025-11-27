package piscine

// NodeI represents a single node in the linked list with integer data
type NodeI struct {
	Data int
	Next *NodeI
}

// SortListInsert inserts a new node into a sorted linked list
// so that the list remains sorted in ascending order
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	// Create the new node
	newNode := &NodeI{Data: data_ref}

	// Case 1: empty list or new node should be the new head
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// Case 2: find the correct position in the list
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert the new node after current
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
