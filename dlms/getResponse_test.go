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
	result := []byte{196, 2, 81, 1, 0, 0, 0, 1, 0, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNewGetResponseWithList(t *testing.T) {
	var gr GetResponse
	var gdr1 GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)

	a := gr.New(TagGetResponseWithList)
	a = *CreateGetResponseWithList(69, []GetDataResult{gdr1})
	t1 := a.Encode()
	result := []byte{196, 3, 69, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt1 DlmsData = *CreateAxdrDoubleLong(1)
	var gdr2 GetDataResult = *CreateGetDataResultAsData(dt1)
	b := *CreateGetResponseWithList(69, []GetDataResult{gdr1, gdr2})
	t2 := b.Encode()
	result = []byte{196, 3, 69, 0, 0, 1, 5, 0, 0, 0, 1}
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
