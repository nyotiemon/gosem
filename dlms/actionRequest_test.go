package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

func TestNew_ActionRequestNormal(t *testing.T) {
	var mthDesc MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DlmsData = *CreateAxdrOctetString("0102030405")
	var a ActionRequestNormal = *CreateActionRequestNormal(81, mthDesc, &dt)
	t1 := a.Encode()
	result := []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with nil Data
	var nilData *DlmsData = nil
	var b ActionRequestNormal = *CreateActionRequestNormal(81, mthDesc, nilData)
	t2 := b.Encode()
	result = []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

}

func TestNew_ActionRequestNextPBlock(t *testing.T) {
	var a ActionRequestNextPBlock = *CreateActionRequestNextPBlock(81, 1)
	t1 := a.Encode()
	result := []byte{195, 2, 81, 0, 0, 0, 1}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_ActionRequestWithList(t *testing.T) {
	// with 1 MethodDescriptor & 1 DlmsData
	var md1 MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt1 DlmsData = *CreateAxdrOctetString("0102030405")
	var a ActionRequestWithList = *CreateActionRequestWithList(81, []MethodDescriptor{md1}, []DlmsData{dt1})
	t1 := a.Encode()
	result := []byte{195, 3, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with 2 MethodDescriptor & 2 DlmsData
	var md2 MethodDescriptor = *CreateMethodDescriptor(1, "0.0.8.0.0.255", 2)
	var dt2 DlmsData = *CreateAxdrDoubleLong(69)
	var b ActionRequestWithList = *CreateActionRequestWithList(81, []MethodDescriptor{md1, md2}, []DlmsData{dt1, dt2})
	t2 := b.Encode()
	result = []byte{195, 3, 81, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 2, 9, 5, 1, 2, 3, 4, 5, 5, 0, 0, 0, 69}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestNew_ActionRequestWithFirstPBlock(t *testing.T) {
	var md MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var a ActionRequestWithFirstPBlock = *CreateActionRequestWithFirstPBlock(81, md, dt)
	t1 := a.Encode()
	result := []byte{195, 4, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_ActionRequestWithListAndFirstPBlock(t *testing.T) {
	// with 1 MethodDescriptor
	var a1 MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})

	var a ActionRequestWithListAndFirstPBlock = *CreateActionRequestWithListAndFirstPBlock(81, []MethodDescriptor{a1}, dt)
	t1 := a.Encode()
	result := []byte{195, 5, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with 2 MethodDescriptor
	var a2 MethodDescriptor = *CreateMethodDescriptor(1, "0.0.8.0.0.255", 2)
	var b ActionRequestWithListAndFirstPBlock = *CreateActionRequestWithListAndFirstPBlock(81, []MethodDescriptor{a1, a2}, dt)
	t2 := b.Encode()
	result = []byte{195, 5, 81, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	// with empty MethodDescriptor
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	c := *CreateActionRequestWithListAndFirstPBlock(69, []MethodDescriptor{}, dt)
	c.Encode()
}

func TestNew_ActionRequestWithPBlock(t *testing.T) {
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})

	var a ActionRequestWithPBlock = *CreateActionRequestWithPBlock(81, dt)
	t1 := a.Encode()
	result := []byte{195, 6, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_ActionRequestNormal(t *testing.T) {
	src := []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionRequestNormal(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestNormal. err:%v", err)
	}

	var mthDesc MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DlmsData = *CreateAxdrOctetString("0102030405")
	var b ActionRequestNormal = *CreateActionRequestNormal(81, mthDesc, &dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.MethodInfo.ClassId != b.MethodInfo.ClassId {
		t.Errorf("t1 Failed. MethodInfo.ClassId get: %v, should:%v", a.MethodInfo.ClassId, b.MethodInfo.ClassId)
	}
	res := bytes.Compare(a.MethodInfo.InstanceId.Bytes(), b.MethodInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. MethodInfo.InstanceId get: %v, should:%v", a.MethodInfo.InstanceId.Bytes(), b.MethodInfo.InstanceId.Bytes())
	}
	if a.MethodInfo.MethodId != b.MethodInfo.MethodId {
		t.Errorf("t1 Failed. MethodInfo.MethodId get: %v, should:%v", a.MethodInfo.MethodId, b.MethodInfo.MethodId)
	}

	if a.MethodParam.Tag != b.MethodParam.Tag {
		t.Errorf("t1 Failed. MethodParam.Tag get: %v, should:%v", a.MethodParam.Tag, b.MethodParam.Tag)
	}

	if a.MethodParam.Value != b.MethodParam.Value {
		t.Errorf("t1 Failed. MethodParam.Tag get: %v, should:%v", a.MethodParam.Value, b.MethodParam.Value)
	}

	// --- with nil data

	src = []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0}
	a, err = DecodeActionRequestNormal(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestNormal. err:%v", err)
	}

	if a.MethodParam != nil {
		t.Errorf("t2 Failed. MethodParam should be nil, get: %v", a.MethodParam)
	}

}

func TestDecode_ActionRequestNextPBlock(t *testing.T) {
	var x ActionRequestNextPBlock = *CreateActionRequestNextPBlock(81, 1)
	src := []byte{195, 2, 81, 0, 0, 0, 1}

	a, err := DecodeActionRequestNextPBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.BlockNum != x.BlockNum {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.BlockNum, x.BlockNum)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_ActionRequestWithList(t *testing.T) {
	// ---------------------- with 1 MethodDescriptor
	src := []byte{195, 3, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionRequestWithList(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestWithList. err:%v", err)
	}

	var a1 MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var d1 DlmsData = *CreateAxdrOctetString("0102030405")
	var b ActionRequestWithList = *CreateActionRequestWithList(81, []MethodDescriptor{a1}, []DlmsData{d1})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.MethodInfoList) != len(b.MethodInfoList) {
		t.Errorf("t1 Failed. MethodInfoList count get: %v, should:%v", len(a.MethodInfoList), len(b.MethodInfoList))
	}
	if a.MethodInfoCount != b.MethodInfoCount {
		t.Errorf("t1 Failed. MethodInfoCount get: %v, should:%v", a.MethodInfoCount, b.MethodInfoCount)
	}
	aDescObis := a.MethodInfoList[0].InstanceId.String()
	bDescObis := b.MethodInfoList[0].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. MethodInfoList[0].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}
	if len(a.MethodParamList) != len(b.MethodParamList) {
		t.Errorf("t1 Failed. MethodParamList count get: %v, should:%v", len(a.MethodParamList), len(b.MethodParamList))
	}
	if a.MethodParamCount != b.MethodParamCount {
		t.Errorf("t1 Failed. MethodParamCount get: %v, should:%v", a.MethodParamCount, b.MethodParamCount)
	}
	aDataTag := a.MethodParamList[0].Tag
	bDataTag := b.MethodParamList[0].Tag
	if aDataTag != bDataTag {
		t.Errorf("t1 Failed. MethodParamList[0].Tag get: %v, should:%v", aDataTag, bDataTag)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 MethodDescriptor
	src = []byte{195, 3, 81, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 2, 9, 5, 1, 2, 3, 4, 5, 5, 0, 0, 0, 69}
	a, err = DecodeActionRequestWithList(&src)
	if err != nil {
		t.Errorf("t2 Failed to DecodeActionRequestWithList. err:%v", err)
	}

	var a2 MethodDescriptor = *CreateMethodDescriptor(1, "0.0.8.0.0.255", 2)
	var d2 DlmsData = *CreateAxdrDoubleLong(69)
	b = *CreateActionRequestWithList(81, []MethodDescriptor{a1, a2}, []DlmsData{d1, d2})

	if len(a.MethodInfoList) != len(b.MethodInfoList) {
		t.Errorf("t2 Failed. MethodInfoList count get: %v, should:%v", len(a.MethodInfoList), len(b.MethodInfoList))
	}
	if a.MethodInfoCount != b.MethodInfoCount {
		t.Errorf("t2 Failed. MethodInfoCount get: %v, should:%v", a.MethodInfoCount, b.MethodInfoCount)
	}
	aDescObis = a.MethodInfoList[1].InstanceId.String()
	bDescObis = b.MethodInfoList[1].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t2 Failed. MethodInfoList[1].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}
	if len(a.MethodParamList) != len(b.MethodParamList) {
		t.Errorf("t2 Failed. MethodParamList count get: %v, should:%v", len(a.MethodParamList), len(b.MethodParamList))
	}
	if a.MethodParamCount != b.MethodParamCount {
		t.Errorf("t2 Failed. MethodParamCount get: %v, should:%v", a.MethodParamCount, b.MethodParamCount)
	}
	aDataTag = a.MethodParamList[1].Tag
	bDataTag = b.MethodParamList[1].Tag
	if aDataTag != bDataTag {
		t.Errorf("t2 Failed. MethodParamList[1].Tag get: %v, should:%v", aDataTag, bDataTag)
	}

	if len(src) > 0 {
		t.Errorf("t2 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_ActionRequestWithFirstPBlock(t *testing.T) {
	src := []byte{195, 4, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionRequestWithFirstPBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestWithFirstPBlock. err:%v", err)
	}

	var md MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b ActionRequestWithFirstPBlock = *CreateActionRequestWithFirstPBlock(81, md, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}

	if a.MethodInfo.ClassId != b.MethodInfo.ClassId {
		t.Errorf("t1 Failed. MethodInfo.ClassId get: %v, should:%v", a.MethodInfo.ClassId, b.MethodInfo.ClassId)
	}
	res := bytes.Compare(a.MethodInfo.InstanceId.Bytes(), b.MethodInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. MethodInfo.InstanceId get: %v, should:%v", a.MethodInfo.InstanceId.Bytes(), b.MethodInfo.InstanceId.Bytes())
	}
	if a.MethodInfo.MethodId != b.MethodInfo.MethodId {
		t.Errorf("t1 Failed. MethodInfo.MethodId get: %v, should:%v", a.MethodInfo.MethodId, b.MethodInfo.MethodId)
	}

	if a.PBlock.LastBlock != b.PBlock.LastBlock {
		t.Errorf("t1 Failed. PBlock.LastBlock get: %v, should:%v", a.PBlock.LastBlock, b.PBlock.LastBlock)
	}
	if a.PBlock.BlockNumber != b.PBlock.BlockNumber {
		t.Errorf("t1 Failed. PBlock.BlockNumber get: %v, should:%v", a.PBlock.BlockNumber, b.PBlock.BlockNumber)
	}
	res = bytes.Compare(a.PBlock.Raw, a.PBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. PBlock.Raw get: %v, should:%v", a.PBlock.Raw, a.PBlock.Raw)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_ActionRequestWithListAndFirstPBlock(t *testing.T) {
	src := []byte{195, 5, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionRequestWithListAndFirstPBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestWithListAndFirstPBlock. err:%v", err)
	}

	var a1 MethodDescriptor = *CreateMethodDescriptor(1, "1.0.0.3.0.255", 2)
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b ActionRequestWithListAndFirstPBlock = *CreateActionRequestWithListAndFirstPBlock(81, []MethodDescriptor{a1}, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}

	if len(a.MethodInfoList) != len(b.MethodInfoList) {
		t.Errorf("t1 Failed. MethodInfoList count get: %v, should:%v", len(a.MethodInfoList), len(b.MethodInfoList))
	}
	if a.MethodInfoCount != b.MethodInfoCount {
		t.Errorf("t1 Failed. MethodInfoCount get: %v, should:%v", a.MethodInfoCount, b.MethodInfoCount)
	}
	aDescObis := a.MethodInfoList[0].InstanceId.String()
	bDescObis := b.MethodInfoList[0].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t1 Failed. MethodInfoList[0].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}

	if a.PBlock.LastBlock != b.PBlock.LastBlock {
		t.Errorf("t1 Failed. PBlock.LastBlock get: %v, should:%v", a.PBlock.LastBlock, b.PBlock.LastBlock)
	}
	if a.PBlock.BlockNumber != b.PBlock.BlockNumber {
		t.Errorf("t1 Failed. PBlock.BlockNumber get: %v, should:%v", a.PBlock.BlockNumber, b.PBlock.BlockNumber)
	}
	res := bytes.Compare(a.PBlock.Raw, a.PBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. PBlock.Raw get: %v, should:%v", a.PBlock.Raw, a.PBlock.Raw)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 MethodDescriptor
	src = []byte{195, 5, 81, 2, 0, 1, 1, 0, 0, 3, 0, 255, 2, 0, 1, 0, 0, 8, 0, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err = DecodeActionRequestWithListAndFirstPBlock(&src)
	if err != nil {
		t.Errorf("t2 Failed to DecodeActionRequestWithList. err:%v", err)
	}

	var a2 MethodDescriptor = *CreateMethodDescriptor(1, "0.0.8.0.0.255", 2)
	b = *CreateActionRequestWithListAndFirstPBlock(81, []MethodDescriptor{a1, a2}, dt)

	if len(a.MethodInfoList) != len(b.MethodInfoList) {
		t.Errorf("t2 Failed. MethodInfoList count get: %v, should:%v", len(a.MethodInfoList), len(b.MethodInfoList))
	}
	if a.MethodInfoCount != b.MethodInfoCount {
		t.Errorf("t2 Failed. MethodInfoCount get: %v, should:%v", a.MethodInfoCount, b.MethodInfoCount)
	}
	aDescObis = a.MethodInfoList[1].InstanceId.String()
	bDescObis = b.MethodInfoList[1].InstanceId.String()
	if aDescObis != bDescObis {
		t.Errorf("t2 Failed. MethodInfoList[1].InstanceId get: %v, should:%v", aDescObis, bDescObis)
	}

	if len(src) > 0 {
		t.Errorf("t2 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_ActionRequestWithPBlock(t *testing.T) {
	src := []byte{195, 6, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionRequestWithPBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionRequestWithPBlock. err:%v", err)
	}

	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b ActionRequestWithPBlock = *CreateActionRequestWithPBlock(81, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}

	if a.PBlock.LastBlock != b.PBlock.LastBlock {
		t.Errorf("t1 Failed. PBlock.LastBlock get: %v, should:%v", a.PBlock.LastBlock, b.PBlock.LastBlock)
	}
	if a.PBlock.BlockNumber != b.PBlock.BlockNumber {
		t.Errorf("t1 Failed. PBlock.BlockNumber get: %v, should:%v", a.PBlock.BlockNumber, b.PBlock.BlockNumber)
	}
	res := bytes.Compare(a.PBlock.Raw, a.PBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. PBlock.Raw get: %v, should:%v", a.PBlock.Raw, a.PBlock.Raw)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_ActionRequest(t *testing.T) {
	var sr ActionRequest

	// ------------------  ActionRequestNormal
	src := []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	_, e := sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestNormal Failed. err:%v", e)
	}

	// ------------------  ActionRequestNextPBlock
	src = []byte{195, 2, 81, 0, 0, 0, 1}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestNextPBlock Failed. err:%v", e)
	}

	// ------------------  ActionRequestWithList
	src = []byte{195, 3, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithList Failed. err:%v", e)
	}

	// ------------------  ActionRequestWithFirstPBlock
	src = []byte{195, 4, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithFirstPBlock Failed. err:%v", e)
	}

	// ------------------  ActionRequestWithListAndFirstPBlock
	src = []byte{195, 5, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithListAndFirstPBlock Failed. err:%v", e)
	}

	// ------------------  ActionRequestWithPBlock
	src = []byte{195, 6, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithPBlock Failed. err:%v", e)
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := sr.Decode(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
