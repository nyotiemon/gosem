package cosem

import (
	"bytes"
	. "gosem/axdr"
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

func (gr *GetRequest) New(tag getRequestTag) CosemPDU {
	switch tag {
	case TagGetRequestNormal:
		return &GetRequestNormal{}
	case TagGetRequestNext:
		return &GetRequestNext{}
	case TagGetRequestWithList:
		return &GetRequestWithList{}
	default:
		panic("Tag not recognized!")
	}
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

func (gr GetRequestNormal) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestNormal))
	buf.WriteByte(byte(gr.InvokePriority))
	buf.Write(gr.AttributeInfo.Encode())
	if gr.SelectiveAccessInfo == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		buf.Write(gr.SelectiveAccessInfo.Encode())
	}

	return buf.Bytes()
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

func (gr GetRequestNext) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestNext))
	buf.WriteByte(byte(gr.InvokePriority))
	blockNum, _ := EncodeDoubleLongUnsigned(gr.BlockNum)
	buf.Write(blockNum)

	return buf.Bytes()
}

// GetRequestWithList implement CosemPDU
type GetRequestWithList struct {
	InvokePriority    uint8
	AttributeInfoList []AttributeDescriptor
}

func CreateGetRequestWithList(invokeId uint8, attList []AttributeDescriptor) *GetRequestWithList {
	if len(attList) < 1 {
		panic("AttributeInfoList cannot have zero member")
	}
	return &GetRequestWithList{
		InvokePriority:    invokeId,
		AttributeInfoList: attList,
	}
}

func (gr GetRequestWithList) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetRequest))
	buf.WriteByte(byte(TagGetRequestWithList))
	buf.WriteByte(byte(gr.InvokePriority))
	for _, attr := range gr.AttributeInfoList {
		buf.Write(attr.Encode())
	}

	return buf.Bytes()
}

func DecodeGetRequestNormal(src *[]byte) (out GetRequestNormal, err error) {
	if (*src)[0] != TagGetRequest.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetRequest))
		return
	}
	if (*src)[1] != TagGetRequestNormal.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetRequestNormal))
		return
	}
	out.InvokePriority = (*src)[2]
	(*src) = (*src)[3:]
	out.AttributeInfo, err = DecodeAttributeDescriptor(src)
	if err != nil {
		return
	}

	haveAccDesc := (*src)[0]
	(*src) = (*src)[1:]
	// SelectiveAccessInfo
	if haveAccDesc == 0 {
		var nilAccsDesc *SelectiveAccessDescriptor = nil
		out.SelectiveAccessInfo = nilAccsDesc
	} else {
		accDesc, e := DecodeSelectiveAccessDescriptor(src)
		if e != nil {
			err = e
			return
		}
		out.SelectiveAccessInfo = &accDesc
	}

	return
}
