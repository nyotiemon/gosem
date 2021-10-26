package axdr

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEncodeLength(t *testing.T) {
	t1, err := EncodeLength(125)
	res := bytes.Compare(t1, []byte{125})
	if res != 0 {
		t.Errorf("t1 failed. val: %d, err:%v", t1, err)
	}

	t2, err := EncodeLength(128)
	res = bytes.Compare(t2, []byte{129, 128})
	if res != 0 {
		t.Errorf("t2 failed. val: %d, err:%v", t2, err)
	}

	t3, err := EncodeLength(255)
	res = bytes.Compare(t3, []byte{129, 255})
	if res != 0 {
		t.Errorf("t3 failed. val: %d, err:%v", t3, err)
	}

	t4, err := EncodeLength(256)
	res = bytes.Compare(t4, []byte{130, 1, 0})
	if res != 0 {
		t.Errorf("t4 failed. val: %d, err:%v", t4, err)
	}

	t5, err := EncodeLength(65535)
	res = bytes.Compare(t5, []byte{130, 255, 255})
	if res != 0 {
		t.Errorf("t5 failed. val: %d, err:%v", t5, err)
	}

	t6, err := EncodeLength(65536)
	res = bytes.Compare(t6, []byte{131, 1, 0, 0})
	if res != 0 {
		t.Errorf("t6 failed. val: %d, err:%v", t6, err)
	}

	t7, err := EncodeLength(uint(18446744073709551615))
	res = bytes.Compare(t7, []byte{136, 255, 255, 255, 255, 255, 255, 255, 255})
	if res != 0 {
		t.Errorf("t7 failed. val: %d, err:%v", t7, err)
	}

	t8, err := EncodeLength(uint64(18446744073709551615))
	res = bytes.Compare(t8, []byte{136, 255, 255, 255, 255, 255, 255, 255, 255})
	if res != 0 {
		t.Errorf("t8 failed. val: %d, err:%v", t8, err)
	}

	t9, err := EncodeLength("123")
	if err == nil {
		t.Errorf("t9. String should have failed. val: %d, err:%v", t9, err)
	}

	t10, err := EncodeLength(3.14)
	if err == nil {
		t.Errorf("t10. Float should have failed. val: %d, err:%v", t10, err)
	}

	t11, err := EncodeLength(-500)
	if err == nil {
		t.Errorf("t11. Negative number should have failed. val: %d, err:%v", t11, err)
	}

	t12, err := EncodeLength(int64(-500000000))
	if err == nil {
		t.Errorf("t12. Negative int64 number should have failed. val: %d, err:%v", t12, err)
	}

	t13, err := EncodeLength(0)
	res = bytes.Compare(t13, []byte{0})
	if res != 0 || err != nil {
		t.Errorf("t13 failed. val: %d, err:%v", t13, err)
	}
}

