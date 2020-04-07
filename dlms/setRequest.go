package cosem

import (
	"bytes"
	. "gosem/axdr"
)

type setRequestTag uint8

const (
	TagSetRequestNormal                    setRequestTag = 0x1
	TagSetRequestWithFirstDataBlock        setRequestTag = 0x2
	TagSetRequestWithDataBlock             setRequestTag = 0x3
	TagSetRequestWithList                  setRequestTag = 0x4
	TagSetRequestWithListAndFirstDataBlock setRequestTag = 0x5
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s setRequestTag) Value() uint8 {
	return uint8(s)
}

// SetRequest implement CosemI
type SetRequest struct{}

func (gr *SetRequest) New(tag setRequestTag) CosemPDU {
	switch tag {
	case TagSetRequestNormal:
		return &SetRequestNormal{}
	case TagSetRequestWithFirstDataBlock:
		return &SetRequestWithFirstDataBlock{}
	case TagSetRequestWithDataBlock:
		return &SetRequestWithDataBlock{}
	case TagSetRequestWithList:
		return &SetRequestWithList{}
	case TagSetRequestWithListAndFirstDataBlock:
		return &SetRequestWithListAndFirstDataBlock{}
	default:
		panic("Tag not recognized!")
	}
}

// TODO
func (gr *SetRequest) Decode(src *[]byte) (out CosemPDU, err error) {
	return
}

// SetRequestNormal implement CosemPDU
type SetRequestNormal struct {
	InvokePriority      uint8
	AttributeInfo       AttributeDescriptor
	SelectiveAccessInfo *SelectiveAccessDescriptor
	Value               DlmsData
}

func CreateSetRequestNormal(invokeId uint8, att AttributeDescriptor, acc *SelectiveAccessDescriptor, dt DlmsData) *SetRequestNormal {
	return &SetRequestNormal{
		InvokePriority:      invokeId,
		AttributeInfo:       att,
		SelectiveAccessInfo: acc,
		Value:               dt,
	}
}

func (sr SetRequestNormal) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestNormal))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.Write(sr.AttributeInfo.Encode())
	if sr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		buf.Write(sr.SelectiveAccessInfo.Encode())
	}
	buf.Write(sr.Value.Encode())

	return buf.Bytes()
}

// SetRequestWithFirstDataBlock implement CosemPDU
type SetRequestWithFirstDataBlock struct {
	InvokePriority      uint8
	AttributeInfo       AttributeDescriptor
	SelectiveAccessInfo *SelectiveAccessDescriptor
	DataBlock           DataBlockSA
}

func CreateSetRequestWithFirstDataBlock(invokeId uint8, att AttributeDescriptor, acc *SelectiveAccessDescriptor, dt DataBlockSA) *SetRequestWithFirstDataBlock {
	return &SetRequestWithFirstDataBlock{
		InvokePriority:      invokeId,
		AttributeInfo:       att,
		SelectiveAccessInfo: acc,
		DataBlock:           dt,
	}
}

func (sr SetRequestWithFirstDataBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithFirstDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.Write(sr.AttributeInfo.Encode())
	if sr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		buf.Write(sr.SelectiveAccessInfo.Encode())
	}
	buf.Write(sr.DataBlock.Encode())

	return buf.Bytes()
}

// SetRequestWithDataBlock implement CosemPDU
type SetRequestWithDataBlock struct {
	InvokePriority uint8
	DataBlock      DataBlockSA
}

func CreateSetRequestWithDataBlock(invokeId uint8, dt DataBlockSA) *SetRequestWithDataBlock {
	return &SetRequestWithDataBlock{
		InvokePriority: invokeId,
		DataBlock:      dt,
	}
}

func (sr SetRequestWithDataBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.Write(sr.DataBlock.Encode())

	return buf.Bytes()
}

// SetRequestWithList implement CosemPDU
type SetRequestWithList struct {
	InvokePriority    uint8
	AttributeCount    uint8
	AttributeInfoList []AttributeDescriptor
	ValueCount        uint8
	ValueList         []DlmsData
}

func CreateSetRequestWithList(invokeId uint8, attList []AttributeDescriptor, valList []DlmsData) *SetRequestWithList {
	if len(attList) < 1 || len(attList) > 255 {
		panic("AttributeInfoList cannot have zero or >255 member")
	}
	if len(valList) < 1 || len(valList) > 255 {
		panic("ValueList cannot have zero or >255 member")
	}
	return &SetRequestWithList{
		InvokePriority:    invokeId,
		AttributeCount:    uint8(len(attList)),
		AttributeInfoList: attList,
		ValueCount:        uint8(len(valList)),
		ValueList:         valList,
	}
}

func (sr SetRequestWithList) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithList))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.WriteByte(byte(sr.AttributeCount))
	for _, attr := range sr.AttributeInfoList {
		buf.Write(attr.Encode())
	}
	buf.WriteByte(byte(sr.ValueCount))
	for _, val := range sr.ValueList {
		buf.Write(val.Encode())
	}

	return buf.Bytes()
}

// SetRequestWithListAndFirstDataBlock implement CosemPDU
type SetRequestWithListAndFirstDataBlock struct {
	InvokePriority    uint8
	AttributeCount    uint8
	AttributeInfoList []AttributeDescriptor
	DataBlock         DataBlockSA
}

func CreateSetRequestWithListAndFirstDataBlock(invokeId uint8, attList []AttributeDescriptor, dt DataBlockSA) *SetRequestWithListAndFirstDataBlock {
	if len(attList) < 1 || len(attList) > 255 {
		panic("AttributeInfoList cannot have zero or >255 member")
	}
	return &SetRequestWithListAndFirstDataBlock{
		InvokePriority:    invokeId,
		AttributeCount:    uint8(len(attList)),
		AttributeInfoList: attList,
		DataBlock:         dt,
	}
}

func (sr SetRequestWithListAndFirstDataBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithListAndFirstDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.WriteByte(byte(sr.AttributeCount))
	for _, attr := range sr.AttributeInfoList {
		buf.Write(attr.Encode())
	}
	buf.Write(sr.DataBlock.Encode())

	return buf.Bytes()
}
