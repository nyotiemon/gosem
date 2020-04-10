package cosem

import (
	"fmt"
	"strconv"
	"strings"
)

// this object doesn't have header/length part, unlike DlmsData OctetString
type Obis struct {
	stringValue string
	byteValue   [6]byte
}

func (o *Obis) Set(str string) error {
	s := strings.Split(str, ".")
	if len(s) < 6 {
		return fmt.Errorf("Obis code is built of 6 uint8 connected with dots. sample: 1.0.0.3.0.255")
	}
	var bv [6]byte
	for i, v := range s {
		bt, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return err
		}
		bv[i] = uint8(bt)
	}

	o.stringValue = str
	o.byteValue = bv

	return nil
}

func CreateObis(str string) *Obis {
	o := &Obis{}
	o.Set(str)

	return o
}

func (o Obis) String() string {
	return o.stringValue
}

func (o Obis) Bytes() []byte {
	return o.byteValue[:]
}

func DecodeObis(src *[]byte) (outVal Obis, err error) {
	if len(*src) < 6 {
		err = fmt.Errorf("byte slice length must be at least 6 bytes")
		return
	}
	btVal := [6]byte{(*src)[0], (*src)[1], (*src)[2], (*src)[3], (*src)[4], (*src)[5]}
	strVal := fmt.Sprintf("%v.%v.%v.%v.%v.%v", (*src)[0], (*src)[1], (*src)[2], (*src)[3], (*src)[4], (*src)[5])
	outVal = Obis{stringValue: strVal, byteValue: btVal}
	(*src) = (*src)[6:]
	return
}
