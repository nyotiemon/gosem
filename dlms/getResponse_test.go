package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

func TestNewGetResponseNormal(t *testing.T) {
	var gr GetResponse
	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var ret GetDataResult = *CreateGetDataResultAsData(dt)

	a := gr.New(TagGetResponseNormal)
	a = *CreateGetResponseNormal(81, ret)
	t1 := a.Encode()
	result := []byte{196, 1, 81, 1, 5, 0, 0, 0, 69}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNewGetResponseWithDataBlock(t *testing.T) {
	var gr GetResponse
	var dbg DataBlockG = *CreateDataBlockGAsData(true, 1, "07D20C04030A060BFF007800")

	a := gr.New(TagGetResponseWithDataBlock)
	a = *CreateGetResponseWithDataBlock(81, dbg)
	t1 := a.Encode()
	result := []byte{196, 2, 81, 1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// ---
	b := gr.New(TagGetResponseWithDataBlock)
	dbg = *CreateDataBlockGAsResult(true, 1, TagAccSuccess)
	b = *CreateGetResponseWithDataBlock(81, dbg)
	t2 := b.Encode()
	result = []byte{196, 2, 81, 1, 0, 0, 0, 1, 1, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestNewGetResponseWithList(t *testing.T) {
	var gr GetResponse
	var gdr1 GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)

	a := gr.New(TagGetResponseWithList)
	a = *CreateGetResponseWithList(69, []GetDataResult{gdr1})
	t1 := a.Encode()
	result := []byte{196, 3, 69, 1, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt1 DlmsData = *CreateAxdrDoubleLong(1)
	var gdr2 GetDataResult = *CreateGetDataResultAsData(dt1)
	b := *CreateGetResponseWithList(69, []GetDataResult{gdr1, gdr2})
	t2 := b.Encode()
	result = []byte{196, 3, 69, 2, 0, 0, 1, 5, 0, 0, 0, 1}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on empty GetDataResult list")
		}
	}()
	c := *CreateGetResponseWithList(69, []GetDataResult{})
	c.Encode()
}

func TestDecode_GetResponseNormal(t *testing.T) {
	src := []byte{196, 1, 81, 1, 5, 0, 0, 0, 69}
	a, err := DecodeGetResponseNormal(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetResponseNormal. err:%v", err)
	}
	if a.InvokePriority != 81 {
		t.Errorf("t1 Failed. InvokePriority should 81, get: %v", a.InvokePriority)
	}
	if a.Result.IsData != true {
		t.Errorf("t1 Failed. Result.IsData should true, get: %v", a.Result.IsData)
	}

	val := a.Result.Value.(DlmsData)
	if val.Tag != TagDoubleLong {
		t.Errorf("t1 Failed. get: %d, should:%v", val.Tag, TagDoubleLong)
	}
	if v := val.Value.(int32); v != 69 {
		t.Errorf("t1 Failed. get: %d, should:%v", v, 69)
	}
}

func TestDecode_GetResponseWithDataBlock(t *testing.T) {
	src := []byte{196, 2, 81, 1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	a, err := DecodeGetResponseWithDataBlock(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetResponseWithDataBlock. err:%v", err)
	}
	if a.InvokePriority != 81 {
		t.Errorf("t1 Failed. InvokePriority should be 81, get: %v", a.InvokePriority)
	}
	val := a.Result.Result.([]byte)
	res := bytes.Compare(val, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 {
		t.Errorf("t1 Failed. Result incorrect, get: %v", val)
	}

	// ---
	src = []byte{196, 2, 81, 1, 0, 0, 0, 1, 1, 0}
	a, err = DecodeGetResponseWithDataBlock(&src)
	if err != nil {
		t.Errorf("t2 Failed to DecodeGetResponseWithDataBlock. err:%v", err)
	}
	if a.InvokePriority != 81 {
		t.Errorf("t2 Failed. InvokePriority should be 81, get: %v", a.InvokePriority)
	}
	val2 := a.Result.Result.(accessResultTag)
	if val2 != TagAccSuccess {
		t.Errorf("t2 Failed. Result should be TagAccSuccess, get: %v", val2)
	}
}

func TestDecode_GetResponseWithList(t *testing.T) {
	src := []byte{196, 3, 69, 2, 0, 0, 1, 5, 0, 0, 0, 1}
	a, err := DecodeGetResponseWithList(&src)

	if err != nil {
		t.Errorf("t1 Failed to DecodeGetResponseWithList. err:%v", err)
	}

	var gdr1 GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)
	var dt1 DlmsData = *CreateAxdrDoubleLong(1)
	var gdr2 GetDataResult = *CreateGetDataResultAsData(dt1)

	var b GetResponseWithList = *CreateGetResponseWithList(69, []GetDataResult{gdr1, gdr2})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should: %v", a.InvokePriority, b.InvokePriority)
	}
	if len(a.ResultList) != len(b.ResultList) {
		t.Errorf("t1 Failed. ResultList count get: %v, should: %v", len(a.ResultList), len(b.ResultList))
	}
	if a.ResultCount != b.ResultCount {
		t.Errorf("t1 Failed. ResultCount get: %v, should: %v", a.ResultCount, b.ResultCount)
	}
	if a.ResultList[0].IsData {
		t.Errorf("t1 Failed. ResultList[0].IsData should be false, get: %v", a.ResultList[0].IsData)
	}
	a1DescObis := a.ResultList[0].ValueAsAccess()
	b1DescObis := b.ResultList[0].ValueAsAccess()
	if a1DescObis != b1DescObis {
		t.Errorf("t1 Failed. ResultList[0].Value get: %v, should: %v", a1DescObis, b1DescObis)
	}
	if !a.ResultList[1].IsData {
		t.Errorf("t1 Failed. ResultList[0].IsData should be true, get: %v", a.ResultList[1].IsData)
	}
	a2DescObis := uint32(a.ResultList[1].ValueAsData().Value.(int32))
	b2DescObis := uint32(b.ResultList[1].ValueAsData().Value.(int))
	if a2DescObis != b2DescObis {
		t.Errorf("t1 Failed. ResultList[1].Value get: %v, should: %v", a2DescObis, b2DescObis)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}
