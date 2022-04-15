package main

import (
	"testing"
	"time"
)

func TestOpenReadUpdateJSON(t *testing.T) {
	f := "../testdata/kata2_input.json"
	n := time.Now()
	s := n.Format(time.Stamp)
	k := "test"
	err := openReadUpdateJSON(f, k, s)
	if err != nil {
		t.Error(err)
	}
}