func TestEncodeBoolean(t *testing.T) {
	ts, err := EncodeBoolean(true)
	res := bytes.Compare(ts, []byte{255})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBoolean(false)
	res = bytes.Compare(ts, []byte{0})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeBitString(t *testing.T) {
	ts, err := EncodeBitString("11111000")
	res := bytes.Compare(ts, []byte{248})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBitString("111100000001")
	res = bytes.Compare(ts, []byte{240, 16})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBitString("0000111111110000111111110000000101010101")
	res = bytes.Compare(ts, []byte{15, 240, 255, 1, 85})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBitString("00001111 11110000 11111111 00000001 01010101 1")
	res = bytes.Compare(ts, []byte{15, 240, 255, 1, 85, 128})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeDoubleLong(t *testing.T) {
	ts, err := EncodeDoubleLong(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLong(255)
	res = bytes.Compare(ts, []byte{0, 0, 0, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLong(-25)
	res = bytes.Compare(ts, []byte{255, 255, 255, 231})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLong(65535)
	res = bytes.Compare(ts, []byte{0, 0, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLong(2147483647)
	res = bytes.Compare(ts, []byte{127, 255, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLong(-2147483647)
	res = bytes.Compare(ts, []byte{128, 0, 0, 1})
	if res != 0 || err != nil {
		t.Errorf("t5 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeDoubleLongUnsigned(t *testing.T) {
	ts, err := EncodeDoubleLongUnsigned(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLongUnsigned(255)
	res = bytes.Compare(ts, []byte{0, 0, 0, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLongUnsigned(65535)
	res = bytes.Compare(ts, []byte{0, 0, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeDoubleLongUnsigned(4294967295)
	res = bytes.Compare(ts, []byte{255, 255, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeOctetString(t *testing.T) {
	ts, err := EncodeOctetString("07D20C04030A060BFF007800")
	res := bytes.Compare(ts, []byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeOctetString("1.0.0.3.0.255")
	res = bytes.Compare(ts, []byte{1, 0, 0, 3, 0, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeVisibleString(t *testing.T) {
	ts, err := EncodeVisibleString("ABCD")
	res := bytes.Compare(ts, []byte{65, 66, 67, 68})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeVisibleString("a1 -")
	res = bytes.Compare(ts, []byte{97, 49, 32, 45})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeVisibleString("{}[]()!;")
	res = bytes.Compare(ts, []byte{123, 125, 91, 93, 40, 41, 33, 59})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeVisibleString("ÆÁÉÍÓÚ")
	if err == nil {
		t.Errorf("t4 should've error on non-ascii value. val: %d, err:%v", ts, err)
	}
}

func TestEncodeUTF8String(t *testing.T) {
	ts, err := EncodeUTF8String("ABCD")
	res := bytes.Compare(ts, []byte{65, 66, 67, 68})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeUTF8String("aфᐃ𝕫")
	res = bytes.Compare(ts, []byte{97, 209, 132, 225, 144, 131, 240, 157, 149, 171})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeUTF8String("我愛你")
	res = bytes.Compare(ts, []byte{230, 136, 145, 230, 132, 155, 228, 189, 160})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeBCDAndBCDs(t *testing.T) {
	ts, err := EncodeBCD(int8(127))
	res := bytes.Compare(ts, []byte{127})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBCD(int8(-1))
	res = bytes.Compare(ts, []byte{255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBCDs("1234")
	res = bytes.Compare(ts, []byte{18, 52})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeBCDs("12345")
	res = bytes.Compare(ts, []byte{18, 52, 80})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeInteger(t *testing.T) {
	ts, err := EncodeInteger(-128)
	res := bytes.Compare(ts, []byte{128})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeInteger(0)
	res = bytes.Compare(ts, []byte{0})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeInteger(127)
	res = bytes.Compare(ts, []byte{127})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeInteger(-1)
	res = bytes.Compare(ts, []byte{255})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeLong(t *testing.T) {
	ts, err := EncodeLong(0)
	res := bytes.Compare(ts, []byte{0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong(256)
	res = bytes.Compare(ts, []byte{1, 0})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong(1<<15 - 1)
	res = bytes.Compare(ts, []byte{127, 255})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong(-1 << 15)
	res = bytes.Compare(ts, []byte{128, 0})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeUnsigned(t *testing.T) {
	ts, err := EncodeUnsigned(0)
	res := bytes.Compare(ts, []byte{0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeUnsigned(255)
	res = bytes.Compare(ts, []byte{1<<8 - 1})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeLongUnsigned(t *testing.T) {
	ts, err := EncodeLongUnsigned(0)
	res := bytes.Compare(ts, []byte{0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLongUnsigned(1<<16 - 1)
	res = bytes.Compare(ts, []byte{255, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeLong64(t *testing.T) {
	ts, err := EncodeLong64(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong64(1<<63 - 1)
	res = bytes.Compare(ts, []byte{127, 255, 255, 255, 255, 255, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong64(-1 << 63)
	res = bytes.Compare(ts, []byte{128, 0, 0, 0, 0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong64(-1)
	res = bytes.Compare(ts, []byte{255, 255, 255, 255, 255, 255, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeLong64Unsigned(t *testing.T) {
	ts, err := EncodeLong64Unsigned(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeLong64Unsigned(1<<64 - 1)
	res = bytes.Compare(ts, []byte{255, 255, 255, 255, 255, 255, 255, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeFloat32(t *testing.T) {
	ts, err := EncodeFloat32(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat32(float32(3.14))
	res = bytes.Compare(ts, []byte{64, 72, 245, 195})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat32(float32(-3.14))
	res = bytes.Compare(ts, []byte{192, 72, 245, 195})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat32(4294967295)
	res = bytes.Compare(ts, []byte{79, 128, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeFloat64(t *testing.T) {
	ts, err := EncodeFloat64(0)
	res := bytes.Compare(ts, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat64(float64(3.14))
	res = bytes.Compare(ts, []byte{64, 9, 30, 184, 81, 235, 133, 31})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat64(float64(-3.14))
	res = bytes.Compare(ts, []byte{192, 9, 30, 184, 81, 235, 133, 31})
	if res != 0 || err != nil {
		t.Errorf("t3 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat64(4294967295)
	res = bytes.Compare(ts, []byte{65, 239, 255, 255, 255, 224, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t4 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeFloat64(3.1415926535)
	res = bytes.Compare(ts, []byte{64, 9, 33, 251, 84, 65, 23, 68})
	if res != 0 || err != nil {
		t.Errorf("t5 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeDate(t *testing.T) {
	dt := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)
	ts, err := EncodeDate(dt)
	res := bytes.Compare(ts, []byte{7, 217, 11, 10, 2})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	dt = time.Date(1500, time.January, 1, 0, 0, 0, 0, time.UTC)
	ts, err = EncodeDate(dt)
	res = bytes.Compare(ts, []byte{5, 220, 1, 1, 1})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeTime(t *testing.T) {
	dt := time.Date(2020, time.January, 1, 10, 0, 0, 255, time.UTC)
	ts, err := EncodeTime(dt)
	res := bytes.Compare(ts, []byte{10, 0, 0, 255})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	dt = time.Date(2020, time.January, 1, 23, 59, 59, 255, time.UTC)
	ts, err = EncodeTime(dt)
	res = bytes.Compare(ts, []byte{23, 59, 59, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeDateTime(t *testing.T) {
	dt := time.Date(20000, time.December, 30, 23, 59, 59, 255, time.UTC)
	ts, err := EncodeDateTime(dt)
	res := bytes.Compare(ts, []byte{78, 32, 12, 30, 6, 23, 59, 59, 255, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	dt = time.Date(1500, time.January, 1, 0, 0, 0, 255, time.UTC)
	ts, err = EncodeDateTime(dt)
	res = bytes.Compare(ts, []byte{5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestDlmsData(t *testing.T) {
	tDD := DlmsData{Tag: TagBoolean, Value: true}
	encodedBool, err := tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData EncodeBoolean get error. %d", err)
	}
	res := bytes.Compare(encodedBool, []byte{byte(TagBoolean), 255})
	if res != 0 {
		t.Errorf("DlmsData EncodeBoolean get raw failed. val: %d", encodedBool)
	}
	res = bytes.Compare(tDD.RawValue(), []byte{255})
	if res != 0 {
		t.Errorf("DlmsData EncodeBoolean get rawValue failed. val: %d", tDD.RawValue())
	}
	tDD.Tag = TagBitString
	tDD.Value = "0000111111110000111111110000000101010101"
	tDD.Encode()
	res = bytes.Compare(tDD.Raw(), []byte{byte(TagBitString), 40, 15, 240, 255, 1, 85})
	if res != 0 {
		t.Errorf("t3 failed. val: %d", tDD.Raw())
	}
}

func TestDlmsData_NilValue(t *testing.T) {
	tDD := DlmsData{Tag: TagBoolean, Value: nil}

	_, err := tDD.Encode()
	if err == nil {
		t.Errorf("Should've panic on nil value")
	}
}

func TestDlmsData_WrongBoolValue(t *testing.T) {
	tDD := DlmsData{Tag: TagBoolean, Value: 1234}

	_, err := tDD.Encode()
	if err == nil {
		t.Errorf("Should've panic on wrong Value to Tag")
	}
}

func TestDlmsData_WrongBitStringValue(t *testing.T) {
	tDD := DlmsData{Tag: TagBitString, Value: "ABCDEFG"}

	_, err := tDD.Encode()
	if err == nil {
		t.Errorf("Should've panic on value is not strings of binary")
	}
}

func TestDlmsData_DateTime(t *testing.T) {
	tDD := DlmsData{Tag: TagDateTime, Value: "9999-12-30 23:59:59"}
	encoded, err := tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData Encode DateTime get error. %d", err)
	}
	res := bytes.Compare(encoded, []byte{byte(TagDateTime), 39, 15, 12, 30, 4, 23, 59, 59, 0, 0, 0, 0})
	if res != 0 {
		t.Errorf("DlmsData Encode DateTime get raw failed. val: %d", encoded)
	}

	dt := time.Date(20000, time.December, 30, 23, 59, 59, 0, time.UTC)
	tDD.Value = dt
	encoded, err = tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData Encode DateTime get error. %d", err)
	}
	res = bytes.Compare(encoded, []byte{byte(TagDateTime), 78, 32, 12, 30, 6, 23, 59, 59, 0, 0, 0, 0})
	if res != 0 {
		t.Errorf("DlmsData Encode DateTime get raw failed. val: %d", encoded)
	}
}

func TestArray(t *testing.T) {
	d1 := DlmsData{Tag: TagBoolean, Value: true}
	d2 := DlmsData{Tag: TagBitString, Value: "111"}
	d3 := DlmsData{Tag: TagDateTime, Value: "2020-03-11 18:00:00"}

	ls := []*DlmsData{&d1, &d2, &d3}
	ts, err := EncodeArray(ls)
	res := bytes.Compare(ts, []byte{byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	res = bytes.Compare(d1.Raw(), []byte{byte(TagBoolean), 255})
	if res != 0 {
		t.Errorf("t1.1 failed. val: %d", d1.Raw())
	}
	res = bytes.Compare(d2.Raw(), []byte{byte(TagBitString), 3, 224})
	if res != 0 {
		t.Errorf("t1.2 failed. val: %d", d2.Raw())
	}
	res = bytes.Compare(d3.Raw(), []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0})
	if res != 0 {
		t.Errorf("t1.3 failed. val: %d", d3.Raw())
	}

	tables := []struct {
		x DlmsData
		y DlmsData
		z DlmsData
		r []byte
	}{
		{d1, d2, d3, []byte{byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0}},
		{d2, d1, d3, []byte{byte(TagBitString), 3, 224, byte(TagBoolean), 255, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0}},
		{d3, d2, d1, []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0, byte(TagBitString), 3, 224, byte(TagBoolean), 255}},
	}
	for idx, table := range tables {
		ts, err = EncodeArray([]*DlmsData{&table.x, &table.y, &table.z})
		res = bytes.Compare(ts, table.r)
		if res != 0 || err != nil {
			t.Errorf("combination %v failed. get: %d, should:%v", idx, ts, table.r)
		}
	}
}

func TestDlmsData_Array(t *testing.T) {
	d1 := DlmsData{Tag: TagBoolean, Value: true}
	d2 := DlmsData{Tag: TagBitString, Value: "111"}
	d3 := DlmsData{Tag: TagDateTime, Value: "2020-03-11 18:00:00"}
	tDD := DlmsData{Tag: TagArray, Value: []*DlmsData{&d1, &d2, &d3}}
	encoded, err := tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData Encode Array get error. %d", err)
	}
	res := bytes.Compare(encoded, []byte{byte(TagArray), 3, byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0})
	if res != 0 {
		t.Errorf("t1 failed. val: %d", encoded)
	}

	res = bytes.Compare(d1.Raw(), []byte{byte(TagBoolean), 255})
	if res != 0 {
		t.Errorf("t1.1 failed. val: %d", d1.Raw())
	}
	res = bytes.Compare(d2.Raw(), []byte{byte(TagBitString), 3, 224})
	if res != 0 {
		t.Errorf("t1.2 failed. val: %d", d2.Raw())
	}
	res = bytes.Compare(d3.Raw(), []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 0, 0, 0, 0})
	if res != 0 {
		t.Errorf("t1.3 failed. val: %d", d3.Raw())
	}

	tDD = DlmsData{Tag: TagArray, Value: []*DlmsData{{Tag: TagBoolean, Value: true}, {Tag: TagBoolean, Value: false}}}
	encoded, err = tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData Encode Array get error. %d", err)
	}
	res = bytes.Compare(encoded, []byte{byte(TagArray), 2, byte(TagBoolean), 255, byte(TagBoolean), 0})
	if res != 0 {
		t.Errorf("t2 failed. val: %d", encoded)
	}

	tDD = DlmsData{Tag: TagArray, Value: []*DlmsData{}}
	encoded, err = tDD.Encode()
	if err != nil {
		t.Errorf("DlmsData Encode Array get error. %d", err)
	}
	res = bytes.Compare(encoded, []byte{byte(TagArray), 0})
	if res != 0 {
		t.Errorf("t3 failed. val: %d", encoded)
	}
}

// ---------- decoding tests

func TestDecodeLength(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val uint64
	}{
		{[]byte{2, 1, 2, 3}, []byte{2}, 2},
		{[]byte{131, 1, 0, 0, 1, 2, 3}, []byte{131, 1, 0, 0}, 65536},
		{[]byte{136, 255, 255, 255, 255, 255, 255, 255, 255, 1, 2, 3}, []byte{136, 255, 255, 255, 255, 255, 255, 255, 255}, 18446744073709551615},
	}
	for idx, table := range tables {
		bt, val, err := DecodeLength(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %d, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %d, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %d, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeBoolean(t *testing.T) {
	src := []byte{255, 1, 2, 3}
	bt, val, err := DecodeBoolean(&src)
	if err != nil {
		t.Errorf("t1 failed. got an error:%v", err)
	}
	sameByte := bytes.Compare(bt, []byte{255})
	if sameByte != 0 {
		t.Errorf("t1 failed. val: %d", sameByte)
	}
	sameValue := (val == true)
	if !sameValue {
		t.Errorf("t1 failed. Value get: %v", val)
	}
	sameReminder := bytes.Compare(src, []byte{1, 2, 3})
	if sameReminder != 0 {
		t.Errorf("t1 failed. Reminder get: %d, should:[1, 2, 3]", src)
	}
}

func TestDecodeBitString(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val string
	}{
		{[]byte{248, 1, 2, 3}, []byte{248}, "11111000"},
		{[]byte{15, 240, 255, 1, 85, 1, 2, 3}, []byte{15, 240, 255, 1, 85}, "0000111111110000111111110000000101010101"},
		{[]byte{15, 240, 255, 1, 85, 128, 1, 2, 3}, []byte{15, 240, 255, 1, 85, 128}, "00001111111100001111111100000001010101011"},
	}
	for idx, table := range tables {
		bt, val, err := DecodeBitString(&table.src, uint64(len(table.val)))
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %s, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeDoubleLong(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val int32
	}{
		{[]byte{255, 255, 255, 231, 1, 2, 3}, []byte{255, 255, 255, 231}, -25},
		{[]byte{127, 255, 255, 255, 1, 2, 3}, []byte{127, 255, 255, 255}, 2147483647},
		{[]byte{128, 0, 0, 1, 1, 2, 3}, []byte{128, 0, 0, 1}, -2147483647},
	}
	for idx, table := range tables {
		bt, val, err := DecodeDoubleLong(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeDoubleLongUnsigned(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val uint32
	}{
		{[]byte{0, 0, 0, 255, 1, 2, 3}, []byte{0, 0, 0, 255}, 255},
		{[]byte{0, 0, 255, 255, 1, 2, 3}, []byte{0, 0, 255, 255}, 65535},
		{[]byte{255, 255, 255, 255, 1, 2, 3}, []byte{255, 255, 255, 255}, 4294967295},
	}
	for idx, table := range tables {
		bt, val, err := DecodeDoubleLongUnsigned(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeOctetString(t *testing.T) {
	tables := []struct {
		src []byte
		lt  uint64
		val string
	}{
		{[]byte{7, 210, 12, 4, 3, 10, 6, 11, 255, 0, 120, 0, 1, 2, 3}, 12, "07D20C04030A060BFF007800"},
		{[]byte{1, 0, 0, 3, 0, 255, 1, 2, 3}, 6, "0100000300FF"},
	}
	for idx, table := range tables {
		answer := table.src[:table.lt]
		bt, val, err := DecodeOctetString(&table.src, table.lt)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(answer, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, answer)
		}
		// compare length value
		sameValue := (table.val == strings.ToUpper(val))
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %s, should:%v", idx, strings.ToUpper(val), table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeVisibleString(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val string
	}{
		{[]byte{65, 66, 67, 68, 1, 2, 3}, []byte{65, 66, 67, 68}, "ABCD"},
		{[]byte{123, 125, 91, 93, 40, 41, 33, 59, 1, 2, 3}, []byte{123, 125, 91, 93, 40, 41, 33, 59}, "{}[]()!;"},
	}
	for idx, table := range tables {
		bt, val, err := DecodeVisibleString(&table.src, uint64(len(table.val)))
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %s, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeUTF8String(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val string
	}{
		{[]byte{65, 66, 67, 68, 1, 2, 3}, []byte{65, 66, 67, 68}, "ABCD"},
		{[]byte{97, 209, 132, 225, 144, 131, 240, 157, 149, 171, 1, 2, 3}, []byte{97, 209, 132, 225, 144, 131, 240, 157, 149, 171}, "aфᐃ𝕫"},
		{[]byte{230, 136, 145, 230, 132, 155, 228, 189, 160, 1, 2, 3}, []byte{230, 136, 145, 230, 132, 155, 228, 189, 160}, "我愛你"},
	}
	for idx, table := range tables {
		bt, val, err := DecodeUTF8String(&table.src, uint64(len(table.val)))
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %s, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeBCD(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val int8
	}{
		{[]byte{127, 1, 2, 3}, []byte{127}, 127},
		{[]byte{255, 1, 2, 3}, []byte{255}, -1},
	}
	for idx, table := range tables {
		bt, val, err := DecodeBCD(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

// DecodeInteger == DecodeBCD == DecodeEnum

func TestDecodeLong(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val int16
	}{
		{[]byte{127, 255, 1, 2, 3}, []byte{127, 255}, 1<<15 - 1},
		{[]byte{128, 0, 1, 2, 3}, []byte{128, 0}, -1 << 15},
	}
	for idx, table := range tables {
		bt, val, err := DecodeLong(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeUnsigned(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val uint8
	}{
		{[]byte{255, 1, 2, 3}, []byte{255}, 255},
		{[]byte{0, 1, 2, 3}, []byte{0}, 0},
	}
	for idx, table := range tables {
		bt, val, err := DecodeUnsigned(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeLongUnsigned(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val uint16
	}{
		{[]byte{255, 255, 1, 2, 3}, []byte{255, 255}, 1<<16 - 1},
		{[]byte{0, 0, 1, 2, 3}, []byte{0, 0}, 0},
	}
	for idx, table := range tables {
		bt, val, err := DecodeLongUnsigned(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeLong64(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val int64
	}{
		{[]byte{127, 255, 255, 255, 255, 255, 255, 255, 1, 2, 3}, []byte{127, 255, 255, 255, 255, 255, 255, 255}, 1<<63 - 1},
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}, []byte{128, 0, 0, 0, 0, 0, 0, 0}, -1 << 63},
		{[]byte{255, 255, 255, 255, 255, 255, 255, 255, 1, 2, 3}, []byte{255, 255, 255, 255, 255, 255, 255, 255}, -1},
	}
	for idx, table := range tables {
		bt, val, err := DecodeLong64(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeLong64Unsigned(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val uint64
	}{
		{[]byte{255, 255, 255, 255, 255, 255, 255, 255, 1, 2, 3}, []byte{255, 255, 255, 255, 255, 255, 255, 255}, 1<<64 - 1},
		{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}, []byte{0, 0, 0, 0, 0, 0, 0, 0}, 0},
	}
	for idx, table := range tables {
		bt, val, err := DecodeLong64Unsigned(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeFloat32(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val float32
	}{
		{[]byte{64, 72, 245, 195, 1, 2, 3}, []byte{64, 72, 245, 195}, 3.14},
		{[]byte{79, 128, 0, 0, 1, 2, 3}, []byte{79, 128, 0, 0}, 4294967295},
		{[]byte{192, 72, 245, 195, 1, 2, 3}, []byte{192, 72, 245, 195}, -3.14},
	}
	for idx, table := range tables {
		bt, val, err := DecodeFloat32(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeFloat64(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val float64
	}{
		{[]byte{64, 9, 30, 184, 81, 235, 133, 31, 1, 2, 3}, []byte{64, 9, 30, 184, 81, 235, 133, 31}, 3.14},
		{[]byte{64, 9, 33, 251, 84, 65, 23, 68, 1, 2, 3}, []byte{64, 9, 33, 251, 84, 65, 23, 68}, 3.1415926535},
		{[]byte{65, 239, 255, 255, 255, 224, 0, 0, 1, 2, 3}, []byte{65, 239, 255, 255, 255, 224, 0, 0}, 4294967295},
	}
	for idx, table := range tables {
		bt, val, err := DecodeFloat64(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare length byte
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare length value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeDate(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val time.Time
	}{
		{[]byte{7, 217, 11, 10, 2, 1, 2, 3}, []byte{7, 217, 11, 10, 2}, time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)},
		{[]byte{5, 220, 1, 1, 1, 1, 2, 3}, []byte{5, 220, 1, 1, 1}, time.Date(1500, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	for idx, table := range tables {
		bt, val, err := DecodeDate(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare byte value
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare time value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeTime(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val time.Time
	}{
		{[]byte{10, 0, 0, 255, 1, 2, 3}, []byte{10, 0, 0, 255}, time.Date(0, time.January, 1, 10, 0, 0, 255, time.UTC)},
		{[]byte{23, 59, 59, 255, 1, 2, 3}, []byte{23, 59, 59, 255}, time.Date(0, time.January, 1, 23, 59, 59, 255, time.UTC)},
	}
	for idx, table := range tables {
		bt, val, err := DecodeTime(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare byte value
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare time value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecodeDateTime(t *testing.T) {
	tables := []struct {
		src []byte
		bt  []byte
		val time.Time
	}{
		{[]byte{78, 32, 12, 30, 6, 23, 59, 59, 255, 0, 0, 0, 1, 2, 3}, []byte{78, 32, 12, 30, 6, 23, 59, 59, 255, 0, 0, 0}, time.Date(20000, time.December, 30, 23, 59, 59, 255, time.UTC)},
		{[]byte{5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0, 1, 2, 3}, []byte{5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0}, time.Date(1500, time.January, 1, 0, 0, 0, 255, time.UTC)},
	}
	for idx, table := range tables {
		bt, val, err := DecodeDateTime(&table.src)
		if err != nil {
			t.Errorf("combination %v failed. got an error:%v", idx, err)
		}
		// compare byte value
		sameByte := bytes.Compare(table.bt, bt)
		if sameByte != 0 {
			t.Errorf("combination %v failed. Byte get: %v, should:%v", idx, bt, table.bt)
		}
		// compare time value
		sameValue := (table.val == val)
		if !sameValue {
			t.Errorf("combination %v failed. Value get: %v, should:%v", idx, val, table.val)
		}
		// compare remainder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %v, should:[1, 2, 3]", idx, table.src)
		}
	}
}

func TestDecoder1(t *testing.T) {
	d1 := DlmsData{Tag: TagLongUnsigned, Value: uint16(60226)}
	d2 := DlmsData{Tag: TagDateTime, Value: time.Date(2020, time.March, 16, 0, 0, 0, 255, time.UTC)}
	d3 := DlmsData{Tag: TagBitString, Value: "0"}
	d4 := DlmsData{Tag: TagDoubleLongUnsigned, Value: uint32(33426304)}
	d5 := DlmsData{Tag: TagLongUnsigned, Value: uint16(3105)}
	// da := DlmsData{Tag: TagArray, Value: DlmsData{Tag: TagStructure, Value: []DlmsData{d1, d2, d3, d4, d5}}}
	str := "0101020512EB421907E40310FF000000FF8000000401000601FE0B80120C21"

	src, _ := hex.DecodeString(str)

	dec := NewDataDecoder(&src)
	t1, err2 := dec.Decode(&src)
	if err2 != nil {
		t.Errorf("got an error when decoding:%v", err2)
	}
	if t1.Tag != TagArray {
		t.Errorf("First level should be TagArray, received: %v", reflect.TypeOf(t1.Tag).Kind())
	}

	t2 := t1.Value.([]*DlmsData)[0]
	if t2.Tag != TagStructure {
		t.Errorf("Second level should be TagStructure, received: %v", reflect.TypeOf(t2.Tag).Kind())
	}

	t3 := t2.Value.([]*DlmsData)
	if t3[0].Value != d1.Value {
		t.Errorf("should be same as d1 %v, received: %v", d1.Value, t3[0].Value)
	}
	if t3[1].Value != d2.Value {
		t.Errorf("should be same as d2 %v, received: %v", d2.Value, t3[1].Value)
	}
	if t3[2].Value != d3.Value {
		t.Errorf("should be same as d3 %v, received: %v", d3.Value, t3[2].Value)
	}
	if t3[3].Value != d4.Value {
		t.Errorf("should be same as d4 %v, received: %v", d4.Value, t3[3].Value)
	}
	if t3[4].Value != d5.Value {
		t.Errorf("should be same as d5 %v, received: %v", d5.Value, t3[4].Value)
	}

	src = []byte{1, 3, 4, 3, 224, 3, 255, 123, 123, 123}
	decoder := NewDataDecoder(&src)
	oriLength := len(src)
	_, err4 := decoder.Decode(&src)
	if err4 == nil {
		t.Errorf("t4 should be error")
	}
	if len(src) != oriLength {
		t.Errorf("src after error should still be the same length (%v)", src)
	}
}
