package prodneutral

import "testing"

func TestUsingPaging(t *testing.T) {
	err := usingPaging()
	if err != nil {
		t.Errorf("paging: error: %v\n", err)
	}
}
