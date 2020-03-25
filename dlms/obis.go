package cosem

import (
	"strconv"
	"strings"
)

// this object doesn't have header/length part, unlike DlmsData OctetString
type Obis struct {
	stringValue string
	byteValue   [6]byte
}

func (o *Obis) Set(str string) {
	s := strings.Split(str, ".")
	if len(s) < 6 {
		panic("Obis code is built of 6 uint8 connected with dots. sample: 1.0.0.3.0.255")
	}
	var bv [6]byte
	for i, v := range s {
		bt, ok := strconv.ParseUint(v, 10, 8)
		if ok != nil {
			panic(ok)
		}
		bv[i] = uint8(bt)
	}

	o.stringValue = str
	o.byteValue = bv
}

func CreateObis(str string) *Obis {
	o := &Obis{}
	o.Set(str)

	return o
}

func (o *Obis) String() string {
	return o.stringValue
}

func (o *Obis) Bytes() []byte {
	return o.byteValue[:]
}
