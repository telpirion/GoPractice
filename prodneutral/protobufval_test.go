package prodneutral

import "testing"

func TestCreatingProtobufValue(t *testing.T) {
	err := creatingProtobufValue()
	if err != nil {
		t.Errorf("protobuf value: error: %v\n", err)
	}
}
