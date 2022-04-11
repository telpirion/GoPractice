package kata2

import (
	"testing"
	"time"
)

func TestOpenReadUpdateJSON(t *testing.T) {
	fname := "../testdata/kata2_input.json"
	tNow := time.Now()
	tStr := tNow.Format(time.Stamp)
	k := "test"
	err := openReadUpdateJSON(fname, k, tStr)
	if err != nil {
		t.Error(err)
	}
}
