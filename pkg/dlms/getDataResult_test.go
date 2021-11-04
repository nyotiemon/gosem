package dlms

import (
	"bytes"
	"gosem/pkg/axdr"
	"testing"
)

func TestAccessResult(t *testing.T) {
	t1 := TagAccSuccess
	if t1.String() != "success" {
		t.Errorf("t1 should return string with value 'success'")
	}
	t2 := TagAccObjectUnavailable
	if t2.String() != "object-unavailable" {
		t.Errorf("t1 should return string with value 'object-unavailable'")
	}
}

func TestGetDataResultAsResult(t *testing.T) {
	var a GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{0, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestGetDataResultAsData(t *testing.T) {
	var dt axdr.DlmsData = *axdr.CreateAxdrDoubleLong(69)
	var a GetDataResult = *CreateGetDataResultAsData(dt)

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{1, 5, 0, 0, 0, 69}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestGetDataResult(t *testing.T) {
	rs := TagAccSuccess
	var a GetDataResult = *CreateGetDataResult(rs)

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt axdr.DlmsData = *axdr.CreateAxdrDoubleLong(69)
	var b GetDataResult = *CreateGetDataResult(dt)
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
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

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{255, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockGAsData(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{255, 0, 0, 0, 1, 0, 6, 1, 0, 0, 3, 0, 255}

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

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{255, 0, 0, 0, 1, 1, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDataBlockG(t *testing.T) {
	var a DataBlockG = *CreateDataBlockG(true, 1, "07D20C04030A060BFF007800")

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{255, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockG(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{255, 0, 0, 0, 1, 0, 6, 1, 0, 0, 3, 0, 255}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	var c DataBlockG = *CreateDataBlockG(true, 1, TagAccSuccess)

	t3, e := c.Encode()
	if e != nil {
		t.Errorf("t3 Encode Failed. err: %v", e)
	}
	result = []byte{255, 0, 0, 0, 1, 1, 0}

	res = bytes.Compare(t3, result)
	if res != 0 {
		t.Errorf("t3 Failed. get: %d, should:%v", t3, result)
	}
}

func TestDataBlockSA(t *testing.T) {
	// with hexstring
	var a DataBlockSA = *CreateDataBlockSA(true, 1, "07D20C04030A060BFF007800")

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{255, 0, 0, 0, 1, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with byte slice
	var b DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{255, 0, 0, 0, 1, 6, 1, 0, 0, 3, 0, 255}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}

	// with wrong value
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t3 should've panic on wrong Value")
		}
	}()
	var c DataBlockSA = *CreateDataBlockSA(true, 1, TagAccSuccess)
	c.Encode()
}

func TestActResponse(t *testing.T) {
	var dt axdr.DlmsData = *axdr.CreateAxdrDoubleLong(69)
	var ret GetDataResult = *CreateGetDataResultAsData(dt)
	var a ActResponse = *CreateActResponse(TagActSuccess, &ret)

	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{0, 1, 1, 5, 0, 0, 0, 69}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with nil GetDataResult
	var nilRet *GetDataResult
	var b ActResponse = *CreateActResponse(TagActReadWriteDenied, nilRet)
	t2, e := b.Encode()
	if e != nil {
		t.Errorf("t2 Encode Failed. err: %v", e)
	}
	result = []byte{3, 0}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestDecode_GetDataResult(t *testing.T) {
	// with AccessResultTag
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

	// with dlms data
	src = []byte{1, 5, 0, 0, 0, 69}
	b, be := DecodeGetDataResult(&src)
	if be != nil {
		t.Errorf("t2 Failed. got error: %v", be)
	}
	if !b.IsData {
		t.Errorf("t2 Failed. Value should be data")
	}
	val := b.Value.(axdr.DlmsData)
	if val.Tag != axdr.TagDoubleLong {
		t.Errorf("t2 Failed. get: %d, should:%v", val.Tag, axdr.TagDoubleLong)
	}
	if v := val.Value.(int32); v != 69 {
		t.Errorf("t2 Failed. get: %d, should:%v", v, 69)
	}
}

func TestDecode_DataBlockG(t *testing.T) {
	// with byte slice
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
	val, _ := a.ResultAsBytes()
	res := bytes.Compare(val, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 {
		t.Errorf("t1 Failed. Result is not correct (%v)", val)
	}

	// with AccessResultTag
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
	_, eTag := b.ResultAsAccess()
	if eTag != nil {
		t.Errorf("t2 Failed. Result should be TagAccSuccess (%v)", eTag)
	}
}

func TestDecode_DataBlockSA(t *testing.T) {
	src := []byte{1, 0, 0, 0, 1, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	a, ae := DecodeDataBlockSA(&src)

	if ae != nil {
		t.Errorf("t1 Failed. got error: %v", ae)
	}
	if !a.LastBlock {
		t.Errorf("t1 Failed. LastBlock should be true")
	}
	if a.BlockNumber != 1 {
		t.Errorf("t1 Failed. BlockNumber should be 1 (%v)", a.BlockNumber)
	}
	res := bytes.Compare(a.Raw, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 {
		t.Errorf("t1 Failed. Result is not correct (%v)", a.Raw)
	}
}

func TestDecode_ActResponse(t *testing.T) {
	src := []byte{0, 1, 0, 0}
	a, ae := DecodeActResponse(&src)

	if ae != nil {
		t.Errorf("Error on DecodeActResponse: %v", ae)
	}
	if a.Result != TagActSuccess {
		t.Errorf("Result should be TagActSuccess")
	}
	if a.ReturnParam == nil {
		t.Errorf("ReturnParam should not be nil (%v)", a.ReturnParam)
	}
	if a.ReturnParam.IsData == true {
		t.Errorf("ReturnParam.IsData should not be true (%v)", a.ReturnParam.IsData)
	}
	tag, err := a.ReturnParam.ValueAsAccess()
	if tag != TagAccSuccess || err != nil {
		t.Errorf("ReturnParam.Value should be TagAccSuccess (%v, err: %v)", tag, err)
	}
}
