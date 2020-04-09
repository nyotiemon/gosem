package cosem

import (
	"bytes"
	"fmt"
	. "gosem/axdr"
)

type actionRequestTag uint8

const (
	TagActionRequestNormal                 actionRequestTag = 0x1
	TagActionRequestNextPBlock             actionRequestTag = 0x2
	TagActionRequestWithList               actionRequestTag = 0x3
	TagActionRequestWithFirstPBlock        actionRequestTag = 0x4
	TagActionRequestWithListAndFirstPBlock actionRequestTag = 0x5
	TagActionRequestWithPBlock             actionRequestTag = 0x6
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s actionRequestTag) Value() uint8 {
	return uint8(s)
}

// ActionRequest implement CosemI
type ActionRequest struct{}

func (gr *ActionRequest) New(tag actionRequestTag) CosemPDU {
	switch tag {
	case TagActionRequestNormal:
		return &ActionRequestNormal{}
	case TagActionRequestNextPBlock:
		return &ActionRequestNextPBlock{}
	case TagActionRequestWithList:
		return &ActionRequestWithList{}
	case TagActionRequestWithFirstPBlock:
		return &ActionRequestWithFirstPBlock{}
	case TagActionRequestWithListAndFirstPBlock:
		return &ActionRequestWithListAndFirstPBlock{}
	case TagActionRequestWithPBlock:
		return &ActionRequestWithPBlock{}
	default:
		panic("Tag not recognized!")
	}
}

func (gr *ActionRequest) Decode(src *[]byte) (out CosemPDU, err error) {
	if (*src)[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagActionRequest))
		return
	}

	switch (*src)[1] {
	case TagActionRequestNormal.Value():
		out, err = DecodeActionRequestNormal(src)
	case TagActionRequestNextPBlock.Value():
		out, err = DecodeActionRequestNextPBlock(src)
	case TagActionRequestWithList.Value():
		out, err = DecodeActionRequestWithList(src)
	case TagActionRequestWithFirstPBlock.Value():
		out, err = DecodeActionRequestWithFirstPBlock(src)
	case TagActionRequestWithListAndFirstPBlock.Value():
		out, err = DecodeActionRequestWithListAndFirstPBlock(src)
	case TagActionRequestWithPBlock.Value():
		out, err = DecodeActionRequestWithPBlock(src)
	default:
		err = fmt.Errorf("byte tag not recognized (%v)", (*src)[1])
	}

	return
}

// ActionRequestNormal implement CosemPDU
type ActionRequestNormal struct {
	InvokePriority uint8
	MethodInfo     MethodDescriptor
	MethodParam    *DlmsData
}

func CreateActionRequestNormal(invokeId uint8, mth MethodDescriptor, dt *DlmsData) *ActionRequestNormal {
	return &ActionRequestNormal{
		InvokePriority: invokeId,
		MethodInfo:     mth,
		MethodParam:    dt,
	}
}

func (ar ActionRequestNormal) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestNormal))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.Write(ar.MethodInfo.Encode())
	if ar.MethodParam == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		buf.Write(ar.MethodParam.Encode())
	}

	return buf.Bytes()
}

func DecodeActionRequestNormal(ori *[]byte) (out ActionRequestNormal, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestNormal.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestNormal))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]
	out.MethodInfo, err = DecodeMethodDescriptor(&src)
	if err != nil {
		return
	}

	haveMethodParam := src[0]
	src = src[1:]
	if haveMethodParam == 0 {
		var nilData *DlmsData = nil
		out.MethodParam = nilData
	} else {
		decoder := NewDataDecoder(&src)
		dt, e := decoder.Decode(&src)
		if e != nil {
			err = e
			return
		}
		out.MethodParam = &dt
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// ActionRequestNextPBlock implement CosemPDU
type ActionRequestNextPBlock struct {
	InvokePriority uint8
	BlockNum       uint32
}

func CreateActionRequestNextPBlock(invokeId uint8, blockNum uint32) *ActionRequestNextPBlock {
	return &ActionRequestNextPBlock{
		InvokePriority: invokeId,
		BlockNum:       blockNum,
	}
}

func (ar ActionRequestNextPBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestNextPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	blockNum, _ := EncodeDoubleLongUnsigned(ar.BlockNum)
	buf.Write(blockNum)

	return buf.Bytes()
}

func DecodeActionRequestNextPBlock(ori *[]byte) (out ActionRequestNextPBlock, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestNextPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestNextPBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	_, v, e := DecodeDoubleLongUnsigned(&src)
	if e != nil {
		err = e
		return
	}
	out.BlockNum = v

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// ActionRequestWithList implement CosemPDU
type ActionRequestWithList struct {
	InvokePriority   uint8
	MethodInfoCount  uint8
	MethodInfoList   []MethodDescriptor
	MethodParamCount uint8
	MethodParamList  []DlmsData
}

func CreateActionRequestWithList(invokeId uint8, mthList []MethodDescriptor, valList []DlmsData) *ActionRequestWithList {
	if len(mthList) < 1 || len(mthList) > 255 {
		panic("MethodInfoList cannot have zero or >255 member")
	}
	if len(valList) < 1 || len(valList) > 255 {
		panic("MethodParamList cannot have zero or >255 member")
	}
	return &ActionRequestWithList{
		InvokePriority:   invokeId,
		MethodInfoCount:  uint8(len(mthList)),
		MethodInfoList:   mthList,
		MethodParamCount: uint8(len(valList)),
		MethodParamList:  valList,
	}
}

func (ar ActionRequestWithList) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestWithList))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.WriteByte(byte(ar.MethodInfoCount))
	for _, attr := range ar.MethodInfoList {
		buf.Write(attr.Encode())
	}
	buf.WriteByte(byte(ar.MethodParamCount))
	for _, val := range ar.MethodParamList {
		buf.Write(val.Encode())
	}

	return buf.Bytes()
}

