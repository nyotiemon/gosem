package cosem

import (
	"bytes"
	"testing"
)

func TestObis(t *testing.T) {
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
}
