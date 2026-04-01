package prodneutral

import "testing"

func TestGettingLRO(t *testing.T) {
	err := getLROName()
	if err != nil {
		t.Errorf("lro: error: %v\n", err)
	}
}

func TestAccessingLRO(t *testing.T) {
	err := getLRODataFromName()
	if err != nil {
		t.Errorf("lro: error: %v\n", err)
	}
}

func TestGettingError(t *testing.T) {
	err := getErrorCode()
	if err == nil {
		t.Errorf("should have gotten error")
	}
}
