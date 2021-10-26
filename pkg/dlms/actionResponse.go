package dlms

import (
	"bytes"
	"fmt"
	"gosem/pkg/axdr"
)

type actionResponseTag uint8

const (
	TagActionResponseNormal     actionResponseTag = 0x1
	TagActionResponseWithPBlock actionResponseTag = 0x2
	TagActionResponseWithList   actionResponseTag = 0x3
	TagActionResponseNextPBlock actionResponseTag = 0x4
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s actionResponseTag) Value() uint8 {
	return uint8(s)
}

// ActionResponse implement CosemI
type ActionResponse struct{}

func (gr *ActionResponse) New(tag actionResponseTag) (out CosemPDU, err error) {
	switch tag {
	case TagActionResponseNormal:
		out = &ActionResponseNormal{}
	case TagActionResponseWithPBlock:
		out = &ActionResponseWithPBlock{}
	case TagActionResponseWithList:
		out = &ActionResponseWithList{}
	case TagActionResponseNextPBlock:
		out = &ActionResponseNextPBlock{}
	default:
		err = fmt.Errorf("tag not recognized")
	}
	return
}

func (gr *ActionResponse) Decode(src *[]byte) (out CosemPDU, err error) {
	if (*src)[0] != TagActionResponse.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagActionResponse))
		return
	}

	switch (*src)[1] {
	case TagActionResponseNormal.Value():
		out, err = DecodeActionResponseNormal(src)
	case TagActionResponseWithPBlock.Value():
		out, err = DecodeActionResponseWithPBlock(src)
	case TagActionResponseWithList.Value():
		out, err = DecodeActionResponseWithList(src)
	case TagActionResponseNextPBlock.Value():
		out, err = DecodeActionResponseNextPBlock(src)
	default:
		err = fmt.Errorf("byte tag not recognized (%v)", (*src)[1])
	}

	return
}

type ActionResponseNormal struct {
	InvokePriority uint8
	Response       ActResponse
}

func CreateActionResponseNormal(invokeId uint8, res ActResponse) *ActionResponseNormal {
	return &ActionResponseNormal{
		InvokePriority: invokeId,
		Response:       res,
	}
}

func (ar ActionResponseNormal) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionResponse))
	buf.WriteByte(byte(TagActionResponseNormal))
	buf.WriteByte(byte(ar.InvokePriority))
	val, e := ar.Response.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeActionResponseNormal(ori *[]byte) (out ActionResponseNormal, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagActionResponse.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionResponse))
		return
	}
	if src[1] != TagActionResponseNormal.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionResponseNormal))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	out.Response, err = DecodeActResponse(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

type ActionResponseWithPBlock struct {
	InvokePriority uint8
	PBlock         DataBlockSA
}

func CreateActionResponseWithPBlock(invokeId uint8, dt DataBlockSA) *ActionResponseWithPBlock {
	return &ActionResponseWithPBlock{
		InvokePriority: invokeId,
		PBlock:         dt,
	}
}

func (ar ActionResponseWithPBlock) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionResponse))
	buf.WriteByte(byte(TagActionResponseWithPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	val, e := ar.PBlock.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeActionResponseWithPBlock(ori *[]byte) (out ActionResponseWithPBlock, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagActionResponse.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionResponse))
		return
	}
	if src[1] != TagActionResponseWithPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionResponseWithPBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	out.PBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

type ActionResponseWithList struct {
	InvokePriority uint8
	ResponseCount  uint8
	ResponseList   []ActResponse
}

func CreateActionResponseWithList(invokeId uint8, resList []ActResponse) *ActionResponseWithList {
	if len(resList) < 1 || len(resList) > 255 {
		panic("ResponseList cannot have zero or >255 member")
	}
	return &ActionResponseWithList{
		InvokePriority: invokeId,
		ResponseCount:  uint8(len(resList)),
		ResponseList:   resList,
	}
}

func (ar ActionResponseWithList) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionResponse))
	buf.WriteByte(byte(TagActionResponseWithList))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.WriteByte(byte(ar.ResponseCount))
	for _, res := range ar.ResponseList {
		val, e := res.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(val)
	}

	out = buf.Bytes()
	return
}

func DecodeActionResponseWithList(ori *[]byte) (out ActionResponseWithList, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagActionResponse.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionResponse))
		return
	}
	if src[1] != TagActionResponseWithList.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionResponseWithList))
		return
	}
	out.InvokePriority = src[2]

	out.ResponseCount = uint8(src[3])
	src = src[4:]
	for i := 0; i < int(out.ResponseCount); i++ {
		v, e := DecodeActResponse(&src)
		if e != nil {
			err = e
			return
		}
		out.ResponseList = append(out.ResponseList, v)
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

type ActionResponseNextPBlock struct {
	InvokePriority uint8
	BlockNum       uint32
}

func CreateActionResponseNextPBlock(invokeId uint8, blockNum uint32) *ActionResponseNextPBlock {
	return &ActionResponseNextPBlock{
		InvokePriority: invokeId,
		BlockNum:       blockNum,
	}
}

func (ar ActionResponseNextPBlock) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionResponse))
	buf.WriteByte(byte(TagActionResponseNextPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	blockNum, _ := axdr.EncodeDoubleLongUnsigned(ar.BlockNum)
	buf.Write(blockNum)

	out = buf.Bytes()
	return
}

func DecodeActionResponseNextPBlock(ori *[]byte) (out ActionResponseNextPBlock, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] != TagActionResponse.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionResponse))
		return
	}
	if src[1] != TagActionResponseNextPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionResponseNextPBlock))
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
