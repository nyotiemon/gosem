package cosem

import (
	"bytes"
	"testing"
)

func TestNewGetRequestNormal(t *testing.T) {
	var gr GetRequest
	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})

	a := gr.New(TagGetRequestNormal)
	a = *CreateGetRequestNormal(81, attrDesc, &accsDesc)
	t1 := a.Encode()
	result := []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	b := *CreateGetRequestNormal(81, attrDesc, nilAccsDesc)
	t2 := b.Encode()
	result = []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}
}

func TestNewGetRequestNext(t *testing.T) {
	var gr GetRequest

	a := gr.New(TagGetRequestNext)
	a = *CreateGetRequestNext(81, 2)
	t1 := a.Encode()
	result := []byte{192, 2, 81, 0, 0, 0, 2}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNewGetRequestWithList(t *testing.T) {
	var gr GetRequest
	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)

	a := gr.New(TagGetRequestWithList)
	a = *CreateGetRequestWithList(69, []AttributeDescriptor{a1})
	t1 := a.Encode()
	result := []byte{192, 3, 69, 0, 1, 1, 0, 0, 3, 0, 255, 2}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	b := *CreateGetRequestWithList(69, []AttributeDescriptor{a1, a2})
	t2 := b.Encode()
	result = []byte{192, 3, 69, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	c := *CreateGetRequestWithList(69, []AttributeDescriptor{})
	c.Encode()
}
