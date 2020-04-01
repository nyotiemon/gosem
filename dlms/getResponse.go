package cosem

import (
	"bytes"
)

type getResponseTag uint8

const (
	TagGetResponseNormal        getResponseTag = 0x1
	TagGetResponseWithDataBlock getResponseTag = 0x2
	TagGetResponseWithList      getResponseTag = 0x3
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s getResponseTag) Value() uint8 {
	return uint8(s)
}

// GetResponse implement CosemI
type GetResponse struct{}

func (gr *GetResponse) New(tag getResponseTag) CosemPDU {
	switch tag {
	case TagGetResponseNormal:
		return &GetResponseNormal{}
	case TagGetResponseWithDataBlock:
		return &GetResponseWithDataBlock{}
	case TagGetResponseWithList:
		return &GetResponseWithList{}
	default:
		panic("Tag not recognized!")
	}
}

// GetResponseNormal implement CosemPDU. SelectiveAccessDescriptor is optional
type GetResponseNormal struct {
	InvokePriority uint8
	Result         GetDataResult
}

func CreateGetResponseNormal(invokeId uint8, res GetDataResult) *GetResponseNormal {
	return &GetResponseNormal{
		InvokePriority: invokeId,
		Result:         res,
	}
}

func (gr GetResponseNormal) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetResponse))
	buf.WriteByte(byte(TagGetResponseNormal))
	buf.WriteByte(byte(gr.InvokePriority))
	buf.Write(gr.Result.Encode())

	return buf.Bytes()
}

// GetResponseNext implement CosemPDU
type GetResponseWithDataBlock struct {
	InvokePriority uint8
	Result         DataBlockG
}

func CreateGetResponseWithDataBlock(invokeId uint8, res DataBlockG) *GetResponseWithDataBlock {
	return &GetResponseWithDataBlock{
		InvokePriority: invokeId,
		Result:         res,
	}
}

func (gr GetResponseWithDataBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetResponse))
	buf.WriteByte(byte(TagGetResponseWithDataBlock))
	buf.WriteByte(byte(gr.InvokePriority))
	buf.Write(gr.Result.Encode())

	return buf.Bytes()
}

// GetResponseWithList implement CosemPDU
type GetResponseWithList struct {
	InvokePriority uint8
	ResultList     []GetDataResult
}

func CreateGetResponseWithList(invokeId uint8, resList []GetDataResult) *GetResponseWithList {
	if len(resList) < 1 {
		panic("ResultList cannot have zero member")
	}
	return &GetResponseWithList{
		InvokePriority: invokeId,
		ResultList:     resList,
	}
}

func (gr GetResponseWithList) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetResponse))
	buf.WriteByte(byte(TagGetResponseWithList))
	buf.WriteByte(byte(gr.InvokePriority))
	for _, res := range gr.ResultList {
		buf.Write(res.Encode())
	}

	return buf.Bytes()
}

func DecodeGetResponseNormal(src *[]byte) (out GetResponseNormal, err error) {
	if (*src)[0] != TagGetResponse.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponse))
		return
	}
	if (*src)[1] != TagGetResponseNormal.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponseNormal))
		return
	}
	out.InvokePriority = (*src)[2]
	(*src) = (*src)[3:]

	out.Result, err = DecodeGetDataResult(src)

	return
}
