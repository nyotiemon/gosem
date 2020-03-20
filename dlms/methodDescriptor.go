package cosem

import "encoding/binary"

type MethodDescriptor struct {
	ClassId     uint16
	InstanceId  Obis
	AttributeId int8
}

func CreateMethodDescriptor(c uint16, i string, a int8) *MethodDescriptor {
	var ob Obis = *CreateObis(i)

	return &MethodDescriptor{ClassId: c, InstanceId: ob, AttributeId: a}
}

func (ad *MethodDescriptor) Encode() []byte {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.ByteValue[:]...)
	output = append(output, byte(ad.AttributeId))

	return output
}