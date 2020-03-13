package axdr

import (
	"encoding/binary"
	"fmt"
)

type Decoder struct {
	tag dataTag
}

func NewDecoder(dt dataTag) *Decoder {
	return &Decoder{tag: dt}
}

func (dec *Decoder) Decode(src []byte) (r DlmsData, err error) {
	haveLength, _ := lengthAfterTag[dec.tag]
	if haveLength {

	}

	return
}

func decodeLength(src *[]byte) (l []byte, v uint64, err error) {
	if (*src)[0] > byte(128) {
		varL := int((*src)[0]) - 128
		l = (*src)[0 : varL+1]
		rl := (*src)[1 : varL+1]

		buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}

		if len(rl) > 8 {
			err = fmt.Errorf("Length value is bigger than uint64 max value. This decoder is limited to uint64")
		} else {
			bufIndex := 7
			for i := len(rl) - 1; i > 0; i-- {
				fmt.Printf("set buf[%v], from %v to %v\n", bufIndex, buf[bufIndex], rl[i])
				buf[bufIndex] = rl[i]
				bufIndex -= 1
			}
		}
		fmt.Printf("buf now : %v\n", buf[:])
		v = binary.BigEndian.Uint64(buf[:])
		(*src) = (*src)[1+len(rl):]

	} else {
		l = append(l, (*src)[0])
		v = uint64((*src)[0])
		(*src) = (*src)[1:]
	}

	return
}
