package prodneutral

import "testing"

func TestSettingRetries(t *testing.T) {
	err := settingRetries(5)

	if err != nil {
		t.Errorf("retries: error: %v\n", err)
	}
}

func TestSettingBackoffPolicy(t *testing.T) {
	err := settingBackoffPolicy(3, 1, 5, 1.5, 0.3)
	if err != nil {
		t.Errorf("backoff: error: %v\n", err)
	}
}
