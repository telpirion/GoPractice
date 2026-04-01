package prodneutral

import "testing"

func TestWithContext(t *testing.T) {
	err := usingContextTimeout(2)
	if err != nil {
		t.Errorf("timeouts: couldn't set timeout: %v", err)
	}
}

func TestWithCallOptions(t *testing.T) {
	err := usingCallOptions(2)
	if err != nil {
		t.Errorf("timeouts: couldn't set timeout: %v", err)
	}
}

func TestSettingKeepAlive(t *testing.T) {
	err := settingKeepAlive(2)
	if err != nil {
		t.Errorf("timeouts: couldn't set keep-alive: %v", err)
	}
}
