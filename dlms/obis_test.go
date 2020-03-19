package cosem

import (
	"bytes"
	"testing"
)

func TestObis(t *testing.T) {
	i := Obis{StringValue: "1.0.0.3.0.255", ByteValue: [6]byte{1, 0, 0, 3, 0, 255}}

	var o Obis = *CreateObis("1.0.0.3.0.255")

	res := bytes.Compare(i.ByteValue[:], o.ByteValue[:])
	if res != 0 {
		t.Errorf("Failed to convert string obis to byte. get: %d, should:%v", o.ByteValue, i.ByteValue)
	}
}
