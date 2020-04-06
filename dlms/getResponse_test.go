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
