package main

import "testing"

func TestOpenReadFile(t *testing.T) {
	tf := "../testdata/input.txt"
	s := "test line!!"
	err := openReadFile(tf, s)

	if err != nil {
		t.Error(err)
	}
}
