package dlms

import (
	"bytes"
	"testing"
)

func TestNew_ExceptionResponse(t *testing.T) {
	var a ExceptionResponse = *CreateExceptionResponse(1, 2)
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{216, 1, 2}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_ExceptionResponse(t *testing.T) {
	src := []byte{216, 1, 2}
	a, err := DecodeExceptionResponse(&src)
	if err != nil {
		t.Errorf("t1 failed on DecodeExceptionResponse. Err: %v", err)
	}

	var b ExceptionResponse = *CreateExceptionResponse(1, 2)

	if a.StateError != b.StateError {
		t.Errorf("t1 err StateError. get: %v, should: %v", a.StateError, b.StateError)
	}

	if a.ServiceError != b.ServiceError {
		t.Errorf("t1 err ServiceError. get: %v, should: %v", a.ServiceError, b.ServiceError)
	}

	src = []byte{216, 1}
	_, err = DecodeExceptionResponse(&src)
	if err == nil {
		t.Errorf("t1 should failed on DecodeExceptionResponse. Err: %v", err)
	}

	src = []byte{217, 1, 2}
	_, err = DecodeExceptionResponse(&src)
	if err == nil {
		t.Errorf("t1 should failed on DecodeExceptionResponse. Err: %v", err)
	}
}
