package prodneutral

import "testing"

func TestWithADC(t *testing.T) {
	err := withADCCredentials()

	if err != nil {
		t.Errorf("auth: couldn't construct client: %v", err)
	}
}

func TestWithManualCredentials(t *testing.T) {
	err := withManualCredentials("/Users/erschmid/Devtools/vidint-erschmid.json")
	if err != nil {
		t.Errorf("auth: couldn't construct client: %v", err)
	}
}
