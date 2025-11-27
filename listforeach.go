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

// ListForEach applies the given function f to each node of the list
func ListForEach(l *List, f func(*NodeL)) {
	// Start from the head of the list
	current := l.Head
	// Traverse the list until the end
	for current != nil {
		// Apply the function to the current node
		f(current)
		// Move to the next node
		current = current.Next
	}
}

// Add2_node modifies the node's Data by adding 2 (if int) or appending "2" (if string)
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node modifies the node's Data by subtracting 3 (if int) or appending "-3" (if string)
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
