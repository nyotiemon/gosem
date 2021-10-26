package dlms

import (
	"bytes"
	"fmt"
	"gosem/pkg/axdr"
)

type getRequestTag uint8

const (
	TagGetRequestNormal   getRequestTag = 0x1
	TagGetRequestNext     getRequestTag = 0x2
	TagGetRequestWithList getRequestTag = 0x3
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s getRequestTag) Value() uint8 {
	return uint8(s)
}

// GetRequest implement CosemI
type GetRequest struct{}

func (gr *GetRequest) New(tag getRequestTag) (out CosemPDU, err error) {
	switch tag {
	case TagGetRequestNormal:
		out = &GetRequestNormal{}
	case TagGetRequestNext:
		out = &GetRequestNext{}
	case TagGetRequestWithList:
		out = &GetRequestWithList{}
	default:
		err = fmt.Errorf("tag not recognized")
	}

	return
}

func (gr *GetRequest) Decode(src *[]byte) (out CosemPDU, err error) {
	if (*src)[0] != TagGetRequest.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetRequest))
		return
	}

	switch (*src)[1] {
	case TagGetRequestNormal.Value():
		out, err = DecodeGetRequestNormal(src)
	case TagGetRequestNext.Value():
		out, err = DecodeGetRequestNext(src)
	case TagGetRequestWithList.Value():
		out, err = DecodeGetRequestWithList(src)
	default:
		err = fmt.Errorf("byte tag not recognized (%v)", (*src)[1])
	}

	return
}

// GetRequestNormal implement CosemPDU. SelectiveAccessDescriptor is optional
type GetRequestNormal struct {
	InvokePriority      uint8
	AttributeInfo       AttributeDescriptor
	SelectiveAccessInfo *SelectiveAccessDescriptor
}

func CreateGetRequestNormal(invokeId uint8, att AttributeDescriptor, acc *SelectiveAccessDescriptor) *GetRequestNormal {
	return &GetRequestNormal{
		InvokePriority:      invokeId,
		AttributeInfo:       att,
		SelectiveAccessInfo: acc,
	}
}

func (gr GetRequestNormal) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestNormal))
	buf.WriteByte(byte(gr.InvokePriority))
	attInfo, e := gr.AttributeInfo.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(attInfo)
	if gr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		selInfo, e := gr.SelectiveAccessInfo.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(selInfo)
	}

	out = buf.Bytes()
	return
}

func DecodeGetRequestNormal(ori *[]byte) (out GetRequestNormal, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagGetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagGetRequest))
		return
	}
	if src[1] != TagGetRequestNormal.Value() {
		err = ErrWrongTag(1, src[1], byte(TagGetRequestNormal))
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

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// GetRequestNext implement CosemPDU
type GetRequestNext struct {
	InvokePriority uint8
	BlockNum       uint32
}

func CreateGetRequestNext(invokeId uint8, blockNum uint32) *GetRequestNext {
	return &GetRequestNext{
		InvokePriority: invokeId,
		BlockNum:       blockNum,
	}
}

func (gr GetRequestNext) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestNext))
	buf.WriteByte(byte(gr.InvokePriority))
	blockNum, _ := axdr.EncodeDoubleLongUnsigned(gr.BlockNum)
	buf.Write(blockNum)

	out = buf.Bytes()
	return
}

func DecodeGetRequestNext(ori *[]byte) (out GetRequestNext, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagGetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagGetRequest))
		return
	}
	if src[1] != TagGetRequestNext.Value() {
		err = ErrWrongTag(1, src[1], byte(TagGetRequestNext))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	_, v, e := axdr.DecodeDoubleLongUnsigned(&src)
	if e != nil {
		err = e
		return
	}
	out.BlockNum = v

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// GetRequestWithList implement CosemPDU
type GetRequestWithList struct {
	InvokePriority    uint8
	AttributeCount    uint8
	AttributeInfoList []AttributeDescriptorWithSelection
}

func CreateGetRequestWithList(invokeId uint8, attList []AttributeDescriptorWithSelection) *GetRequestWithList {
	if len(attList) < 1 || len(attList) > 255 {
		panic("AttributeInfoList cannot have zero or >255 member")
	}
	return &GetRequestWithList{
		InvokePriority:    invokeId,
		AttributeCount:    uint8(len(attList)),
		AttributeInfoList: attList,
	}
}

func (gr GetRequestWithList) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestWithList))
	buf.WriteByte(byte(gr.InvokePriority))
	buf.WriteByte(byte(len(gr.AttributeInfoList)))
	for _, attr := range gr.AttributeInfoList {
		val, e := attr.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(val)
	}

	out = buf.Bytes()
	return
}

func DecodeGetRequestWithList(ori *[]byte) (out GetRequestWithList, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagGetRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagGetRequest))
		return
	}
	if src[1] != TagGetRequestWithList.Value() {
		err = ErrWrongTag(1, src[1], byte(TagGetRequestWithList))
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

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
