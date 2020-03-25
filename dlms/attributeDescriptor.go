package cosem

import "encoding/binary"

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
