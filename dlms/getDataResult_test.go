package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

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
	var a GetDataResult = *CreateGetDataResult(false, rs)

	t1 := a.Encode()
	result := []byte{0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var b GetDataResult = *CreateGetDataResult(true, dt)
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
	var c GetDataResult = *CreateGetDataResult(false, 999)
	c.Encode()
}

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
