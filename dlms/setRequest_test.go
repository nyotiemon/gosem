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
	result := []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	var b SetRequestWithFirstDataBlock = *CreateSetRequestWithFirstDataBlock(81, attrDesc, nilAccsDesc, dt)
	t2 := b.Encode()
	result = []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

}

func TestNew_SetRequestWithDataBlock(t *testing.T) {
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var a SetRequestWithDataBlock = *CreateSetRequestWithDataBlock(81, dt)
	t1 := a.Encode()
	result := []byte{193, 3, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
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
	result := []byte{193, 5, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	var b SetRequestWithListAndFirstDataBlock = *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{a1, a2}, dt)
	t2 := b.Encode()
	result = []byte{193, 5, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
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

func TestDecode_SetRequestNormal(t *testing.T) {
	src := []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 9, 5, 1, 2, 3, 4, 5}
	a, err := DecodeSetRequestNormal(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
	}

	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var dt DlmsData = *CreateAxdrOctetString("0102030405")
	var b SetRequestNormal = *CreateSetRequestNormal(81, attrDesc, &accsDesc, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.AttributeInfo.ClassId != b.AttributeInfo.ClassId {
		t.Errorf("t1 Failed. AttributeInfo.ClassId get: %v, should:%v", a.AttributeInfo.ClassId, b.AttributeInfo.ClassId)
	}
	res := bytes.Compare(a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. AttributeInfo.InstanceId get: %v, should:%v", a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	}
	if a.AttributeInfo.AttributeId != b.AttributeInfo.AttributeId {
		t.Errorf("t1 Failed. AttributeInfo.AttributeId get: %v, should:%v", a.AttributeInfo.AttributeId, b.AttributeInfo.AttributeId)
	}
	if a.SelectiveAccessInfo.AccessSelector != b.SelectiveAccessInfo.AccessSelector {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector get: %v, should:%v", a.SelectiveAccessInfo.AccessSelector, b.SelectiveAccessInfo.AccessSelector)
	}
	aByte := a.SelectiveAccessInfo.AccessParameter.Encode()
	bByte := b.SelectiveAccessInfo.AccessParameter.Encode()
	res = bytes.Compare(aByte, bByte)
	if res != 0 {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessParameter get: %v, should:%v", aByte, bByte)
	}
	if a.Value.Tag != b.Value.Tag {
		t.Errorf("t1 Failed. Value.Tag get: %v, should:%v", a.Value.Tag, b.Value.Tag)
	}
	if a.Value.Value != b.Value.Value {
		t.Errorf("t1 Failed. Value.Value get: %v, should:%v", a.Value.Value, b.Value.Value)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ------------------ t2 without SelectiveAccessDescriptor

	src = []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 9, 5, 1, 2, 3, 4, 5}
	a, err = DecodeSetRequestNormal(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestNormal. err:%v", err)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	b = *CreateSetRequestNormal(81, attrDesc, nilAccsDesc, dt)

	if a.SelectiveAccessInfo != nilAccsDesc {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector should be nil get: %v", a.SelectiveAccessInfo)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetRequestWithFirstDataBlock(t *testing.T) {
	src := []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeSetRequestWithFirstDataBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestWithFirstDataBlock. err:%v", err)
	}

	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b SetRequestWithFirstDataBlock = *CreateSetRequestWithFirstDataBlock(81, attrDesc, &accsDesc, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.AttributeInfo.ClassId != b.AttributeInfo.ClassId {
		t.Errorf("t1 Failed. AttributeInfo.ClassId get: %v, should:%v", a.AttributeInfo.ClassId, b.AttributeInfo.ClassId)
	}
	res := bytes.Compare(a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. AttributeInfo.InstanceId get: %v, should:%v", a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	}
	if a.AttributeInfo.AttributeId != b.AttributeInfo.AttributeId {
		t.Errorf("t1 Failed. AttributeInfo.AttributeId get: %v, should:%v", a.AttributeInfo.AttributeId, b.AttributeInfo.AttributeId)
	}
	if a.SelectiveAccessInfo.AccessSelector != b.SelectiveAccessInfo.AccessSelector {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector get: %v, should:%v", a.SelectiveAccessInfo.AccessSelector, b.SelectiveAccessInfo.AccessSelector)
	}
	aByte := a.SelectiveAccessInfo.AccessParameter.Encode()
	bByte := b.SelectiveAccessInfo.AccessParameter.Encode()
	res = bytes.Compare(aByte, bByte)
	if res != 0 {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessParameter get: %v, should:%v", aByte, bByte)
	}
	if a.DataBlock.LastBlock != b.DataBlock.LastBlock {
		t.Errorf("t1 Failed. DataBlock.LastBlock get: %v, should:%v", a.DataBlock.LastBlock, b.DataBlock.LastBlock)
	}
	if a.DataBlock.BlockNumber != b.DataBlock.BlockNumber {
		t.Errorf("t1 Failed. DataBlock.BlockNumber get: %v, should:%v", a.DataBlock.BlockNumber, b.DataBlock.BlockNumber)
	}
	res = bytes.Compare(a.DataBlock.Raw, a.DataBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. DataBlock.Raw get: %v, should:%v", a.DataBlock.Raw, a.DataBlock.Raw)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ------------------ t2 without SelectiveAccessDescriptor

	src = []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err = DecodeSetRequestWithFirstDataBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestWithFirstDataBlock. err:%v", err)
	}

	var nilAccsDesc *SelectiveAccessDescriptor = nil
	b = *CreateSetRequestWithFirstDataBlock(81, attrDesc, nilAccsDesc, dt)

	if a.SelectiveAccessInfo != nilAccsDesc {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector should be nil get: %v", a.SelectiveAccessInfo)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetRequestWithDataBlock(t *testing.T) {
	src := []byte{193, 3, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeSetRequestWithDataBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestWithDataBlock. err:%v", err)
	}

	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b SetRequestWithDataBlock = *CreateSetRequestWithDataBlock(81, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.DataBlock.LastBlock != b.DataBlock.LastBlock {
		t.Errorf("t1 Failed. DataBlock.LastBlock get: %v, should:%v", a.DataBlock.LastBlock, b.DataBlock.LastBlock)
	}
	if a.DataBlock.BlockNumber != b.DataBlock.BlockNumber {
		t.Errorf("t1 Failed. DataBlock.BlockNumber get: %v, should:%v", a.DataBlock.BlockNumber, b.DataBlock.BlockNumber)
	}
	res := bytes.Compare(a.DataBlock.Raw, a.DataBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. DataBlock.Raw get: %v, should:%v", a.DataBlock.Raw, a.DataBlock.Raw)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_SetRequestWithList(t *testing.T) {
	src := []byte{193, 4, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	a, err := DecodeSetRequestWithList(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestWithList. err:%v", err)
	}

	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var d1 DlmsData = *CreateAxdrOctetString("0102030405")
	var b SetRequestWithList = *CreateSetRequestWithList(69, []AttributeDescriptor{a1}, []DlmsData{d1})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
	}
	if a.AttributeCount != b.AttributeCount {
		t.Errorf("t1 Failed. AttributeCount get: %v, should:%v", a.AttributeCount, b.AttributeCount)
	}
	aDescObis := a.AttributeInfoList[0].InstanceId.String()
	bDescObis := b.AttributeInfoList[0].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. AttributeInfoList[0].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}
	if len(a.ValueList) != len(b.ValueList) {
		t.Errorf("t1 Failed. ValueList count get: %v, should:%v", len(a.ValueList), len(b.ValueList))
	}
	if a.ValueCount != b.ValueCount {
		t.Errorf("t1 Failed. ValueCount get: %v, should:%v", a.ValueCount, b.ValueCount)
	}
	aDataTag := a.ValueList[0].Tag
	bDataTag := b.ValueList[0].Tag
	if aDataTag != bDataTag {
		t.Errorf("t1 Failed. ValueList[0].Tag get: %v, should:%v", aDataTag, bDataTag)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 AttributeDescriptor
	src = []byte{193, 4, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 2, 9, 5, 1, 2, 3, 4, 5, 5, 0, 0, 0, 69}
	a, err = DecodeSetRequestWithList(&src)

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	var d2 DlmsData = *CreateAxdrDoubleLong(69)
	b = *CreateSetRequestWithList(69, []AttributeDescriptor{a1, a2}, []DlmsData{d1, d2})

	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
	}
	if a.AttributeCount != b.AttributeCount {
		t.Errorf("t1 Failed. AttributeCount get: %v, should:%v", a.AttributeCount, b.AttributeCount)
	}
	aDescObis = a.AttributeInfoList[1].InstanceId.String()
	bDescObis = b.AttributeInfoList[1].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. AttributeInfoList[1].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}
	if len(a.ValueList) != len(b.ValueList) {
		t.Errorf("t1 Failed. ValueList count get: %v, should:%v", len(a.ValueList), len(b.ValueList))
	}
	if a.ValueCount != b.ValueCount {
		t.Errorf("t1 Failed. ValueCount get: %v, should:%v", a.ValueCount, b.ValueCount)
	}
	aDataTag = a.ValueList[1].Tag
	bDataTag = b.ValueList[1].Tag
	if aDataTag != bDataTag {
		t.Errorf("t1 Failed. ValueList[1].Tag get: %v, should:%v", aDataTag, bDataTag)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_SetRequestWithListAndFirstDataBlock(t *testing.T) {
	src := []byte{193, 5, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeSetRequestWithListAndFirstDataBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestWithListAndFirstDataBlock. err:%v", err)
	}

	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b SetRequestWithListAndFirstDataBlock = *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{a1}, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
	}
	if a.AttributeCount != b.AttributeCount {
		t.Errorf("t1 Failed. AttributeCount get: %v, should:%v", a.AttributeCount, b.AttributeCount)
	}
	aDescObis := a.AttributeInfoList[0].InstanceId.String()
	bDescObis := b.AttributeInfoList[0].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. AttributeInfoList[0].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}
	if a.DataBlock.LastBlock != b.DataBlock.LastBlock {
		t.Errorf("t1 Failed. DataBlock.LastBlock get: %v, should:%v", a.DataBlock.LastBlock, b.DataBlock.LastBlock)
	}
	if a.DataBlock.BlockNumber != b.DataBlock.BlockNumber {
		t.Errorf("t1 Failed. DataBlock.BlockNumber get: %v, should:%v", a.DataBlock.BlockNumber, b.DataBlock.BlockNumber)
	}
	res := bytes.Compare(a.DataBlock.Raw, a.DataBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. DataBlock.Raw get: %v, should:%v", a.DataBlock.Raw, a.DataBlock.Raw)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 AttributeDescriptor
	src = []byte{193, 5, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err = DecodeSetRequestWithListAndFirstDataBlock(&src)

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	b = *CreateSetRequestWithListAndFirstDataBlock(69, []AttributeDescriptor{a1, a2}, dt)

	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
	}
	if a.AttributeCount != b.AttributeCount {
		t.Errorf("t1 Failed. AttributeCount get: %v, should:%v", a.AttributeCount, b.AttributeCount)
	}
	aDescObis = a.AttributeInfoList[1].InstanceId.String()
	bDescObis = b.AttributeInfoList[1].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. AttributeInfoList[1].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_SetRequest(t *testing.T) {
	var sr SetRequest

	// ------------------  SetRequestNormal
	src := []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 9, 5, 1, 2, 3, 4, 5}
	_, e := sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestNormal Failed. err:%v", e)
	}

	// ------------------  SetRequestWithFirstDataBlock
	src = []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestWithFirstDataBlock Failed. err:%v", e)
	}

	// ------------------  SetRequestWithDataBlock
	src = []byte{193, 3, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestWithDataBlock Failed. err:%v", e)
	}

	// ------------------  SetRequestWithList
	src = []byte{193, 4, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestWithList Failed. err:%v", e)
	}

	// ------------------  SetRequestWithDataBlock
	src = []byte{193, 5, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestWithListAndFirstDataBlock Failed. err:%v", e)
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := sr.Decode(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
