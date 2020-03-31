package cosem

import (
	"bytes"
	"testing"
)

func TestAttributeDescriptorWithSelection(t *testing.T) {
	var nilSAD *SelectiveAccessDescriptor = nil
	var a AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, nilSAD)
	t1 := a.Encode()
	result := []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("Test 1 with nil SelectiveAccessDescriptor failed. get: %d, should:%v", t1, result)
	}

	var sad SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var b AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, &sad)
	t2 := b.Encode()
	result = []byte{0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("Test 2 with SelectiveAccessDescriptor failed. get: %d, should:%v", t2, result)
	}
}
