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
	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)

	a := gr.New(TagGetRequestWithList)
	a = *CreateGetRequestWithList(69, []AttributeDescriptor{a1})
	t1 := a.Encode()
	result := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	b := *CreateGetRequestWithList(69, []AttributeDescriptor{a1, a2})
	t2 := b.Encode()
	result = []byte{192, 3, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2}
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
	src := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2}
	a, err := DecodeGetRequestWithList(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetRequestWithList. err:%v", err)
	}

	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var b GetRequestWithList = *CreateGetRequestWithList(69, []AttributeDescriptor{a1})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
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
	src = []byte{192, 3, 69, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2}
	a, err = DecodeGetRequestWithList(&src)

	var a2 AttributeDescriptor = *CreateAttributeDescriptor(1, "0.0.8.0.0.255", 2)
	b = *CreateGetRequestWithList(69, []AttributeDescriptor{a1, a2})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.AttributeInfoList) != len(b.AttributeInfoList) {
		t.Errorf("t1 Failed. AttributeInfoList count get: %v, should:%v", len(a.AttributeInfoList), len(b.AttributeInfoList))
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
	x, assertGetRequestNormal := a.(GetRequestNormal)
	if !assertGetRequestNormal {
		t.Errorf("Decode supposed to return %v instead of %v", reflect.TypeOf(GetRequestNormal{}).Name(), reflect.TypeOf(a).Name())
	}

	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var accsDesc SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	var t1 GetRequestNormal = *CreateGetRequestNormal(81, attrDesc, &accsDesc)

	if x.InvokePriority != t1.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", x.InvokePriority, t1.InvokePriority)
	}
	if x.AttributeInfo.ClassId != t1.AttributeInfo.ClassId {
		t.Errorf("t1 Failed. AttributeInfo.ClassId get: %v, should:%v", x.AttributeInfo.ClassId, t1.AttributeInfo.ClassId)
	}
	res := bytes.Compare(x.AttributeInfo.InstanceId.Bytes(), t1.AttributeInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. AttributeInfo.InstanceId get: %v, should:%v", x.AttributeInfo.InstanceId.Bytes(), t1.AttributeInfo.InstanceId.Bytes())
	}
	if x.AttributeInfo.AttributeId != t1.AttributeInfo.AttributeId {
		t.Errorf("t1 Failed. AttributeInfo.AttributeId get: %v, should:%v", x.AttributeInfo.AttributeId, t1.AttributeInfo.AttributeId)
	}
	if x.SelectiveAccessInfo.AccessSelector != t1.SelectiveAccessInfo.AccessSelector {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessSelector get: %v, should:%v", x.SelectiveAccessInfo.AccessSelector, t1.SelectiveAccessInfo.AccessSelector)
	}
	xByte := x.SelectiveAccessInfo.AccessParameter.Encode()
	t1Byte := t1.SelectiveAccessInfo.AccessParameter.Encode()
	res = bytes.Compare(xByte, t1Byte)
	if res != 0 {
		t.Errorf("t1 Failed. SelectiveAccessInfo.AccessParameter get: %v, should:%v", xByte, t1Byte)
	}
	if len(srcNormal) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", srcNormal)
	}

	// ------------------  GetRequestNext
	srcNext := []byte{192, 2, 81, 0, 0, 0, 2}
	b, e2 := gr.Decode(&srcNext)
	if e2 != nil {
		t.Errorf("Decode for GetRequestNext Failed. err:%v", e2)
	}
	y, assertGetRequestNext := b.(GetRequestNext)
	if !assertGetRequestNext {
		t.Errorf("Decode supposed to return GetRequestNext instead of %v", reflect.TypeOf(b).Name())
	}

	var t2 GetRequestNext = *CreateGetRequestNext(81, 2)

	if y.InvokePriority != t2.InvokePriority {
		t.Errorf("t2 Failed. InvokePriority get: %v, should:%v", y.InvokePriority, t2.InvokePriority)
	}
	if y.BlockNum != t2.BlockNum {
		t.Errorf("t2 Failed. BlockNum get: %v, should:%v", y.BlockNum, t2.BlockNum)
	}
	if len(srcNext) > 0 {
		t.Errorf("t2 Failed. src should be empty. get: %v", srcNext)
	}

	// ------------------  GetRequestWithList
	srcWithList := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2}
	c, e3 := gr.Decode(&srcWithList)
	if e3 != nil {
		t.Errorf("Decode for GetRequestWithList Failed. err:%v", e3)
	}
	z, assertGetRequestWithList := c.(GetRequestWithList)
	if !assertGetRequestWithList {
		t.Errorf("Decode supposed to return GetRequestWithList instead of %v", reflect.TypeOf(c).Name())
	}

	var a1 AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var t3 GetRequestWithList = *CreateGetRequestWithList(69, []AttributeDescriptor{a1})

	if z.InvokePriority != t3.InvokePriority {
		t.Errorf("t3 Failed. InvokePriority get: %v, should:%v", z.InvokePriority, t3.InvokePriority)
	}
	if len(z.AttributeInfoList) != len(t3.AttributeInfoList) {
		t.Errorf("t3 Failed. AttributeInfoList count get: %v, should:%v", len(z.AttributeInfoList), len(t3.AttributeInfoList))
	}
	zDescObis := z.AttributeInfoList[0].InstanceId.String()
	t3DescObis := t3.AttributeInfoList[0].InstanceId.String()
	if zDescObis != t3DescObis {
		t.Errorf("t3 Failed. AttributeInfoList[0].InstanceId get: %v, should:%v", zDescObis, t3DescObis)
	}
	if len(srcWithList) > 0 {
		t.Errorf("t3 Failed. src should be empty. get: %v", srcWithList)
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := gr.Decode(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
