package piscine

// NodeI represents a single node in the linked list with integer data
type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the linked list in ascending order
func ListSort(l *NodeI) *NodeI {
	// If the list is empty or has only one element, nothing to sort
	if l == nil || l.Next == nil {
		return l
	}

	swapped := true
	for swapped {
		swapped = false
		current := l
		for current.Next != nil {
			// Compare current node with the next node
			if current.Data > current.Next.Data {
				// Swap the values
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}
	return l
}
