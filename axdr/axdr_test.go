package axdr

import (
	"bytes"
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
	if err == nil {
		t.Errorf("t13. Zero should have failed. val: %d, err:%v", t13, err)
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
	ts, err := EncodeOctetString("ABCD")
	res := bytes.Compare(ts, []byte{65, 66, 67, 68})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeOctetString("A1 -")
	res = bytes.Compare(ts, []byte{65, 49, 32, 45})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}

	ts, err = EncodeOctetString("A1 -")
	res = bytes.Compare(ts, []byte{65, 49, 32, 45})
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
	dt := time.Date(2020, time.January, 1, 10, 0, 0, 0, time.UTC)
	ts, err := EncodeTime(dt)
	res := bytes.Compare(ts, []byte{10, 0, 0, 255})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	dt = time.Date(2020, time.January, 1, 23, 59, 59, 0, time.UTC)
	ts, err = EncodeTime(dt)
	res = bytes.Compare(ts, []byte{23, 59, 59, 255})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestEncodeDateTime(t *testing.T) {
	dt := time.Date(20000, time.December, 30, 23, 59, 59, 0, time.UTC)
	ts, err := EncodeDateTime(dt)
	res := bytes.Compare(ts, []byte{78, 32, 12, 30, 6, 23, 59, 59, 255, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t1 failed. val: %d, err:%v", ts, err)
	}

	dt = time.Date(1500, time.January, 1, 0, 0, 0, 0, time.UTC)
	ts, err = EncodeDateTime(dt)
	res = bytes.Compare(ts, []byte{5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0})
	if res != 0 || err != nil {
		t.Errorf("t2 failed. val: %d, err:%v", ts, err)
	}
}

func TestDlmsData(t *testing.T) {
	var tDD DlmsData
	tDD = DlmsData{Tag: TagBoolean, Value: true}
	encodedBool := tDD.Encode()
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
	var tDD DlmsData
	tDD = DlmsData{Tag: TagBoolean, Value: nil}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should've panic on nil value")
		}
	}()
	tDD.Encode()
}

func TestDlmsData_WrongBoolValue(t *testing.T) {
	var tDD DlmsData
	tDD = DlmsData{Tag: TagBoolean, Value: 1234}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should've panic on wrong Value to Tag")
		}
	}()
	tDD.Encode()
}

func TestDlmsData_WrongBitStringValue(t *testing.T) {
	var tDD DlmsData
	tDD = DlmsData{Tag: TagBitString, Value: "ABCDEFG"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should've panic on value is not strings of binary")
		}
	}()
	tDD.Encode()
}

func TestDlmsData_DateTime(t *testing.T) {
	tDD := DlmsData{Tag: TagDateTime, Value: "9999-12-30 23:59:59"}
	encoded := tDD.Encode()
	res := bytes.Compare(encoded, []byte{byte(TagDateTime), 39, 15, 12, 30, 4, 23, 59, 59, 255, 0, 0, 0})
	if res != 0 {
		t.Errorf("DlmsData Encode DateTime get raw failed. val: %d", encoded)
	}

	dt := time.Date(20000, time.December, 30, 23, 59, 59, 0, time.UTC)
	tDD.Value = dt
	encoded = tDD.Encode()
	res = bytes.Compare(encoded, []byte{byte(TagDateTime), 78, 32, 12, 30, 6, 23, 59, 59, 255, 0, 0, 0})
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
	res := bytes.Compare(ts, []byte{byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0})
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
	res = bytes.Compare(d3.Raw(), []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0})
	if res != 0 {
		t.Errorf("t1.3 failed. val: %d", d3.Raw())
	}

	tables := []struct {
		x DlmsData
		y DlmsData
		z DlmsData
		r []byte
	}{
		{d1, d2, d3, []byte{byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0}},
		{d2, d1, d3, []byte{byte(TagBitString), 3, 224, byte(TagBoolean), 255, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0}},
		{d3, d2, d1, []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0, byte(TagBitString), 3, 224, byte(TagBoolean), 255}},
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
	encoded := tDD.Encode()
	res := bytes.Compare(encoded, []byte{byte(TagArray), 3, byte(TagBoolean), 255, byte(TagBitString), 3, 224, byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0})
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
	res = bytes.Compare(d3.Raw(), []byte{byte(TagDateTime), 7, 228, 3, 11, 3, 18, 0, 0, 255, 0, 0, 0})
	if res != 0 {
		t.Errorf("t1.3 failed. val: %d", d3.Raw())
	}

	tDD = DlmsData{Tag: TagArray, Value: []*DlmsData{&DlmsData{Tag: TagBoolean, Value: true}, &DlmsData{Tag: TagBoolean, Value: false}}}
	encoded = tDD.Encode()
	res = bytes.Compare(encoded, []byte{byte(TagArray), 2, byte(TagBoolean), 255, byte(TagBoolean), 0})
	if res != 0 {
		t.Errorf("t2 failed. val: %d", encoded)
	}
}

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
		bt, val, err := decodeLength(&table.src)
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
		// compare reminder bytes of src
		sameReminder := bytes.Compare(table.src, []byte{1, 2, 3})
		if sameReminder != 0 {
			t.Errorf("combination %v failed. Reminder get: %d, should:[1, 2, 3]", idx, table.src)
		}
	}

}
