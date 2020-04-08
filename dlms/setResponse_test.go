package cosem

import (
	"bytes"
	"testing"
)

func TestNew_SetResponseNormal(t *testing.T) {

	var a SetResponseNormal = *CreateSetResponseNormal(81, TagAccSuccess)
	t1 := a.Encode()
	result := []byte{197, 1, 81, 0}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_SetResponseDataBlock(t *testing.T) {

	var a SetResponseDataBlock = *CreateSetResponseDataBlock(81, 1)
	t1 := a.Encode()
	result := []byte{197, 2, 81, 0, 0, 0, 1}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_SetResponseLastDataBlock(t *testing.T) {

	var a SetResponseLastDataBlock = *CreateSetResponseLastDataBlock(81, TagAccSuccess, 1)
	t1 := a.Encode()
	result := []byte{197, 3, 81, 0, 0, 0, 0, 1}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestNew_SetResponseLastDataBlockWithList(t *testing.T) {

	resList := []accessResultTag{TagAccSuccess, TagAccHardwareFault, TagAccOtherReason}
	var a SetResponseLastDataBlockWithList = *CreateSetResponseLastDataBlockWithList(81, resList, 1)
	t1 := a.Encode()
	result := []byte{197, 4, 81, 3, 0, 1, 250, 0, 0, 0, 1}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t2 should've panic on wrong Value")
		}
	}()
	b := *CreateSetResponseLastDataBlockWithList(81, []accessResultTag{}, 1)
	b.Encode()
}

func TestNew_SetResponseWithList(t *testing.T) {

	resList := []accessResultTag{TagAccSuccess, TagAccHardwareFault, TagAccOtherReason}
	var a SetResponseWithList = *CreateSetResponseWithList(81, resList)
	t1 := a.Encode()
	result := []byte{197, 5, 81, 3, 0, 1, 250}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("t2 should've panic on wrong Value")
		}
	}()
	b := *CreateSetResponseWithList(81, []accessResultTag{})
	b.Encode()
}

func TestDecode_SetResponseNormal(t *testing.T) {
	var x SetResponseNormal = *CreateSetResponseNormal(81, TagAccSuccess)
	src := []byte{197, 1, 81, 0}

	a, err := DecodeSetResponseNormal(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.Result != x.Result {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.Result, x.Result)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetResponseDataBlock(t *testing.T) {
	var x SetResponseDataBlock = *CreateSetResponseDataBlock(81, 1)
	src := []byte{197, 2, 81, 0, 0, 0, 1}

	a, err := DecodeSetResponseDataBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
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

func TestDecode_SetResponseLastDataBlock(t *testing.T) {
	var x SetResponseLastDataBlock = *CreateSetResponseLastDataBlock(81, TagAccSuccess, 1)
	src := []byte{197, 3, 81, 0, 0, 0, 0, 1}

	a, err := DecodeSetResponseLastDataBlock(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.Result != x.Result {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.Result, x.Result)
	}

	if a.BlockNum != x.BlockNum {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.BlockNum, x.BlockNum)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetResponseLastDataBlockWithList(t *testing.T) {
	resList := []accessResultTag{TagAccSuccess, TagAccHardwareFault, TagAccOtherReason}
	var x SetResponseLastDataBlockWithList = *CreateSetResponseLastDataBlockWithList(81, resList, 1)
	src := []byte{197, 4, 81, 3, 0, 1, 250, 0, 0, 0, 1}

	a, err := DecodeSetResponseLastDataBlockWithList(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.ResultCount != x.ResultCount {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.ResultCount, x.ResultCount)
	}

	if a.ResultList[0].Value() != resList[0].Value() {
		t.Errorf("t1 Failed. ResultList[0].Value get: %v, should:%v", a.ResultList[0].Value(), resList[0].Value())
	}

	if a.ResultList[1].Value() != resList[1].Value() {
		t.Errorf("t1 Failed. ResultList[1].Value get: %v, should:%v", a.ResultList[1].Value(), resList[1].Value())
	}

	if a.ResultList[2].Value() != resList[2].Value() {
		t.Errorf("t1 Failed. ResultList[2].Value get: %v, should:%v", a.ResultList[2].Value(), resList[2].Value())
	}

	if a.BlockNum != x.BlockNum {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.BlockNum, x.BlockNum)
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetResponseWithList(t *testing.T) {
	resList := []accessResultTag{TagAccSuccess, TagAccHardwareFault, TagAccOtherReason}
	var x SetResponseWithList = *CreateSetResponseWithList(81, resList)
	src := []byte{197, 5, 81, 3, 0, 1, 250}

	a, err := DecodeSetResponseWithList(&src)
	if err != nil {
		t.Errorf("t1 Failed to DecodeSetRequestNormal. err:%v", err)
	}

	if a.InvokePriority != x.InvokePriority {
		t.Errorf("t1 Failed. InvokePriority get: %v, should:%v", a.InvokePriority, x.InvokePriority)
	}

	if a.ResultCount != x.ResultCount {
		t.Errorf("t1 Failed. Result get: %v, should:%v", a.ResultCount, x.ResultCount)
	}

	if a.ResultList[0].Value() != resList[0].Value() {
		t.Errorf("t1 Failed. ResultList[0].Value get: %v, should:%v", a.ResultList[0].Value(), resList[0].Value())
	}

	if a.ResultList[1].Value() != resList[1].Value() {
		t.Errorf("t1 Failed. ResultList[1].Value get: %v, should:%v", a.ResultList[1].Value(), resList[1].Value())
	}

	if a.ResultList[2].Value() != resList[2].Value() {
		t.Errorf("t1 Failed. ResultList[2].Value get: %v, should:%v", a.ResultList[2].Value(), resList[2].Value())
	}

	if len(src) > 0 {
		t.Errorf("t1 Failed. src should be empty. get: %v", src)
	}
}

func TestDecode_SetResponse(t *testing.T) {
	var sr SetResponse

	// ------------------  SetResponseNormal
	src := []byte{197, 1, 81, 0}
	_, e := sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetResponseNormal Failed. err:%v", e)
	}

	// ------------------  SetResponseDataBlock
	src = []byte{197, 2, 81, 0, 0, 0, 1}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetResponseDataBlock Failed. err:%v", e)
	}

	// ------------------  SetResponseLastDataBlock
	src = []byte{197, 3, 81, 0, 0, 0, 0, 1}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetResponseLastDataBlock Failed. err:%v", e)
	}

	// ------------------  SetRequestWithList
	src = []byte{197, 4, 81, 3, 0, 1, 250, 0, 0, 0, 1}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetRequestWithList Failed. err:%v", e)
	}

	// ------------------  SetResponseLastDataBlockWithList
	src = []byte{197, 5, 81, 3, 0, 1, 250}
	_, e = sr.Decode(&src)
	if e != nil {
		t.Errorf("Decode for SetResponseLastDataBlockWithList Failed. err:%v", e)
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := sr.Decode(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
