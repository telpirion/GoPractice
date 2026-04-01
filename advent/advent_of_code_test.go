package misc

import (
	"testing"
)

func TestDay1(t *testing.T) {
	day1File := "resources/day1.txt"
	tot, _, err := doDay1(day1File)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := 66306

	if tot != 66306 {
		t.Errorf("Expected %d, got %v", want, tot)
	}
}

func TestDay2(t *testing.T) {
	day2File := "resources/day2.txt"
	tot, err := doDay2(day2File)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := 13565

	if tot != want {
		t.Errorf("Expected %d, got %v", want, tot)
	}
}
