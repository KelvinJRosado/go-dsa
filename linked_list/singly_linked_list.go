package linked_list

// Defines a node that will be used in the linked list
// Uses a pointer to the next node to avoid self-recursice struct definition
type SingleNode[T any] struct {
	Data T
	next *SingleNode[T]
}

// Returns the next element pointed to by the node
func (sn *SingleNode[T]) Next() SingleNode[T] {
	return *sn.next
}

// A singly linked list
type SinglyLinkedList[T any] struct {
	root *SingleNode[T]
}

// Get the first element in the list
func (sl *SinglyLinkedList[T]) GetRoot() SingleNode[T] {
	return *sl.root
}

// Inserts a node as the new root
func (sl *SinglyLinkedList[T]) InsertNewRoot(val T) {

	newRoot := SingleNode[T]{
		Data: val,
		next: nil,
	}

	// Special case: Replaces nil root that gets created on init
	if sl.root == nil {
		sl.root = &newRoot
		return
	}

	oldRoot := sl.GetRoot()
	newRoot.next = &oldRoot
}

// Inserts a new node into the list, after an existing node
// Updates `next` pointers to maintain correctness
func (sl *SinglyLinkedList[T]) InsertAfterNode(node *SingleNode[T], newNode *SingleNode[T]) {
	newNode.next = node.next
	node.next = newNode
}

// Removes the root element
func (sl *SinglyLinkedList[T]) RemoveRoot() {

	// Skip if empty

	newRoot := sl.GetRoot().next
	sl.root = newRoot
}
