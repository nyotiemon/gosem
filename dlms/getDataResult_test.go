package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

func TestAccessResult(t *testing.T) {
	t1 := TagAccSuccess
	if "success" != t1.String() {
		t.Errorf("t1 should return string with value 'success'")
	}
	t2 := TagAccObjectUnavailable
	if "object-unavailable" != t2.String() {
		t.Errorf("t1 should return string with value 'success'")
	}
}

func TestGetDataResultAsResult(t *testing.T) {
	var a GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)

	t1 := a.Encode()
	result := []byte{0, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestGetDataResultAsData(t *testing.T) {
	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var a GetDataResult = *CreateGetDataResultAsData(dt)

	t1 := a.Encode()
	result := []byte{1, 5, 0, 0, 0, 69}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestGetDataResult(t *testing.T) {
	rs := TagAccSuccess
	var a GetDataResult = *CreateGetDataResult(rs)

	t1 := a.Encode()
	result := []byte{0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var b GetDataResult = *CreateGetDataResult(dt)
	t2 := b.Encode()
	result = []byte{1, 5, 0, 0, 0, 69}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	var c GetDataResult = *CreateGetDataResult(999)
	c.Encode()
}

func TestDataBlockGAsData(t *testing.T) {
	var a DataBlockG = *CreateDataBlockGAsData(true, 1, "07D20C04030A060BFF007800")

	t1 := a.Encode()
	result := []byte{1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockGAsData(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2 := b.Encode()
	result = []byte{1, 0, 0, 0, 1, 0, 6, 1, 0, 0, 3, 0, 255}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	var c DataBlockG = *CreateDataBlockGAsData(true, 1, TagAccSuccess)
	c.Encode()
}

func TestDataBlockGAsResult(t *testing.T) {
	var a DataBlockG = *CreateDataBlockGAsResult(true, 1, TagAccSuccess)

	t1 := a.Encode()
	result := []byte{1, 0, 0, 0, 1, 1, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDataBlockG(t *testing.T) {
	var a DataBlockG = *CreateDataBlockG(true, 1, "07D20C04030A060BFF007800")

	t1 := a.Encode()
	result := []byte{1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockG(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2 := b.Encode()
	result = []byte{1, 0, 0, 0, 1, 0, 6, 1, 0, 0, 3, 0, 255}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	var c DataBlockG = *CreateDataBlockG(true, 1, TagAccSuccess)

	t3 := c.Encode()
	result = []byte{1, 0, 0, 0, 1, 1, 0}

	res = bytes.Compare(t3, result)
	if res != 0 {
		t.Errorf("t3 Failed. get: %d, should:%v", t3, result)
	}
}

func TestDataBlockSA(t *testing.T) {
	var a DataBlockSA = *CreateDataBlockSA(true, 1, "07D20C04030A060BFF007800")

	t1 := a.Encode()
	result := []byte{1, 0, 0, 0, 1, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2 := b.Encode()
	result = []byte{1, 0, 0, 0, 1, 1, 0, 0, 3, 0, 255}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	var c DataBlockSA = *CreateDataBlockSA(true, 1, TagAccSuccess)
	c.Encode()
}

func TestActionResponseWithOptData(t *testing.T) {
	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var ret GetDataResult = *CreateGetDataResultAsData(dt)
	var a ActionResponseWithOptData = *CreateActionResponseWithOptData(TagActSuccess, &ret)

	t1 := a.Encode()
	result := []byte{0, 1, 1, 5, 0, 0, 0, 69}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var nilRet *GetDataResult = nil
	var b ActionResponseWithOptData = *CreateActionResponseWithOptData(TagActReadWriteDenied, nilRet)
	t2 := b.Encode()
	result = []byte{3, 0}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestDecode_GetDataResult(t *testing.T) {
	src := []byte{0, 0}
	a, ae := DecodeGetDataResult(&src)

	if ae != nil {
		t.Errorf("t1 Failed. got error: %v", ae)
	}
	if a.IsData {
		t.Errorf("t1 Failed. Value should be access")
	}
	if a.Value != TagAccSuccess {
		t.Errorf("t1 Failed. get: %d, should:%v", a.Value, TagAccSuccess)
	}

	src = []byte{1, 5, 0, 0, 0, 69}
	b, be := DecodeGetDataResult(&src)
	if be != nil {
		t.Errorf("t2 Failed. got error: %v", be)
	}
	if !b.IsData {
		t.Errorf("t2 Failed. Value should be data")
	}
	val := b.Value.(DlmsData)
	if val.Tag != TagDoubleLong {
		t.Errorf("t2 Failed. get: %d, should:%v", val.Tag, TagDoubleLong)
	}
	if v := val.Value.(int32); v != 69 {
		t.Errorf("t2 Failed. get: %d, should:%v", v, 69)
	}
}

func TestDecode_DataBlockG(t *testing.T) {
	src := []byte{1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	a, ae := DecodeDataBlockG(&src)

	if ae != nil {
		t.Errorf("t1 Failed. got error: %v", ae)
	}
	if !a.LastBlock {
		t.Errorf("t1 Failed. LastBlock should be true")
	}
	if a.BlockNumber != 1 {
		t.Errorf("t1 Failed. BlockNumber should be 1 (%v)", a.BlockNumber)
	}
	if a.IsResult {
		t.Errorf("t1 Failed. IsResult should be false")
	}
	val := a.ResultAsBytes()
	res := bytes.Compare(val, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 {
		t.Errorf("t1 Failed. Result is not correct (%v)", val)
	}

	// ---
	src = []byte{1, 0, 0, 0, 1, 1, 0}
	b, be := DecodeDataBlockG(&src)

	if be != nil {
		t.Errorf("t2 Failed. got error: %v", be)
	}
	if !b.LastBlock {
		t.Errorf("t2 Failed. LastBlock should be true")
	}
	if b.BlockNumber != 1 {
		t.Errorf("t2 Failed. BlockNumber should be 1 (%v)", b.BlockNumber)
	}
	if !b.IsResult {
		t.Errorf("t2 Failed. IsResult should be true")
	}
	if b.ResultAsAccess() != TagAccSuccess {
		t.Errorf("t2 Failed. Result should be TagAccSuccess (%v)", b.Result)
	}

}
