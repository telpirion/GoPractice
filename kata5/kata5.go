package main

import "fmt"

func main() {
	fmt.Println("hello")
}

type LinkedListNode struct {
	Datum interface{}
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) Add(d interface{}) {
	ln := LinkedListNode{Datum: d}

	if l.Head == nil {
		l.Head = &ln
	}

	tail := iterateToEnd(*l.Head)
	tail.Next = &ln
}

func iterateToEnd(n LinkedListNode) *LinkedListNode {
	if n.Next == nil {
		return &n
	}

	return iterateToEnd(*n.Next)
}
