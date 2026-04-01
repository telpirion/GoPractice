package prodneutral

import "testing"

func TestSettingConnectionPool(t *testing.T) {
	err := setConnectionPool(3)
	if err != nil {
		t.Errorf("connection pool: error: %v\n", err)
	}
}
