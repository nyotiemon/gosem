package dlms

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"gosem/pkg/axdr"
)

type GetDataResult struct {
	IsData bool
	Value  interface{}
}

func CreateGetDataResultAsData(value axdr.DlmsData) *GetDataResult {
	return &GetDataResult{true, value}
}

func CreateGetDataResultAsResult(value AccessResultTag) *GetDataResult {
	return &GetDataResult{false, value}
}

func CreateGetDataResult(value interface{}) *GetDataResult {
	switch val := value.(type) {
	case axdr.DlmsData:
		return CreateGetDataResultAsData(val)
	case AccessResultTag:
		return CreateGetDataResultAsResult(val)
	default:
		panic("Value must be either axdr.DlmsData or AccessResultTag")
	}
}

func (dt GetDataResult) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	if dt.IsData == true {
		buf.WriteByte(0x1)
		value := dt.Value.(axdr.DlmsData)
		enc, e := value.Encode()
		if err != nil {
			err = e
			return
		}
		buf.Write(enc)
	} else {
		buf.WriteByte(0x0)
		value := dt.Value.(AccessResultTag)
		buf.WriteByte(byte(value))
	}

	out = buf.Bytes()
	return
}

func (dt GetDataResult) ValueAsData() (out axdr.DlmsData, err error) {
	if !dt.IsData {
		err = fmt.Errorf("value is DataAccessResult!")
		return
	}
	out = dt.Value.(axdr.DlmsData)

	return
}

func (dt GetDataResult) ValueAsAccess() (out AccessResultTag, err error) {
	if dt.IsData {
		err = fmt.Errorf("value is axdr.DlmsData!")
		return
	}
	out = dt.Value.(AccessResultTag)

	return
}

