package main

import "testing"

func TestAddNode(t *testing.T) {
	ll := LinkedList{}
	want := 5

	ll.Add(want)
	got := ll.Head.Datum

	if got != want {
		t.Errorf("data mismatch: wanted %v, got %v", want, got)
	}
}
