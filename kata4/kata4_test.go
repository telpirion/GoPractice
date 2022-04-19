package main

import "testing"

func TestPassMessagesBetweenGoroutines(t *testing.T) {
	passMessagesBetweenGoroutines("Hello there!")
}
