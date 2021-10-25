package dlms

import (
	"bytes"
	"testing"
)

func TestNew_ConfirmedServiceError(t *testing.T) {
	var a ConfirmedServiceError = *CreateConfirmedServiceError(TagErrInitiateError, TagErrInitiate, 1)
	t1, e := a.Encode()
	if e != nil {
		t.Errorf("t1 Encode Failed. err: %v", e)
	}
	result := []byte{14, 1, 6, 1}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_ConfirmedServiceError(t *testing.T) {
	src := []byte{14, 1, 6, 1}
	a, err := DecodeConfirmedServiceError(&src)
	if err != nil {
		t.Errorf("t1 failed on DecodeConfirmedServiceError. Err: %v", err)
	}

	var b ConfirmedServiceError = *CreateConfirmedServiceError(TagErrInitiateError, TagErrInitiate, 1)

	if a.ConfirmedServiceError != b.ConfirmedServiceError {
		t.Errorf("t1 err ConfirmedServiceError. get: %v, should: %v", a.ConfirmedServiceError, b.ConfirmedServiceError)
	}

	if a.ServiceError != b.ServiceError {
		t.Errorf("t1 err ServiceError. get: %v, should: %v", a.ServiceError, b.ServiceError)
	}

	if a.Value != b.Value {
		t.Errorf("t1 err Value. get: %v, should: %v", a.Value, b.Value)
	}

	src = []byte{14, 1, 6}
	_, err = DecodeConfirmedServiceError(&src)
	if err == nil {
		t.Errorf("t1 should failed on DecodeConfirmedServiceError. Err: %v", err)
	}

	src = []byte{15, 1, 6, 1}
	_, err = DecodeConfirmedServiceError(&src)
	if err == nil {
		t.Errorf("t1 should failed on DecodeConfirmedServiceError. Err: %v", err)
	}
}
