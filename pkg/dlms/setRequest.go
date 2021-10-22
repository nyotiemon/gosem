package dlms

import (
	"bytes"
	"fmt"
	"gosem/pkg/axdr"
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

func (gr *SetRequest) New(tag setRequestTag) (out CosemPDU, err error) {
	switch tag {
	case TagSetRequestNormal:
		out = &SetRequestNormal{}
	case TagSetRequestWithFirstDataBlock:
		out = &SetRequestWithFirstDataBlock{}
	case TagSetRequestWithDataBlock:
		out = &SetRequestWithDataBlock{}
	case TagSetRequestWithList:
		out = &SetRequestWithList{}
	case TagSetRequestWithListAndFirstDataBlock:
		out = &SetRequestWithListAndFirstDataBlock{}
	default:
		err = fmt.Errorf("tag not recognized")
	}
	return
}

func (gr *SetRequest) Decode(src *[]byte) (out CosemPDU, err error) {
	if (*src)[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagSetRequest))
		return
	}

	switch (*src)[1] {
	case TagSetRequestNormal.Value():
		out, err = DecodeSetRequestNormal(src)
	case TagSetRequestWithFirstDataBlock.Value():
		out, err = DecodeSetRequestWithFirstDataBlock(src)
	case TagSetRequestWithDataBlock.Value():
		out, err = DecodeSetRequestWithDataBlock(src)
	case TagSetRequestWithList.Value():
		out, err = DecodeSetRequestWithList(src)
	case TagSetRequestWithListAndFirstDataBlock.Value():
		out, err = DecodeSetRequestWithListAndFirstDataBlock(src)
	default:
		err = fmt.Errorf("byte tag not recognized (%v)", (*src)[1])
	}

	return
}

// SetRequestNormal implement CosemPDU
type SetRequestNormal struct {
	InvokePriority      uint8
	AttributeInfo       AttributeDescriptor
	SelectiveAccessInfo *SelectiveAccessDescriptor
	Value               axdr.DlmsData
}

func CreateSetRequestNormal(invokeId uint8, att AttributeDescriptor, acc *SelectiveAccessDescriptor, dt axdr.DlmsData) *SetRequestNormal {
	return &SetRequestNormal{
		InvokePriority:      invokeId,
		AttributeInfo:       att,
		SelectiveAccessInfo: acc,
		Value:               dt,
	}
}

func (sr SetRequestNormal) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestNormal))
	buf.WriteByte(byte(sr.InvokePriority))
	attInfo, e := sr.AttributeInfo.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(attInfo)
	if sr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		selInfo, e := sr.SelectiveAccessInfo.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(selInfo)
	}
	val, eVal := sr.Value.Encode()
	if eVal != nil {
		err = eVal
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeSetRequestNormal(ori *[]byte) (out SetRequestNormal, err error) {
	var src = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagSetRequest))
		return
	}
	if src[1] != TagSetRequestNormal.Value() {
		err = ErrWrongTag(1, src[1], byte(TagSetRequestNormal))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]
	out.AttributeInfo, err = DecodeAttributeDescriptor(&src)
	if err != nil {
		return
	}

	haveAccDesc := src[0]
	src = src[1:]
	// SelectiveAccessInfo
	if haveAccDesc == 0 {
		var nilAccsDesc *SelectiveAccessDescriptor
		out.SelectiveAccessInfo = nilAccsDesc
	} else {
		accDesc, e := DecodeSelectiveAccessDescriptor(&src)
		if e != nil {
			err = e
			return
		}
		out.SelectiveAccessInfo = &accDesc
	}

	decoder := axdr.NewDataDecoder(&src)
	out.Value, err = decoder.Decode(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
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

func (sr SetRequestWithFirstDataBlock) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithFirstDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	attInfo, e := sr.AttributeInfo.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(attInfo)
	if sr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		selInfo, e := sr.SelectiveAccessInfo.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(selInfo)
	}

	val, eVal := sr.DataBlock.Encode()
	if eVal != nil {
		err = eVal
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeSetRequestWithFirstDataBlock(ori *[]byte) (out SetRequestWithFirstDataBlock, err error) {
	var src = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagSetRequest))
		return
	}
	if src[1] != TagSetRequestWithFirstDataBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagSetRequestWithFirstDataBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]
	out.AttributeInfo, err = DecodeAttributeDescriptor(&src)
	if err != nil {
		return
	}

	haveAccDesc := src[0]
	src = src[1:]

	if haveAccDesc == 0 {
		var nilAccsDesc *SelectiveAccessDescriptor
		out.SelectiveAccessInfo = nilAccsDesc
	} else {
		accDesc, e := DecodeSelectiveAccessDescriptor(&src)
		if e != nil {
			err = e
			return
		}
		out.SelectiveAccessInfo = &accDesc
	}

	out.DataBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
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

