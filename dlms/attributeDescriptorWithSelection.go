package cosem

import "encoding/binary"

type AttributeDescriptorWithSelection struct {
	ClassId          uint16
	InstanceId       Obis
	AttributeId      int8
	AccessDescriptor SelectiveAccessDescriptor
}

func CreateAttributeDescriptorWithSelection(c uint16, i string, a int8, sad SelectiveAccessDescriptor) *AttributeDescriptorWithSelection {
	var ob Obis = *CreateObis(i)

	return &AttributeDescriptorWithSelection{ClassId: c, InstanceId: ob, AttributeId: a, AccessDescriptor: sad}
}

func (ad *AttributeDescriptorWithSelection) Encode() []byte {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.ByteValue[:]...)
	output = append(output, byte(ad.AttributeId))
	output = append(output, ad.AccessDescriptor.Encode()[:]...)

	return output
}
