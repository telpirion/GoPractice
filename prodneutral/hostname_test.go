package prodneutral

import "testing"

func TestSettingHostname(t *testing.T) {
	err := settingHostname("eu-vision.googleapis.com:443")
	if err != nil {
		t.Errorf("hostname: cannot set: %v\n", err)
	}
}
