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
