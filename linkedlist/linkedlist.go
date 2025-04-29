package linkedlist

// declare a new struct for the Node for 'any' type
type Node[T any] struct {
	Value T
	next  *Node[T]
}

// declare a new struct for the linked list (single-linked)
type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// factory function to create and return a new empty linked list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Length returns number of items in the list
func (l *LinkedList[T]) Length() int {
	return l.length
}

// Head returns the first item in the list
func (l *LinkedList[T]) Head() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.Value, true
}

// Tail returns the last item in the list
func (l *LinkedList[T]) Tail() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.Value, true
}

// Append adds a new item to the end of the list
func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	// if empty assign new node to head and tail
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		// assign newNode to the end of populated list
		l.tail.next = newNode
		l.tail = newNode
	}
	// increase length
	l.length++
}

// Prepend adds a new node with the given value to the begin of list
func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	// if empty assign new node to head and tail
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		// assign newNode to begin of populated list
		newNode.next = l.head
		l.head = newNode
	}
	// increase length
	l.length++
}

// InsertAt inserts a new node with a given value at a specific index
func (l *LinkedList[T]) InsertAt(index int, value T) bool {
	if index < 0 || index > l.length {
		return false
	}

	// Insert at begin
	if index == 0 {
		l.Prepend(value)
		return true
	}

	// Insert at end
	if index == l.length {
		l.Append(value)
	}

	// Insert in the middle
	newNode := &Node[T]{Value: value}
	current := l.head
	// Traverse the node to before the insertion point
	for range index - 1 {
		// out of bounds check
		if current == nil {
			return false
		}
		current = current.next
	}
	// current is now before the insertion point
	newNode.next = current.next
	current.next = newNode

	l.length++
	return true
}

// Get returns the value of the node at the specified index (and ok=true)
// returns zero value and ok=false if the index is out of bounds
func (l *LinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.length {
		var zero T
		return zero, false
	}

	current := l.head
	for range index {
		if current == nil {
			var zero T
			return zero, false
		}
		current = current.next
	}
	return current.Value, true
}

// RemoveAd removes the node at a specific index and returns its value and ok=true
// returns zero value and ok=false if the index is out of bounds
func (l *LinkedList[T]) RemoveAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.length {
		return zero, false
	}

	var removedNode *Node[T]

	// Remove head
	if index == 0 {
		removedNode = l.head
		l.head = l.head.next
	} else {
		// Remove middle or tail
		prev := l.head
		for range index - 1 {
			if prev.next == nil {
				return zero, false
			}
			prev = prev.next
		}
		// prev is now the node before the target node
		removedNode = prev.next
		prev.next = removedNode.next
		// If removing tail
		if removedNode == l.tail {
			l.tail = prev
		}
	}
	l.length--
	// clear pointer of removed node to prevent possible leaks
	removedNode.next = nil
	return removedNode.Value, true
}
