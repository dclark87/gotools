package linkedlist

import (
	"fmt"
)

// Node is a node in the linked list containing a Value.
type Node struct {
	next, prev *Node
	value      interface{}
}

// List is represents the collection of nodes in the linked list.
type List struct {
	head, tail *Node
	len        int
}

// Next returns the next node in the list.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node in the list.
func (n *Node) Prev() *Node {
	return n.prev
}

// Init initializes and returns a new List l.
func (l *List) Init() *List {
	l.head = nil
	l.tail = nil
	l.len = 0

	return l
}

// New instantiates and initializes a new List.
func NewList() *List {
	return new(List).Init()
}

// Len returns the total length of l.
func (l *List) Len() int {
	return l.len
}

// Append adds a value of v to the end of the list l.
func (l *List) Append(v interface{}) {
	new := &Node{
		next:  nil,
		prev:  l.tail,
		value: v,
	}
	if l.head == nil {
		l.head = new
	}
	if l.tail == nil {
		l.tail = new
	} else {
		l.tail.next = new
	}
	l.tail = new
}

func (l *List) Pop() interface{} {
	if l.tail == nil {
		fmt.Print("list is empty!")
		return 0
	}
	node := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}

	return node.value
}

func (l *List) Print() {
	node := l.head
	fmt.Println("list contents:")
	fmt.Println("head:", l.head.value)
	fmt.Println("tail:", l.tail.value)
	for node != nil {
		fmt.Print("{", node.value, "}..")
		node = node.next
	}
	fmt.Println("")
}