func DecodeGetDataResult(ori *[]byte) (out GetDataResult, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] == 0x0 {
		out.IsData = false
		out.Value, err = GetAccessTag(uint8(src[1]))
		if err == nil {
			src = src[2:]
		}
	} else {
		out.IsData = true
		src = src[1:]
		decoder := axdr.NewDataDecoder(&src)
		out.Value, err = decoder.Decode(&src)
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// DataBlockG is DataBlock for the GET-response. Result must be either byte slice
// or AccessResultTag after creation, or else it will fail on Encode()
type DataBlockG struct {
	LastBlock   bool
	BlockNumber uint32
	IsResult    bool
	Result      interface{}
}

func CreateDataBlockGAsData(lastBlock bool, blockNum uint32, result interface{}) *DataBlockG {
	switch res := result.(type) {
	case string:
		bt, e := hex.DecodeString(res)
		if e != nil {
			panic(e)
		}
		return &DataBlockG{lastBlock, blockNum, false, bt}

	case []byte:
		return &DataBlockG{lastBlock, blockNum, false, res}

	default:
		panic("CreateDataBlockGAsData result only accept hexstring or byte slice.")
	}
}

func CreateDataBlockGAsResult(lastBlock bool, blockNum uint32, result AccessResultTag) *DataBlockG {
	return &DataBlockG{lastBlock, blockNum, true, result}
}

func CreateDataBlockG(lastBlock bool, blockNum uint32, result interface{}) *DataBlockG {
	switch res := result.(type) {
	case string:
		bt, e := hex.DecodeString(res)
		if e != nil {
			panic(e)
		}
		return CreateDataBlockGAsData(lastBlock, blockNum, bt)

	case []byte:
		return CreateDataBlockGAsData(lastBlock, blockNum, res)

	case AccessResultTag:
		return CreateDataBlockGAsResult(lastBlock, blockNum, res)

	default:
		panic("DataBlockG result only accept hexstring, byte slice, or DataAccessResult.")
	}
}

func (dt DataBlockG) Encode() (out []byte, err error) {
	var buf bytes.Buffer

	if dt.LastBlock {
		buf.WriteByte(0x1)
	} else {
		buf.WriteByte(0x0)
	}

	blk, e := axdr.EncodeDoubleLongUnsigned(dt.BlockNumber)
	if e != nil {
		err = e
		return
	}
	buf.Write(blk)

	if dt.IsResult == true {
		buf.WriteByte(0x1)
		value := dt.Result.(AccessResultTag)
		buf.WriteByte(byte(value))
	} else {
		buf.WriteByte(0x0)
		value := dt.Result.([]byte)
		// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
		buf.WriteByte(byte(len(value)))
		buf.Write(value)
	}

	out = buf.Bytes()
	return
}

func (dt DataBlockG) ResultAsBytes() (out []byte, err error) {
	if dt.IsResult {
		err = fmt.Errorf("value is DataAccessResult!")
	} else {
		out = dt.Result.([]byte)
	}

	return
}

func (dt DataBlockG) ResultAsAccess() (out AccessResultTag, err error) {
	if !dt.IsResult {
		err = fmt.Errorf("value is byte slice!")
	} else {
		out = dt.Result.(AccessResultTag)
	}

	return
}

func DecodeDataBlockG(ori *[]byte) (out DataBlockG, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] == 0x0 {
		out.LastBlock = false
	} else {
		out.LastBlock = true
	}
	src = src[1:]

	_, out.BlockNumber, err = axdr.DecodeDoubleLongUnsigned(&src)

	if src[0] == 0x0 {
		out.IsResult = false
	} else {
		out.IsResult = true
	}
	src = src[1:]

	if out.IsResult {
		out.Result, err = GetAccessTag(uint8(src[0]))
		src = src[0:]
	} else {
		// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
		val := uint8(src[0])
		out.Result = src[1 : 1+val]
		src = src[1+val:]
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// DataBlockSA is DataBlock for the SET-request, ACTION-request and ACTION-response
type DataBlockSA struct {
	LastBlock   bool
	BlockNumber uint32
	Raw         []byte
}

func CreateDataBlockSA(lastBlock bool, blockNum uint32, result interface{}) *DataBlockSA {
	switch res := result.(type) {
	case string:
		bt, e := hex.DecodeString(res)
		if e != nil {
			panic(e)
		}
		return &DataBlockSA{lastBlock, blockNum, bt}

	case []byte:
		return &DataBlockSA{lastBlock, blockNum, res}

	default:
		panic("DataBlockSA result only accept hexstring or byte slice.")
	}
}

func (dt DataBlockSA) Encode() (out []byte, err error) {
	var buf bytes.Buffer

	if dt.LastBlock {
		buf.WriteByte(0x1)
	} else {
		buf.WriteByte(0x0)
	}

	blk, e := axdr.EncodeDoubleLongUnsigned(dt.BlockNumber)
	if e != nil {
		err = e
		return
	}
	buf.Write(blk)

	// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
	buf.WriteByte(byte(len(dt.Raw)))
	buf.Write(dt.Raw)

	out = buf.Bytes()
	return
}

func DecodeDataBlockSA(ori *[]byte) (out DataBlockSA, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] == 0x0 {
		out.LastBlock = false
	} else {
		out.LastBlock = true
	}
	src = src[1:]

	_, out.BlockNumber, err = axdr.DecodeDoubleLongUnsigned(&src)

	// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
	val := uint8(src[0])
	out.Raw = src[1 : val+1]
	src = src[val+1:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}

// Response of ActionRequest. ReturnParam is optional parameter therefore pointer
type ActResponse struct {
	Result      ActionResultTag
	ReturnParam *GetDataResult
}

func CreateActResponse(result ActionResultTag, returnParam *GetDataResult) *ActResponse {

	return &ActResponse{Result: result, ReturnParam: returnParam}
}

func (dt ActResponse) Encode() (out []byte, err error) {
	var buf bytes.Buffer

	buf.WriteByte(byte(dt.Result))

	if dt.ReturnParam == nil {
		buf.WriteByte(0x0)
	} else {
		buf.WriteByte(0x1)
		enc, e := dt.ReturnParam.Encode()
		if e != nil {
			err = e
			return
		}
		buf.Write(enc)
	}

	out = buf.Bytes()
	return
}

func DecodeActResponse(ori *[]byte) (out ActResponse, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	out.Result, err = GetActionTag(src[0])
	if err != nil {
		return
	}

	haveReturnParam := src[1]
	src = src[2:]

	if haveReturnParam == 0x0 {
		var gdrNil *GetDataResult = nil
		out.ReturnParam = gdrNil
	} else {
		gdr, e := DecodeGetDataResult(&src)
		if e != nil {
			err = e
			return
		}
		out.ReturnParam = &gdr
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
