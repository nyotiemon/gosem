package dlms

import (
	"encoding/binary"
	"fmt"
	"gosem/pkg/axdr"
)

type MethodDescriptor struct {
	ClassId    uint16
	InstanceId Obis
	MethodId   int8
}

func CreateMethodDescriptor(c uint16, i string, a int8) *MethodDescriptor {
	var ob Obis = *CreateObis(i)

	return &MethodDescriptor{ClassId: c, InstanceId: ob, MethodId: a}
}

func (ad MethodDescriptor) Encode() (out []byte, err error) {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.Bytes()...)
	output = append(output, byte(ad.MethodId))

	out = output
	return
}

func DecodeMethodDescriptor(ori *[]byte) (out MethodDescriptor, err error) {
	src := append([]byte(nil), (*ori)...)

	if len(src) < 9 {
		err = fmt.Errorf("byte slice length must be at least 9 bytes")
		return
	}

	_, out.ClassId, err = axdr.DecodeLongUnsigned(&src)
	if err != nil {
		return
	}
	out.InstanceId, err = DecodeObis(&src)
	if err != nil {
		return
	}
	out.MethodId = int8(src[0])
	src = src[1:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
