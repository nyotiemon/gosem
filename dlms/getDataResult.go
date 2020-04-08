package cosem

import (
	"bytes"
	"encoding/hex"
	. "gosem/axdr"
)

type GetDataResult struct {
	IsData bool
	Value  interface{}
}

func CreateGetDataResultAsData(value DlmsData) *GetDataResult {
	return &GetDataResult{true, value}
}

func CreateGetDataResultAsResult(value AccessResultTag) *GetDataResult {
	return &GetDataResult{false, value}
}

func CreateGetDataResult(value interface{}) *GetDataResult {
	switch val := value.(type) {
	case DlmsData:
		return CreateGetDataResultAsData(val)
	case AccessResultTag:
		return CreateGetDataResultAsResult(val)
	default:
		panic("Value must be either DlmsData or AccessResultTag")
	}
}

func (dt *GetDataResult) Encode() []byte {
	var output bytes.Buffer
	if dt.IsData == true {
		output.WriteByte(0x1)
		value := dt.Value.(DlmsData)
		output.Write(value.Encode())
	} else {
		output.WriteByte(0x0)
		value := dt.Value.(AccessResultTag)
		output.WriteByte(byte(value))
	}

	return output.Bytes()
}

func (dt *GetDataResult) ValueAsData() DlmsData {
	if !dt.IsData {
		panic("Value is DataAccessResult!")
	}

	return dt.Value.(DlmsData)
}

func (dt *GetDataResult) ValueAsAccess() AccessResultTag {
	if dt.IsData {
		panic("Value is DlmsData!")
	}

	return dt.Value.(AccessResultTag)
}

func DecodeGetDataResult(src *[]byte) (out GetDataResult, err error) {
	if (*src)[0] == 0x0 {
		out.IsData = false
		out.Value, err = GetAccessTag(uint8((*src)[1]))
		if err == nil {
			(*src) = (*src)[2:]
		}
	} else {
		out.IsData = true
		(*src) = (*src)[1:]
		decoder := NewDataDecoder(src)
		out.Value, err = decoder.Decode(src)
	}

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

func (dt *DataBlockG) Encode() []byte {
	var output bytes.Buffer

	if dt.LastBlock {
		output.WriteByte(0x1)
	} else {
		output.WriteByte(0x0)
	}

	blk, e := EncodeDoubleLongUnsigned(dt.BlockNumber)
	if e != nil {
		panic(e)
	}
	output.Write(blk)

	if dt.IsResult == true {
		output.WriteByte(0x1)
		value := dt.Result.(AccessResultTag)
		output.WriteByte(byte(value))
	} else {
		output.WriteByte(0x0)
		value := dt.Result.([]byte)
		// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
		output.WriteByte(byte(len(value)))
		output.Write(value)
	}

	return output.Bytes()
}

func (dt *DataBlockG) ResultAsBytes() []byte {
	if dt.IsResult {
		panic("Value is DataAccessResult!")
	}

	return dt.Result.([]byte)
}

func (dt *DataBlockG) ResultAsAccess() AccessResultTag {
	if !dt.IsResult {
		panic("Value is byte slice!")
	}

	return dt.Result.(AccessResultTag)
}

func DecodeDataBlockG(src *[]byte) (out DataBlockG, err error) {
	if (*src)[0] == 0x0 {
		out.LastBlock = false
	} else {
		out.LastBlock = true
	}
	(*src) = (*src)[1:]

	_, out.BlockNumber, err = DecodeDoubleLongUnsigned(src)

	if (*src)[0] == 0x0 {
		out.IsResult = false
	} else {
		out.IsResult = true
	}
	(*src) = (*src)[1:]

	if out.IsResult {
		out.Result, err = GetAccessTag(uint8((*src)[0]))
		(*src) = (*src)[0:]
	} else {
		// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
		val := uint8((*src)[0])
		out.Result = (*src)[1 : 1+val]
		(*src) = (*src)[1+val:]
	}

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

func (dt *DataBlockSA) Encode() []byte {
	var output bytes.Buffer

	if dt.LastBlock {
		output.WriteByte(0x1)
	} else {
		output.WriteByte(0x0)
	}

	blk, e := EncodeDoubleLongUnsigned(dt.BlockNumber)
	if e != nil {
		panic(e)
	}
	output.Write(blk)

	// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
	output.WriteByte(byte(len(dt.Raw)))
	output.Write(dt.Raw)

	return output.Bytes()
}

func DecodeDataBlockSA(src *[]byte) (out DataBlockSA, err error) {
	if (*src)[0] == 0x0 {
		out.LastBlock = false
	} else {
		out.LastBlock = true
	}
	(*src) = (*src)[1:]

	_, out.BlockNumber, err = DecodeDoubleLongUnsigned(src)

	// not sure if length is limited only 1 byte, or does it follow KLV as in axdr
	val := uint8((*src)[0])
	out.Raw = (*src)[1 : val+1]
	(*src) = (*src)[val+1:]

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

func (dt *ActResponse) Encode() []byte {
	var output bytes.Buffer

	output.WriteByte(byte(dt.Result))

	if dt.ReturnParam == nil {
		output.WriteByte(0x0)
	} else {
		output.WriteByte(0x1)
		output.Write(dt.ReturnParam.Encode())
	}

	return output.Bytes()
}

func DecodeActResponse(src *[]byte) (out ActResponse, err error) {
	out.Result, err = GetActionTag((*src)[0])
	if err != nil {
		return
	}

	haveReturnParam := (*src)[1]
	(*src) = (*src)[2:]

	if haveReturnParam == 0x0 {
		var gdrNil *GetDataResult = nil
		out.ReturnParam = gdrNil
	} else {
		gdr, e := DecodeGetDataResult(src)
		if e != nil {
			err = e
			return
		}
		out.ReturnParam = &gdr
	}

	return
}
