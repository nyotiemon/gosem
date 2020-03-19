package cosem

import (
	"strconv"
	"strings"
)

type Obis struct {
	StringValue string
	ByteValue   [6]byte
}

func CreateObis(str string) *Obis {
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

	return &Obis{StringValue: str, ByteValue: bv}
}