func DecodeActionRequestWithList(ori *[]byte) (out ActionRequestWithList, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestWithList.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestWithList))
		return
	}
	out.InvokePriority = src[2]

	out.MethodInfoCount = uint8(src[3])
	src = src[4:]
	for i := 0; i < int(out.MethodInfoCount); i++ {
		v, e := DecodeMethodDescriptor(&src)
		if e != nil {
			err = e
			return
		}
		out.MethodInfoList = append(out.MethodInfoList, v)
	}

	out.MethodParamCount = uint8(src[0])
	src = src[1:]
	for i := 0; i < int(out.MethodParamCount); i++ {
		decoder := NewDataDecoder(&src)
		v, e := decoder.Decode(&src)
		if e != nil {
			err = e
			return
		}
		out.MethodParamList = append(out.MethodParamList, v)
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// ActionRequestWithFirstPBlock implement CosemPDU
type ActionRequestWithFirstPBlock struct {
	InvokePriority uint8
	MethodInfo     MethodDescriptor
	PBlock         DataBlockSA
}

func CreateActionRequestWithFirstPBlock(invokeId uint8, mth MethodDescriptor, dt DataBlockSA) *ActionRequestWithFirstPBlock {
	return &ActionRequestWithFirstPBlock{
		InvokePriority: invokeId,
		MethodInfo:     mth,
		PBlock:         dt,
	}
}

func (ar ActionRequestWithFirstPBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestWithFirstPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.Write(ar.MethodInfo.Encode())
	buf.Write(ar.PBlock.Encode())

	return buf.Bytes()
}

func DecodeActionRequestWithFirstPBlock(ori *[]byte) (out ActionRequestWithFirstPBlock, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestWithFirstPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestWithFirstPBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]
	out.MethodInfo, err = DecodeMethodDescriptor(&src)
	if err != nil {
		return
	}

	out.PBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// ActionRequestWithListAndFirstPBlock implement CosemPDU
type ActionRequestWithListAndFirstPBlock struct {
	InvokePriority  uint8
	MethodInfoCount uint8
	MethodInfoList  []MethodDescriptor
	PBlock          DataBlockSA
}

func CreateActionRequestWithListAndFirstPBlock(invokeId uint8, mthList []MethodDescriptor, dt DataBlockSA) *ActionRequestWithListAndFirstPBlock {
	if len(mthList) < 1 || len(mthList) > 255 {
		panic("MethodInfoList cannot have zero or >255 member")
	}
	return &ActionRequestWithListAndFirstPBlock{
		InvokePriority:  invokeId,
		MethodInfoCount: uint8(len(mthList)),
		MethodInfoList:  mthList,
		PBlock:          dt,
	}
}

func (ar ActionRequestWithListAndFirstPBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestWithListAndFirstPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.WriteByte(byte(ar.MethodInfoCount))
	for _, attr := range ar.MethodInfoList {
		buf.Write(attr.Encode())
	}
	buf.Write(ar.PBlock.Encode())

	return buf.Bytes()
}

func DecodeActionRequestWithListAndFirstPBlock(ori *[]byte) (out ActionRequestWithListAndFirstPBlock, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestWithListAndFirstPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestWithListAndFirstPBlock))
		return
	}
	out.InvokePriority = src[2]

	out.MethodInfoCount = uint8(src[3])
	src = src[4:]
	for i := 0; i < int(out.MethodInfoCount); i++ {
		v, e := DecodeMethodDescriptor(&src)
		if e != nil {
			err = e
			return
		}
		out.MethodInfoList = append(out.MethodInfoList, v)
	}

	out.PBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// ActionRequestWithPBlock implement CosemPDU
type ActionRequestWithPBlock struct {
	InvokePriority uint8
	PBlock         DataBlockSA
}

func CreateActionRequestWithPBlock(invokeId uint8, dt DataBlockSA) *ActionRequestWithPBlock {
	return &ActionRequestWithPBlock{
		InvokePriority: invokeId,
		PBlock:         dt,
	}
}

func (ar ActionRequestWithPBlock) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagActionRequest))
	buf.WriteByte(byte(TagActionRequestWithPBlock))
	buf.WriteByte(byte(ar.InvokePriority))
	buf.Write(ar.PBlock.Encode())

	return buf.Bytes()
}

func DecodeActionRequestWithPBlock(ori *[]byte) (out ActionRequestWithPBlock, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagActionRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagActionRequest))
		return
	}
	if src[1] != TagActionRequestWithPBlock.Value() {
		err = ErrWrongTag(1, src[1], byte(TagActionRequestWithPBlock))
		return
	}
	out.InvokePriority = src[2]
	src = src[3:]

	out.PBlock, err = DecodeDataBlockSA(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
