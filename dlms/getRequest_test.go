package cosem

import (
	"bytes"
	"reflect"
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
	var sad SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var a1 AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, &sad)

	a := gr.New(TagGetRequestWithList)
	a = *CreateGetRequestWithList(69, []AttributeDescriptorWithSelection{a1})
	t1 := a.Encode()
	result := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "0.0.8.0.0.255", 2, &sad)
	b := *CreateGetRequestWithList(69, []AttributeDescriptorWithSelection{a1, a2})
	t2 := b.Encode()
	result = []byte{192, 3, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	c := *CreateGetRequestWithList(69, []AttributeDescriptorWithSelection{})
	c.Encode()
}

func TestDecode_GetRequestNormal(t *testing.T) {
	src := []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	a, err := DecodeGetRequestNormal(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestNormal. err:%v", err)
	}

	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var b GetRequestNormal = *CreateGetRequestNormal(81, attrDesc, &accsDesc)

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
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ------------------ t2 without SelectiveAccessDescriptor

	src = []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0}
	a, err = DecodeGetRequestNormal(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestNormal. err:%v", err)
	}

	attrDesc = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var nilAccsDesc *SelectiveAccessDescriptor = nil
	b = *CreateGetRequestNormal(81, attrDesc, nilAccsDesc)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.AttributeInfo.ClassId != b.AttributeInfo.ClassId {
		t.Errorf("t1 Failed. AttributeInfo.ClassId get: %v, should:%v", a.AttributeInfo.ClassId, b.AttributeInfo.ClassId)
	}
	res = bytes.Compare(a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. AttributeInfo.InstanceId get: %v, should:%v", a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	}
	if a.AttributeInfo.AttributeId != b.AttributeInfo.AttributeId {
		t.Errorf("t1 Failed. AttributeInfo.AttributeId get: %v, should:%v", a.AttributeInfo.AttributeId, b.AttributeInfo.AttributeId)
	}
	if a.SelectiveAccessInfo != nilAccsDesc {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector should be nil get: %v", a.SelectiveAccessInfo)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_GetRequestNext(t *testing.T) {
	src := []byte{192, 2, 81, 0, 0, 0, 2}
	a, err := DecodeGetRequestNext(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestNext. err:%v", err)
	}

	var b GetRequestNext = *CreateGetRequestNext(81, 2)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.BlockNum != b.BlockNum {
		t.Errorf("t1 Failed. BlockNum get: %v, should:%v", a.BlockNum, b.BlockNum)
	}
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_GetRequestWithList(t *testing.T) {
	src := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	a, err := DecodeGetRequestWithList(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestWithList. err:%v", err)
	}

	var sad SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var a1 AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "1.0.0.3.0.255", 2, &sad)
	var b GetRequestWithList = *CreateGetRequestWithList(69, []AttributeDescriptorWithSelection{a1})

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
	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 AttributeDescriptor
	src = []byte{192, 3, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	a, err = DecodeGetRequestWithList(&src)

	var a2 AttributeDescriptorWithSelection = *CreateAttributeDescriptorWithSelection(1, "0.0.8.0.0.255", 2, &sad)
	b = *CreateGetRequestWithList(69, []AttributeDescriptorWithSelection{a1, a2})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
	}
	if a.AttributeCount != b.AttributeCount {
		t.Errorf("t1 Failed. AttributeCount get: %v, should:%v", a.AttributeCount, b.AttributeCount)
	}
	aDescObis = a.AttributeInfoList[0].InstanceId.String()
	bDescObis = b.AttributeInfoList[0].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. AttributeInfoList[0].InstanceId get: %v, should:%v", aDescObis, bDescObis)
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

func TestDecode_GetRequest(t *testing.T) {
	var gr GetRequest

	// ------------------  GetRequestNormal
	srcNormal := []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	a, e1 := gr.Decode(&srcNormal)
	if e1 != nil {
		t.Errorf("Decode for GetRequestNormal Failed. err:%v", e1)
	}
	_, assertGetRequestNormal := a.(GetRequestNormal)
	if !assertGetRequestNormal {
		t.Errorf("Decode supposed to return %v instead of %v", reflect.TypeOf(GetRequestNormal{}).Name(), reflect.TypeOf(a).Name())
	}

	// ------------------  GetRequestNext
	srcNext := []byte{192, 2, 81, 0, 0, 0, 2}
	b, e2 := gr.Decode(&srcNext)
	if e2 != nil {
		t.Errorf("Decode for GetRequestNext Failed. err:%v", e2)
	}
	_, assertGetRequestNext := b.(GetRequestNext)
	if !assertGetRequestNext {
		t.Errorf("Decode supposed to return GetRequestNext instead of %v", reflect.TypeOf(b).Name())
	}

	// ------------------  GetRequestWithList
	srcWithList := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	c, e3 := gr.Decode(&srcWithList)
	if e3 != nil {
		t.Errorf("Decode for GetRequestWithList Failed. err:%v", e3)
	}
	_, assertGetRequestWithList := c.(GetRequestWithList)
	if !assertGetRequestWithList {
		t.Errorf("Decode supposed to return GetRequestWithList instead of %v", reflect.TypeOf(c).Name())
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := gr.Decode(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
