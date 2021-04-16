package dlms

import (
	"bytes"
	"testing"
)

func TestAttributeDescriptor_Encode(t *testing.T) {
	var a AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)

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

func TestAttributeDescriptor_Decode(t *testing.T) {
	var a AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	src := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 3}
	b, e := DecodeAttributeDescriptor(&src)

	if e != nil {
		t.Errorf("t1 failed with err: %v", e)
	}
	if a != b {
		t.Errorf("t1 failed. get: %v, should:%v", b, a)
	}

	res := bytes.Compare(src, []byte{1, 2, 3})
	if res != 0 {
		t.Errorf("t1 reminder failed. get: %v, should: [1, 2, 3]", src)
	}
}
