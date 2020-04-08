package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
)

func TestNew_ActionResponseNormal(t *testing.T) {
	var ret GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)
	var ares ActResponse = *CreateActResponse(TagActSuccess, &ret)
	var a ActionResponseNormal = *CreateActionResponseNormal(81, ares)
	t1 := a.Encode()

	result := []byte{199, 1, 81, 0, 1, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_ActionResponseWithPBlock(t *testing.T) {
	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var a ActionResponseWithPBlock = *CreateActionResponseWithPBlock(81, dt)
	t1 := a.Encode()

	result := []byte{199, 2, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_ActionResponseWithList(t *testing.T) {
	// with 1 ActResponse
	var ret GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)
	var ares1 ActResponse = *CreateActResponse(TagActSuccess, &ret)
	var a ActionResponseWithList = *CreateActionResponseWithList(81, []ActResponse{ares1})
	t1 := a.Encode()

	result := []byte{199, 3, 81, 1, 0, 1, 0, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// with 2 ActResponse
	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var ret2 GetDataResult = *CreateGetDataResultAsData(dt)
	var ares2 ActResponse = *CreateActResponse(TagActSuccess, &ret2)
	var b ActionResponseWithList = *CreateActionResponseWithList(81, []ActResponse{ares1, ares2})
	t2 := b.Encode()

	result = []byte{199, 3, 81, 2, 0, 1, 0, 0, 0, 1, 1, 5, 0, 0, 0, 69}
	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t2, result)
	}
}

func TestNew_ActionResponseNextPBlock(t *testing.T) {
	var a ActionResponseNextPBlock = *CreateActionResponseNextPBlock(81, 1)
	t1 := a.Encode()
	result := []byte{199, 4, 81, 0, 0, 0, 1}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_ActionResponseNormal(t *testing.T) {
	src := []byte{199, 1, 81, 0, 1, 0, 0}
	a, err := DecodeActionResponseNormal(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionResponseNormal. err:%v", err)
	}

	var ret GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)
	var ares ActResponse = *CreateActResponse(TagActSuccess, &ret)
	var b ActionResponseNormal = *CreateActionResponseNormal(81, ares)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.Response.ReturnParam.IsData != b.Response.ReturnParam.IsData {
		t.Errorf("t1 Failed. Response.IsData get: %v, should:%v", a.Response.ReturnParam.IsData, b.Response.ReturnParam.IsData)
	}
	if a.Response.ReturnParam.Value != b.Response.ReturnParam.Value {
		t.Errorf("t1 Failed. Response.Value get: %v, should:%v", a.Response.ReturnParam.Value, b.Response.ReturnParam.Value)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_ActionResponseWithPBlock(t *testing.T) {
	src := []byte{199, 2, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	a, err := DecodeActionResponseWithPBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionResponseWithPBlock. err:%v", err)
	}

	var dt DataBlockSA = *CreateDataBlockSA(true, 1, []byte{1, 2, 3, 4, 5})
	var b ActionResponseWithPBlock = *CreateActionResponseWithPBlock(81, dt)

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}

	if a.PBlock.LastBlock != b.PBlock.LastBlock {
		t.Errorf("t1 Failed. PBlock.LastBlock get: %v, should:%v", a.PBlock.LastBlock, b.PBlock.LastBlock)
	}
	if a.PBlock.BlockNumber != b.PBlock.BlockNumber {
		t.Errorf("t1 Failed. PBlock.BlockNumber get: %v, should:%v", a.PBlock.BlockNumber, b.PBlock.BlockNumber)
	}
	res := bytes.Compare(a.PBlock.Raw, a.PBlock.Raw)
	if res != 0 {
		t.Errorf("t1 Failed. PBlock.Raw get: %v, should:%v", a.PBlock.Raw, a.PBlock.Raw)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_ActionResponseWithList(t *testing.T) {
	// ---------------------- with 1 ActResponse
	src := []byte{199, 3, 81, 1, 0, 1, 0, 0}
	a, err := DecodeActionResponseWithList(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionResponseWithList. err:%v", err)
	}

	var ret GetDataResult = *CreateGetDataResultAsResult(TagAccSuccess)
	var ares1 ActResponse = *CreateActResponse(TagActSuccess, &ret)
	var b ActionResponseWithList = *CreateActionResponseWithList(81, []ActResponse{ares1})

	if a.InvokePriority != b.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, b.InvokePriority)
	}
	if a.ResponseCount != b.ResponseCount {
		t.Errorf("t1 Failed. ResponseCount get: %v, should:%v", a.ResponseCount, b.ResponseCount)
	}

	if a.ResponseList[0].Result != b.ResponseList[0].Result {
		t.Errorf("t1 Failed. ResponseList[0].Result get: %v, should:%v", a.ResponseList[0].Result, b.ResponseList[0].Result)
	}

	aData1 := a.ResponseList[0].ReturnParam.ValueAsAccess()
	bData1 := b.ResponseList[0].ReturnParam.ValueAsAccess()
	if aData1 != bData1 {
		t.Errorf("t1 Failed. ResponseList[0].ReturnParam.Value get: %v, should:%v", aData1, bData1)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}

	// ---------------------- with 2 ActResponse
	src = []byte{199, 3, 81, 2, 0, 1, 0, 0, 0, 1, 1, 5, 0, 0, 0, 69}
	a, err = DecodeActionResponseWithList(&src)
	if err != nil {
		t.Errorf("t2 Failed to DecodeActionResponseWithList. err:%v", err)
	}

	var dt DlmsData = *CreateAxdrDoubleLong(69)
	var ret2 GetDataResult = *CreateGetDataResultAsData(dt)
	var ares2 ActResponse = *CreateActResponse(TagActSuccess, &ret2)
	b = *CreateActionResponseWithList(81, []ActResponse{ares1, ares2})

	if a.ResponseCount != b.ResponseCount {
		t.Errorf("t1 Failed. ResponseCount get: %v, should:%v", a.ResponseCount, b.ResponseCount)
	}

	if a.ResponseList[1].Result != b.ResponseList[1].Result {
		t.Errorf("t1 Failed. ResponseList[1].Result get: %v, should:%v", a.ResponseList[1].Result, b.ResponseList[1].Result)
	}
	if a.ResponseList[1].ReturnParam.IsData != b.ResponseList[1].ReturnParam.IsData {
		t.Errorf("t1 Failed. ResponseList[1].ReturnParam.IsData get: %v, should:%v", a.ResponseList[1].ReturnParam.IsData, b.ResponseList[1].ReturnParam.IsData)
	}

	aData2 := a.ResponseList[1].ReturnParam.ValueAsData().Value
	bData2 := b.ResponseList[1].ReturnParam.ValueAsData().Value
	if aData2 != bData2 {
		t.Errorf("t1 Failed. ResponseList[1].ReturnParam.Value get: %v, should:%v", aData2, bData2)
	}

	if len(src) > 0 {
		t.Errorf("t2 Failed. src should be empty. get: %v", src)
	}

}

func TestDecode_ActionResponseNextPBlock(t *testing.T) {
	var x ActionResponseNextPBlock = *CreateActionResponseNextPBlock(81, 1)
	src := []byte{199, 4, 81, 0, 0, 0, 1}

	a, err := DecodeActionResponseNextPBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeActionResponseNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.BlockNum != x.BlockNum {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.BlockNum, x.BlockNum)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}
