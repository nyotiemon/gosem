package cosem

import (
	"bytes"
	"fmt"
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

func (gr *GetResponse) Decode(src *[]byte) (out CosemPDU, err error) {
	if (*src)[0] != TagGetResponse.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponse))
		return
	}

	switch (*src)[1] {
	case TagGetResponseNormal.Value():
		out, err = DecodeGetResponseNormal(src)
	case TagGetResponseWithDataBlock.Value():
		out, err = DecodeGetResponseWithDataBlock(src)
	case TagGetResponseWithList.Value():
		out, err = DecodeGetResponseWithList(src)
	default:
		err = fmt.Errorf("Byte tag not recognized (%v)", (*src)[1])
	}

	return
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

func DecodeGetResponseWithDataBlock(src *[]byte) (out GetResponseWithDataBlock, err error) {
	if (*src)[0] != TagGetResponse.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponse))
		return
	}
	if (*src)[1] != TagGetResponseWithDataBlock.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponseWithDataBlock))
		return
	}
	out.InvokePriority = (*src)[2]
	(*src) = (*src)[3:]

	out.Result, err = DecodeDataBlockG(src)
	return
}

// GetResponseWithList implement CosemPDU
type GetResponseWithList struct {
	InvokePriority uint8
	ResultCount    uint8
	ResultList     []GetDataResult
}

func CreateGetResponseWithList(invokeId uint8, resList []GetDataResult) *GetResponseWithList {
	if len(resList) < 1 || len(resList) > 255 {
		panic("ResultList cannot have zero or >255 member")
	}
	return &GetResponseWithList{
		InvokePriority: invokeId,
		ResultCount:    uint8(len(resList)),
		ResultList:     resList,
	}
}

func (gr GetResponseWithList) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagGetResponse))
	buf.WriteByte(byte(TagGetResponseWithList))
	buf.WriteByte(byte(gr.InvokePriority))
	buf.WriteByte(byte(len(gr.ResultList)))
	for _, res := range gr.ResultList {
		buf.Write(res.Encode())
	}

	return buf.Bytes()
}

func DecodeGetResponseWithList(src *[]byte) (out GetResponseWithList, err error) {
	if (*src)[0] != TagGetResponse.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponse))
		return
	}
	if (*src)[1] != TagGetResponseWithList.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagGetResponseWithList))
		return
	}
	out.InvokePriority = (*src)[2]

	out.ResultCount = uint8((*src)[3])
	(*src) = (*src)[4:]
	for i := 0; i < int(out.ResultCount); i++ {
		v, e := DecodeGetDataResult(src)
		if e != nil {
			err = e
			return
		}
		out.ResultList = append(out.ResultList, v)
	}

	return
}
