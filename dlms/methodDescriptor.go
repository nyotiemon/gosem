package cosem

import (
	"encoding/binary"
	"fmt"
	. "gosem/axdr"
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

func (ad *MethodDescriptor) Encode() []byte {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.Bytes()...)
	output = append(output, byte(ad.MethodId))

	return output
}

func DecodeMethodDescriptor(src *[]byte) (out MethodDescriptor, err error) {
	if len(*src) < 9 {
		err = fmt.Errorf("byte slice length must be at least 9 bytes")
		return
	}

	_, classId, eClass := DecodeLongUnsigned(src)
	if eClass != nil {
		err = eClass
		return
	}
	obis, errObis := DecodeObis(src)
	if errObis != nil {
		err = errObis
		return
	}
	attribId := int8((*src)[0])
	(*src) = (*src)[1:]

	out = MethodDescriptor{classId, obis, attribId}

	return
}
