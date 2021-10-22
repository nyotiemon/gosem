package dlms

import (
	"bytes"
	"gosem/pkg/axdr"
	"reflect"
	"testing"
)

func TestNewGetResponseNormal(t *testing.T) {
	var dt axdr.DlmsData = *axdr.CreateAxdrDoubleLong(69)
	var ret GetDataResult = *CreateGetDataResultAsData(dt)

	a := *CreateGetResponseNormal(81, ret)
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{196, 1, 81, 1, 5, 0, 0, 0, 69}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNewGetResponseWithDataBlock(t *testing.T) {
	var dbg DataBlockG = *CreateDataBlockGAsData(true, 1, "07D20C04030A060BFF007800")

	a := *CreateGetResponseWithDataBlock(81, dbg)
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{196, 2, 81, 255, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// ---
	dbg = *CreateDataBlockGAsResult(true, 1, TagAccSuccess)
	b := *CreateGetResponseWithDataBlock(81, dbg)
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{196, 2, 81, 255, 0, 0, 0, 1, 1, 0}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestNewGetResponseWithList(t *testing.T) {
	var gdr1 GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)

	a := *CreateGetResponseWithList(69, []GetDataResult{gdr1})
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{196, 3, 69, 1, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt1 axdr.DlmsData = *axdr.CreateAxdrDoubleLong(1)
	var gdr2 GetDataResult = *CreateGetDataResultAsData(dt1)
	b := *CreateGetResponseWithList(69, []GetDataResult{gdr1, gdr2})
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
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

	val := a.Result.Value.(axdr.DlmsData)
	if val.Tag != axdr.TagDoubleLong {
		t.Errorf("t1 Failed. get: %d, should:%v", val.Tag, axdr.TagDoubleLong)
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
	val2 := a.Result.Result.(AccessResultTag)
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
	var dt1 axdr.DlmsData = *axdr.CreateAxdrDoubleLong(1)
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
	a1DescObis, _ := a.ResultList[0].ValueAsAccess()
	b1DescObis, _ := b.ResultList[0].ValueAsAccess()
	if a1DescObis != b1DescObis {
		t.Errorf("t1 Failed. ResultList[0].Value get: %v, should: %v", a1DescObis, b1DescObis)
	}
	if !a.ResultList[1].IsData {
		t.Errorf("t1 Failed. ResultList[0].IsData should be true, get: %v", a.ResultList[1].IsData)
	}
	a2DescObis, _ := a.ResultList[1].ValueAsData()
	b2DescObis, _ := b.ResultList[1].ValueAsData()
	if a2DescObis.Value != b2DescObis.Value {
		t.Errorf("t1 Failed. ResultList[1].Value get: %v, should: %v", a2DescObis.Value, b2DescObis.Value)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_GetResponse(t *testing.T) {
	var gr GetResponse

	srcNormal := []byte{196, 1, 81, 1, 5, 0, 0, 0, 69}
	a, e1 := gr.Decode(&srcNormal)
	if e1 != nil {
		t.Errorf("Decode for GetResponseNormal Failed. err:%v", e1)
	}
	x, assertGetResponseNormal := a.(GetResponseNormal)
	if !assertGetResponseNormal {
		t.Errorf("Decode supposed to return GetResponseNormal instead of %v", reflect.TypeOf(a).Name())
	}
	valX := x.Result.Value.(axdr.DlmsData)
	if valX.Tag != axdr.TagDoubleLong {
		t.Errorf("GetResponseNormal Value wrong. get: %d, should:%v", valX.Tag, axdr.TagDoubleLong)
	}

	srcWithDataBlock := []byte{196, 2, 81, 1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	b, e2 := gr.Decode(&srcWithDataBlock)
	if e2 != nil {
		t.Errorf("Decode for GetResponseWithDataBlock Failed. err:%v", e1)
	}
	y, assertGetResponseWithDataBlock := b.(GetResponseWithDataBlock)
	if !assertGetResponseWithDataBlock {
		t.Errorf("Decode supposed to return GetResponseWithDataBlock instead of %v", reflect.TypeOf(b).Name())
	}
	valY, _ := y.Result.ResultAsBytes()
	res := bytes.Compare(valY, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 {
		t.Errorf("GetResponseWithDataBlock Result wrong. get: %d, should: %v", valY, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	}

	srcWithList := []byte{196, 3, 69, 2, 0, 0, 1, 5, 0, 0, 0, 1}
	c, e3 := gr.Decode(&srcWithList)
	if e3 != nil {
		t.Errorf("Decode for GetResponseWithList Failed. err:%v", e1)
	}
	z, assertGetResponseWithList := c.(GetResponseWithList)
	if !assertGetResponseWithList {
		t.Errorf("Decode supposed to return GetResponseWithList instead of %v", reflect.TypeOf(c).Name())
	}
	if z.ResultCount != 2 {
		t.Errorf("GetResponseNormal ResultCount wrong. get: %d, should: 2", z.ResultCount)
	}
}
