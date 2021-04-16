package dlms

import (
	"reflect"
	"testing"
)

func TestDecode_cosem(t *testing.T) {

	// ------------------  GetRequestNormal
	srcGetRequestNormal := []byte{192, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res, e := DecodeCosem(&srcGetRequestNormal)
	if e != nil {
		t.Errorf("Decode for GetRequestNormal Failed. err:%v", e)
	}
	_, assertTrue := res.(GetRequestNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetRequestNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  GetRequestNext
	srcGetRequestNext := []byte{192, 2, 81, 0, 0, 0, 2}
	res, e = DecodeCosem(&srcGetRequestNext)
	if e != nil {
		t.Errorf("Decode for GetRequestNext Failed. err:%v", e)
	}
	_, assertTrue = res.(GetRequestNext)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetRequestNext instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  GetRequestWithList
	srcGetRequestWithList := []byte{192, 3, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}
	res, e = DecodeCosem(&srcGetRequestWithList)
	if e != nil {
		t.Errorf("Decode for GetRequestWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(GetRequestWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetRequestWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  GetResponseNormal
	srcGetResponseNormal := []byte{196, 1, 81, 1, 5, 0, 0, 0, 69}
	res, e = DecodeCosem(&srcGetResponseNormal)
	if e != nil {
		t.Errorf("Decode for GetResponseNormal Failed. err:%v", e)
	}
	_, assertTrue = res.(GetResponseNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetResponseNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  GetResponseWithDataBlock
	srcGetResponseWithDataBlock := []byte{196, 2, 81, 1, 0, 0, 0, 1, 0, 12, 7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0}
	res, e = DecodeCosem(&srcGetResponseWithDataBlock)
	if e != nil {
		t.Errorf("Decode for GetResponseWithDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(GetResponseWithDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetResponseWithDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  GetResponseWithList
	srcGetResponseWithList := []byte{196, 3, 69, 2, 0, 0, 1, 5, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcGetResponseWithList)
	if e != nil {
		t.Errorf("Decode for GetResponseWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(GetResponseWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return GetResponseWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetRequestNormal
	srcSetRequestNormal := []byte{193, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 9, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcSetRequestNormal)
	if e != nil {
		t.Errorf("Decode for SetRequestNormal Failed. err:%v", e)
	}
	_, assertTrue = res.(SetRequestNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetRequestNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetRequestWithFirstDataBlock
	srcSetRequestWithFirstDataBlock := []byte{193, 2, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcSetRequestWithFirstDataBlock)
	if e != nil {
		t.Errorf("Decode for SetRequestWithFirstDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(SetRequestWithFirstDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetRequestWithFirstDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetRequestWithDataBlock
	srcSetRequestWithDataBlock := []byte{193, 3, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcSetRequestWithDataBlock)
	if e != nil {
		t.Errorf("Decode for SetRequestWithDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(SetRequestWithDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetRequestWithDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetRequestWithList
	srcSetRequestWithList := []byte{193, 4, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 9, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcSetRequestWithList)
	if e != nil {
		t.Errorf("Decode for SetRequestWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(SetRequestWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetRequestWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetRequestWithListAndFirstDataBlock
	srcSetRequestWithListAndFirstDataBlock := []byte{193, 5, 69, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcSetRequestWithListAndFirstDataBlock)
	if e != nil {
		t.Errorf("Decode for SetRequestWithListAndFirstDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(SetRequestWithListAndFirstDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetRequestWithListAndFirstDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetResponseNormal
	srcSetResponseNormal := []byte{197, 1, 81, 0}
	res, e = DecodeCosem(&srcSetResponseNormal)
	if e != nil {
		t.Errorf("Decode for SetResponseNormal Failed. err:%v", e)
	}
	_, assertTrue = res.(SetResponseNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetResponseNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetResponseDataBlock
	srcSetResponseDataBlock := []byte{197, 2, 81, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcSetResponseDataBlock)
	if e != nil {
		t.Errorf("Decode for SetResponseDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(SetResponseDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetResponseDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetResponseLastDataBlock
	srcSetResponseLastDataBlock := []byte{197, 3, 81, 0, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcSetResponseLastDataBlock)
	if e != nil {
		t.Errorf("Decode for SetResponseLastDataBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(SetResponseLastDataBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetResponseLastDataBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetResponseLastDataBlockWithList
	srcSetResponseLastDataBlockWithList := []byte{197, 4, 81, 3, 0, 1, 250, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcSetResponseLastDataBlockWithList)
	if e != nil {
		t.Errorf("Decode for SetResponseLastDataBlockWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(SetResponseLastDataBlockWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetResponseLastDataBlockWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  SetResponseWithList
	srcSetResponseWithList := []byte{197, 5, 81, 3, 0, 1, 250}
	res, e = DecodeCosem(&srcSetResponseWithList)
	if e != nil {
		t.Errorf("Decode for SetResponseWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(SetResponseWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return SetResponseWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestNormal
	srcActionRequestNormal := []byte{195, 1, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionRequestNormal)
	if e != nil {
		t.Errorf("Decode for ActionRequestNormal Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestNextPBlock
	srcActionRequestNextPBlock := []byte{195, 2, 81, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcActionRequestNextPBlock)
	if e != nil {
		t.Errorf("Decode for ActionRequestNextPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestNextPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestNextPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestWithList
	srcActionRequestWithList := []byte{195, 3, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 9, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionRequestWithList)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestWithFirstPBlock
	srcActionRequestWithFirstPBlock := []byte{195, 4, 81, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionRequestWithFirstPBlock)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithFirstPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestWithFirstPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestWithFirstPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestWithListAndFirstPBlock
	srcActionRequestWithListAndFirstPBlock := []byte{195, 5, 81, 1, 0, 1, 1, 0, 0, 3, 0, 255, 2, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionRequestWithListAndFirstPBlock)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithListAndFirstPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestWithListAndFirstPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestWithListAndFirstPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionRequestWithPBlock
	srcActionRequestWithPBlock := []byte{195, 6, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionRequestWithPBlock)
	if e != nil {
		t.Errorf("Decode for ActionRequestWithPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionRequestWithPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionRequestWithPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionResponseNormal
	srcActionResponseNormal := []byte{199, 1, 81, 0, 1, 0, 0}
	res, e = DecodeCosem(&srcActionResponseNormal)
	if e != nil {
		t.Errorf("Decode for ActionResponseNormal Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionResponseNormal)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionResponseNormal instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionResponseWithPBlock
	srcActionResponseWithPBlock := []byte{199, 2, 81, 1, 0, 0, 0, 1, 5, 1, 2, 3, 4, 5}
	res, e = DecodeCosem(&srcActionResponseWithPBlock)
	if e != nil {
		t.Errorf("Decode for ActionResponseWithPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionResponseWithPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionResponseWithPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionResponseWithList
	srcActionResponseWithList := []byte{199, 3, 81, 1, 0, 1, 0, 0}
	res, e = DecodeCosem(&srcActionResponseWithList)
	if e != nil {
		t.Errorf("Decode for ActionResponseWithList Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionResponseWithList)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionResponseWithList instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  ActionResponseNextPBlock
	srcActionResponseNextPBlock := []byte{199, 4, 81, 0, 0, 0, 1}
	res, e = DecodeCosem(&srcActionResponseNextPBlock)
	if e != nil {
		t.Errorf("Decode for ActionResponseNextPBlock Failed. err:%v", e)
	}
	_, assertTrue = res.(ActionResponseNextPBlock)
	if !assertTrue {
		t.Errorf("Decode supposed to return ActionResponseNextPBlock instead of %v", reflect.TypeOf(res).Name())
	}

	// ------------------  Error test
	srcError := []byte{255, 255, 255}
	_, wow := DecodeCosem(&srcError)
	if wow == nil {
		t.Errorf("Decode should've return error.")
	}
}
