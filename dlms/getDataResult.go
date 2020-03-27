package cosem

import (
	"bytes"
	. "gosem/axdr"
)

type GetDataResult struct {
	IsData bool
	Value  interface{}
}

func CreateGetDataResultAsData(value DlmsData) *GetDataResult {
	return &GetDataResult{true, value}
}

func CreateGetDataResultAsResult(value accessResultTag) *GetDataResult {
	return &GetDataResult{false, value}
}

func CreateGetDataResult(value interface{}) *GetDataResult {
	switch val := value.(type) {
	case DlmsData:
		return CreateGetDataResultAsData(val)
	case accessResultTag:
		return CreateGetDataResultAsResult(val)
	default:
		panic("Value must be either DlmsData or accessResultTag")
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
		value := dt.Value.(accessResultTag)
		output.WriteByte(byte(value))
	}

	return output.Bytes()
}
