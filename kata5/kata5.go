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

func iterateToValue(n *LinkedListNode, val interface{}) *LinkedListNode {
	if n.Next == nil {
		return nil
	}

	if n.Next.Datum == val {
		return n
	}

	return iterateToValue(n.Next, val)
}

func (l *LinkedList) Remove(val interface{}) *LinkedListNode {
	if l.Head == nil {
		return nil
	}

	n := iterateToValue(l.Head, val)
	next := n.Next
	n.Next = next.Next
	return next
}

func (l *LinkedList) Add(d interface{}) {
	ln := LinkedListNode{Datum: d}

	if l.Head == nil {
		l.Head = &ln
	}

	tail := iterateToEnd(l.Head, nil)
	tail.Next = &ln
}

func (l *LinkedList) Print() {

	if l.Head == nil {
		return
	}

	print := func(n *LinkedListNode) {
		fmt.Println(n.Datum)
	}

	iterateToEnd(l.Head, &print)
}

func iterateToEnd(n *LinkedListNode, f *func(*LinkedListNode)) *LinkedListNode {

	if f != nil {
		(*f)(n)
	}

	if n.Next == nil {
		return n
	}

	return iterateToEnd(n.Next, f)
}
