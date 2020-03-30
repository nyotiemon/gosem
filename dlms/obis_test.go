package cosem

import (
	"bytes"
	"testing"
)

func TestEncodeObis(t *testing.T) {
	i := Obis{stringValue: "1.0.0.3.0.255", byteValue: [6]byte{1, 0, 0, 3, 0, 255}}
	var o Obis = *CreateObis("1.0.0.3.0.255")

	res := bytes.Compare(i.Bytes(), o.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed to convert string obis to byte. get: %d, should:%v", o.Bytes(), i.Bytes())
	}

	var u Obis
	u.Set("1.0.0.3.0.255")

	res = bytes.Compare(i.Bytes(), u.Bytes())
	if res != 0 {
		t.Errorf("t2 Failed to convert string obis to byte. get: %d, should:%v", u.Bytes(), i.Bytes())
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should've panic on wrong Value")
		}
	}()
	var w Obis
	w.Set("hahaha")
}

func TestDecodeObis(t *testing.T) {
	src := []byte{1, 0, 0, 3, 0, 255, 1, 2, 3}
	ob, e := DecodeObis(&src)
	var o Obis = *CreateObis("1.0.0.3.0.255")

	if e != nil {
		t.Errorf("t1 failed with err: %v", e)
	}
	if o != ob {
		t.Errorf("t1 failed. get: %d, should:%v", ob.Bytes(), o.Bytes())
	}
}
