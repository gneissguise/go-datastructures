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
