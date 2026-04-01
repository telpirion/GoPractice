package misc

import "testing"

func TestSleepingBarbers(t *testing.T) {
	err := barbers(3, 3, 10, 100)
	if err != nil {
		t.Error(err)
	}
}
