package cosem

import (
	"bytes"
	"testing"
)

func TestMethodDescriptor(t *testing.T) {
	var a MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)

	t1 := a.Encode()
	result := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}
