package cosem

import (
	"encoding/binary"
	"fmt"
	. "gosem/axdr"
)

type AttributeDescriptor struct {
	ClassId     uint16
	InstanceId  Obis
	AttributeId int8
}

func CreateAttributeDescriptor(c uint16, i string, a int8) *AttributeDescriptor {
	var ob Obis = *CreateObis(i)

	return &AttributeDescriptor{ClassId: c, InstanceId: ob, AttributeId: a}
}

func (ad *AttributeDescriptor) Encode() []byte {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.Bytes()...)
	output = append(output, byte(ad.AttributeId))

	return output
}

func DecodeAttributeDescriptor(src *[]byte) (out AttributeDescriptor, err error) {
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

	out = AttributeDescriptor{classId, obis, attribId}

	return
}
