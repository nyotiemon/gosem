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
	result := []byte{1, 0, 0, 0, 1, 0, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockGAsData(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2 := b.Encode()
	result = []byte{1, 0, 0, 0, 1, 0, 1, 0, 0, 3, 0, 255}

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
	result := []byte{1, 0, 0, 0, 1, 0, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var b DataBlockG = *CreateDataBlockG(true, 1, []byte{1, 0, 0, 3, 0, 255})
	t2 := b.Encode()
	result = []byte{1, 0, 0, 0, 1, 0, 1, 0, 0, 3, 0, 255}

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
