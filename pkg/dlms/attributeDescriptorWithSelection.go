package dlms

import (
	"encoding/binary"
	"fmt"
	"gosem/pkg/axdr"
)

type AttributeDescriptorWithSelection struct {
	ClassId          uint16
	InstanceId       Obis
	AttributeId      int8
	AccessDescriptor *SelectiveAccessDescriptor
}

// CreateAttributeDescriptorWithSelection will create AttributeDescriptorWithSelection object
// SelectiveAccessDescriptor is allowed to be nil value therefore pointer
func CreateAttributeDescriptorWithSelection(c uint16, i string, a int8, sad *SelectiveAccessDescriptor) *AttributeDescriptorWithSelection {
	var ob Obis = *CreateObis(i)

	return &AttributeDescriptorWithSelection{ClassId: c, InstanceId: ob, AttributeId: a, AccessDescriptor: sad}
}

func (ad AttributeDescriptorWithSelection) Encode() (out []byte, err error) {
	var output []byte
	var c [2]byte
	binary.BigEndian.PutUint16(c[:], ad.ClassId)
	output = append(output, c[:]...)
	output = append(output, ad.InstanceId.Bytes()...)
	output = append(output, byte(ad.AttributeId))
	if ad.AccessDescriptor == nil {
		output = append(output, 0)
	} else {
		output = append(output, 1)
		val, e := ad.AccessDescriptor.Encode()
		if e != nil {
			err = e
			return
		}
		output = append(output, val...)
	}

	out = output
	return
}

func DecodeAttributeDescriptorWithSelection(ori *[]byte) (out AttributeDescriptorWithSelection, err error) {
	src := append([]byte(nil), (*ori)...)

	if len(src) < 11 {
		err = fmt.Errorf("byte slice length must be at least 11 bytes")
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

	out.AttributeId = int8(src[0])
	haveAccDesc := src[1]
	src = src[2:]

	if haveAccDesc == 0x0 {
		var nilAccDesc *SelectiveAccessDescriptor
		out.AccessDescriptor = nilAccDesc
	} else {
		accDesc, errAcc := DecodeSelectiveAccessDescriptor(&src)
		if errAcc != nil {
			err = errAcc
			return
		}
		out.AccessDescriptor = &accDesc
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
