package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

func TestNew_SetResponseNormal(t *testing.T) {
	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var dt DlmsData = *CreateAxdrOctetString("0102030405")

	var a SetRequestNormal = *CreateSetRequestNormal(81, attrDesc, &accsDesc, dt)
	t1 := a.Encode()
	result := []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 9, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	var b SetRequestNormal = *CreateSetRequestNormal(81, attrDesc, nilAccsDesc, dt)
	t2 := b.Encode()
	result = []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 9, 5, 1, 2, 3, 4, 5}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

}

func TestNew_SetRequestWithFirstDataBlock(t *testing.T) {
	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})

	var a SetRequestWithFirstDataBlock = *CreateSetRequestWithFirstDataBlock(81, attrDesc, &accsDesc, dt)
	t1 := a.Encode()
	result := []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	var b SetRequestWithFirstDataBlock = *CreateSetRequestWithFirstDataBlock(81, attrDesc, nilAccsDesc, dt)
	t2 := b.Encode()
	result = []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 0, 1, 1, 2, 3, 4, 5}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

}

func TestNew_SetRequestWithDataBlock(t *testing.T) {
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var a SetRequestWithDataBlock = *CreateSetRequestWithDataBlock(81, dt)
	t1 := a.Encode()
	result := []byte{193, 3, 81, 1, 0, 0, 0, 1, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_SetRequestWithList(t *testing.T) {
	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var d1 DlmsData = *CreateAxdrOctetString("0102030405")

	var a SetRequestWithList = *CreateSetRequestWithList(69, []AttributeDescriptor{a1}, []DlmsData{d1})
	t1 := a.Encode()
	result := []byte{193, 4, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	var d2 DlmsData = *CreateAxdrDoubleLong(69)
	var b SetRequestWithList = *CreateSetRequestWithList(69, []AttributeDescriptor{a1, a2}, []DlmsData{d1, d2})
	t2 := b.Encode()
	result = []byte{193, 4, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 2, 9, 5, 1, 2, 3, 4, 5, 5, 0, 0, 0, 69}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	c := *CreateSetRequestWithList(69, []AttributeDescriptor{}, []DlmsData{d1, d2})
	c.Encode()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t4 should've panic on wrong Value")
		}
	}()
	d := *CreateSetRequestWithList(69, []AttributeDescriptor{a1, a2}, []DlmsData{})
	d.Encode()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t5 should've panic on wrong Value")
		}
	}()
	e := *CreateSetRequestWithList(69, []AttributeDescriptor{}, []DlmsData{})
	e.Encode()
}

func TestNew_SetRequestWithListAndFirstDataBlock(t *testing.T) {
	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})

	var a SetRequestWithListAndFirstDataBlock = *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{a1}, dt)
	t1 := a.Encode()
	result := []byte{193, 5, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	var b SetRequestWithListAndFirstDataBlock = *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{a1, a2}, dt)
	t2 := b.Encode()
	result = []byte{193, 5, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 0, 0, 0, 1, 1, 2, 3, 4, 5}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	c := *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{}, dt)
	c.Encode()
}
