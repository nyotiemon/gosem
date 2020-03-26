package cosem

import (
	"bytes"
	. "gosem/axdr"
)

type GetDataResult struct {
	IsData bool
	Value  interface{}
}

func CreateGetDataResult(isData bool, value interface{}) *GetDataResult {
	if isData == true {
		// val is DlmsData
		if val, ok := value.(DlmsData); ok {
			return &GetDataResult{isData, val}
		} else {
			panic("Value with IsData flag true must be DlmsData")
		}

	} else {
		// val is DataAccessResult
		if val, ok := value.(accessResultTag); ok {
			return &GetDataResult{isData, val}
		} else {
			panic("Value with IsData flag false must be accessResultTag")
		}
	}
}

func CreateGetDataResultAsData(value DlmsData) *GetDataResult {
	return &GetDataResult{true, value}
}

func CreateGetDataResultAsResult(value accessResultTag) *GetDataResult {
	return &GetDataResult{false, value}
}

func (ad *GetDataResult) Encode() []byte {
	var output bytes.Buffer
	if ad.IsData == true {
		output.WriteByte(0x1)
		value := ad.Value.(DlmsData)
		output.Write(value.Encode())
	} else {
		output.WriteByte(0x0)
		value := ad.Value.(accessResultTag)
		output.WriteByte(byte(value))
	}

	return output.Bytes()
}
