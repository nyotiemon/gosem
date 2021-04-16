package dlms

import (
	"bytes"
	"testing"
)

func TestMethodDescriptor(t *testing.T) {
	var a MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_MethodDescriptor(t *testing.T) {
	src := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 3}
	a, err := DecodeMethodDescriptor(&src)
	if err != nil {
		t.Errorf("DecodeMethodDescriptor failed with err: %v", err)
	}

	var b MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)

	if a != b {
		t.Errorf("MethodDescriptor after decode is wrong. get: %v, should:%v", b, a)
	}

	res := bytes.Compare(src, []byte{1, 2, 3})
	if res != 0 {
		t.Errorf("byte reminder wrong. get: %v, should: [1, 2, 3]", src)
	}
}