func (sr SetRequestWithDataBlock) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	val, e := sr.DataBlock.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeSetRequestWithDataBlock(ori *[]byte) (out SetRequestWithDataBlock, err error) {
	var src = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagSetRequest))
		return
	}
	if src[1] != TagSetRequestWithDataBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagSetRequestWithDataBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	out.DataBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// SetRequestWithList implement CosemPDU
type SetRequestWithList struct {
	InvokePriority    uint8
	AttributeCount    uint8
	AttributeInfoList []AttributeDescriptorWithSelection
	ValueCount        uint8
	ValueList         []axdr.DlmsData
}

func CreateSetRequestWithList(invokeId uint8, attList []AttributeDescriptorWithSelection, valList []axdr.DlmsData) *SetRequestWithList {
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

func (sr SetRequestWithList) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithList))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.WriteByte(byte(sr.AttributeCount))
	for _, attr := range sr.AttributeInfoList {
		attInfo, e := attr.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(attInfo)
	}
	buf.WriteByte(byte(sr.ValueCount))
	for _, val := range sr.ValueList {
		dt, e := val.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(dt)
	}

	out = buf.Bytes()
	return
}

func DecodeSetRequestWithList(ori *[]byte) (out SetRequestWithList, err error) {
	var src = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagSetRequest))
		return
	}
	if src[1] != TagSetRequestWithList.Value() {
		err = ErrWrongTag(1, src[1], byte(TagSetRequestWithList))
		return
	}
	out.InvokePriority = src[2]

	out.AttributeCount = uint8(src[3])
	src = src[4:]
	for i := 0; i < int(out.AttributeCount); i++ {
		v, e := DecodeAttributeDescriptorWithSelection(&src)
		if e != nil {
			err = e
			return
		}
		out.AttributeInfoList = append(out.AttributeInfoList, v)
	}

	out.ValueCount = uint8(src[0])
	src = src[1:]
	for i := 0; i < int(out.ValueCount); i++ {
		decoder := axdr.NewDataDecoder(&src)
		v, e := decoder.Decode(&src)
		if e != nil {
			err = e
			return
		}
		out.ValueList = append(out.ValueList, v)
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// SetRequestWithListAndFirstDataBlock implement CosemPDU
type SetRequestWithListAndFirstDataBlock struct {
	InvokePriority    uint8
	AttributeCount    uint8
	AttributeInfoList []AttributeDescriptorWithSelection
	DataBlock         DataBlockSA
}

func CreateSetRequestWithListAndFirstDataBlock(invokeId uint8, attList []AttributeDescriptorWithSelection, dt DataBlockSA) *SetRequestWithListAndFirstDataBlock {
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

func (sr SetRequestWithListAndFirstDataBlock) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagSetRequest))
	buf.WriteByte(byte(TagSetRequestWithListAndFirstDataBlock))
	buf.WriteByte(byte(sr.InvokePriority))
	buf.WriteByte(byte(sr.AttributeCount))
	for _, attr := range sr.AttributeInfoList {
		val, e := attr.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(val)
	}
	val, e := sr.DataBlock.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeSetRequestWithListAndFirstDataBlock(ori *[]byte) (out SetRequestWithListAndFirstDataBlock, err error) {
	var src = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagSetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagSetRequest))
		return
	}
	if src[1] != TagSetRequestWithListAndFirstDataBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagSetRequestWithListAndFirstDataBlock))
		return
	}
	out.InvokePriority = src[2]

	out.AttributeCount = uint8(src[3])
	src = src[4:]
	for i := 0; i < int(out.AttributeCount); i++ {
		v, e := DecodeAttributeDescriptorWithSelection(&src)
		if e != nil {
			err = e
			return
		}
		out.AttributeInfoList = append(out.AttributeInfoList, v)
	}

	out.DataBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
