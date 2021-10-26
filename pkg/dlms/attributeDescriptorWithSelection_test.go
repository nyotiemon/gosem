package dlms

import (
	"bytes"
	"testing"
)

func TestAttributeDescriptorWithSelection(t *testing.T) {
	var nilSAD *SelectiveAccessDescriptor
	var a AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, nilSAD)
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("Test 1 with nil SelectiveAccessDescriptor failed. get: %d, should:%v", t1, result)
	}

	var sad SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var b AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, &sad)
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("Test 2 with SelectiveAccessDescriptor failed. get: %d, should:%v", t2, result)
	}
}

func TestDecode_AttributeDescriptorWithSelection(t *testing.T) {
	src := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 2, 3}
	a, e := DecodeAttributeDescriptorWithSelection(&src)
	if e != nil {
		t.Errorf("t1 failed with err: %v", e)
	}

	var sad SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var b AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, &sad)

	if a.ClassId != b.ClassId {
		t.Errorf("ClassId get: %v, should:%v", a.ClassId, b.ClassId)
	}
	res := bytes.Compare(a.InstanceId.Bytes(), b.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("InstanceId get: %v, should:%v", a.InstanceId.Bytes(), b.InstanceId.Bytes())
	}
	if a.AttributeId != b.AttributeId {
		t.Errorf("AttributeId get: %v, should:%v", a.AttributeId, b.AttributeId)
	}

	if a.AccessDescriptor.AccessSelector != b.AccessDescriptor.AccessSelector {
		t.Errorf("AccessDescriptor.AccessSelector get: %v, should:%v", a.AccessDescriptor.AccessSelector, b.AccessDescriptor.AccessSelector)
	}

	res = bytes.Compare(src, []byte{1, 2, 3})
	if res != 0 {
		t.Errorf("t1 reminder failed. get: %v, should: [1, 2, 3]", src)
	}
}
