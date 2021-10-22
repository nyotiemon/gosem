package axdr

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var validIntType = map[reflect.Kind]reflect.Kind{
	reflect.Int:    reflect.Int,
	reflect.Int8:   reflect.Int8,
	reflect.Int16:  reflect.Int16,
	reflect.Int32:  reflect.Int32,
	reflect.Int64:  reflect.Int64,
	reflect.Uint:   reflect.Uint,
	reflect.Uint8:  reflect.Uint8,
	reflect.Uint16: reflect.Uint16,
	reflect.Uint32: reflect.Uint32,
	reflect.Uint64: reflect.Uint64,
}

// Check if input is a valid u/int 8-64
func ValidNumberType(k interface{}) (reflect.Kind, error) {
	if what, ok := validIntType[reflect.TypeOf(k).Kind()]; ok {
		return what, nil
	}
	return reflect.Invalid, fmt.Errorf("%T is not a number", k)
}

// Encodes a number into correct byte according to A-XDR rule
// high-bit of first byte (x000 0000) will determine if first
// byte is L-of-length or not. If it is, then the rest bits are
// the value of L-of-length, else is the length itself.
// Sample (int -> hex). 3->03, 131->81 83, 25000->82 61 A8
func EncodeLength(dataLength interface{}) ([]byte, error) {
	var output []byte

	dataType, err := ValidNumberType(dataLength)
	if err != nil {
		return output, err
	}

	var trueLength []byte
	switch dataType {
	case reflect.Int:
		value := dataLength.(int)
		if value < 0 {
			return output, fmt.Errorf("%v value cannot be negative", value)
		}
		trueLength = make([]byte, 4)
		binary.BigEndian.PutUint32(trueLength, uint32(value))
	case reflect.Int64:
		value := dataLength.(int64)
		if value < 0 {
			return output, fmt.Errorf("%v value cannot be negative", value)
		}
		trueLength = make([]byte, 8)
		binary.BigEndian.PutUint64(trueLength, uint64(value))
	case reflect.Uint:
		value := dataLength.(uint)
		trueLength = make([]byte, 8)
		binary.BigEndian.PutUint64(trueLength, uint64(value))
	case reflect.Uint64:
		value := dataLength.(uint64)
		trueLength = make([]byte, 8)
		binary.BigEndian.PutUint64(trueLength, value)
	}

	for i, val := range trueLength {
		if val != 0x00 {
			trueLength = trueLength[i:]
			break
		}
		if i == len(trueLength)-1 && val == 0x00 {
			trueLength = trueLength[i:]
		}
	}

	if len(trueLength) == 1 {
		if trueLength[0] > 127 {
			output = append(output, byte(129))
		}
		output = append(output, trueLength...)
	} else {
		output = append(output, byte(128+len(trueLength)))
		output = append(output, trueLength...)
	}

	return output, nil
}

func EncodeArray(data []*DlmsData) ([]byte, error) {
	var output bytes.Buffer

	for _, d := range data {
		res, err := d.Encode()
		if err != nil {
			return []byte{}, nil
		}
		output.Write(res)
	}

	return output.Bytes(), nil
}

func EncodeStructure(data []*DlmsData) ([]byte, error) {
	res, err := EncodeArray(data)
	return res, err
}

func EncodeBoolean(data bool) ([]byte, error) {
	if data {
		return []byte{0xFF}, nil
	} else {
		return []byte{0x00}, nil
	}
}

// Encodes string of binary into byte. If length of bitstring
// is not multiplication of 8 bits, the left over will be put
// as trailing zeros.
// Sample: "111"-> E0
func EncodeBitString(data string) ([]byte, error) {
	var dataBytes bytes.Buffer
	var str string

	data = strings.ReplaceAll(data, " ", "")
	if len(strings.Trim(data, "01")) > 0 {
		return []byte{}, fmt.Errorf("data must be a string of binary, example: 11100000")
	}

	for i := 0; i < len(data); i += 8 {
		if i+8 > len(data) {
			str = data[i:]
			for len(str) < 8 {
				str += "0"
			}
		} else {
			str = data[i : i+8]
		}
		thisByte, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			return []byte{}, err
		}
		dataBytes.WriteByte(byte(thisByte))
	}

	return dataBytes.Bytes(), nil
}

func EncodeDoubleLong(data int32) ([]byte, error) {
	var output [4]byte
	binary.BigEndian.PutUint32(output[:], uint32(data))
	return output[:], nil
}

func EncodeDoubleLongUnsigned(data uint32) ([]byte, error) {
	var output [4]byte
	binary.BigEndian.PutUint32(output[:], data)
	return output[:], nil
}

// An ordered sequence of octets (8 bit bytes)
// Obis / Logical Name is en/decoded using TagOctetString
// other than that, input will be processed as hexstring
func EncodeOctetString(data string) ([]byte, error) {
	// Try to split with dots. if it can, and length is
	// exactly 6, then string is Obis code
	s := strings.Split(data, ".")
	if len(s) == 6 {
		bv := make([]byte, 6)
		for i, v := range s {
			bt, ok := strconv.ParseUint(v, 10, 8)
			if ok != nil {
				return nil, fmt.Errorf("failed to parse input as byte for Obis")
			}
			bv[i] = uint8(bt)
		}

		return bv, nil
	}

	return hex.DecodeString(data)
}

// An ordered sequence of ASCII characters
func EncodeVisibleString(data string) ([]byte, error) {
	for i := 0; i < len(data); i++ {
		if data[i] > '\u007F' { // taken from unicode.MaxASCII
			return []byte{}, fmt.Errorf("data to encode is not a valid ASCII string")
		}
	}
	return []byte(data), nil
}

// An ordered sequence of characters encoded as UTF-8
func EncodeUTF8String(data string) ([]byte, error) {
	if valid := utf8.ValidString(data); !valid {
		return []byte{}, fmt.Errorf("data to encode is not a valid UTF-8 string")
	}
	rs := []rune(data)

	byteSize := 0
	for _, r := range rs {
		byteSize += utf8.RuneLen(r)
	}
	output := make([]byte, byteSize)

	byteSize = 0
	for i := 0; i < len(rs); i++ {
		byteSize += utf8.EncodeRune(output[byteSize:], rs[i])
	}
	return output, nil
}

// binary coded decimal
func EncodeBCD(data int8) ([]byte, error) {
	output := make([]byte, 1)
	output[0] = byte(data)
	return output, nil
}

// Standard 8-4-2-1 decimal-only encoding
// note: this is not part of A-XDR encoder
func EncodeBCDs(data string) ([]byte, error) {
	if _, err := strconv.ParseInt(data, 10, 64); err != nil {
		return []byte{}, fmt.Errorf("data is non-encodable")
	}
	db := []byte(data)
	dl := (len(db) + 1) / 2
	output := make([]byte, dl)

	for i := 0; i < len(db); i++ {
		output[i/2] = (output[i/2] << 4) + db[i]&0xf
	}

	// shift 4 bits to the left if number of digit is not even
	// 12345 >> is 3 bytes, last byte for 5 should be 01010000
	if len(db)-(2*int(len(db)/2)) > 0 {
		output[dl-1] = (output[dl-1] << 4) + 0&0xf
	}

	return output, nil
}

func EncodeInteger(data int8) ([]byte, error) {
	output := make([]byte, 1)
	output[0] = byte(data)
	return output, nil
}

func EncodeLong(data int16) ([]byte, error) {
	var output [2]byte
	binary.BigEndian.PutUint16(output[:], uint16(data))
	return output[:], nil
}

func EncodeUnsigned(data uint8) ([]byte, error) {
	output := make([]byte, 1)
	output[0] = byte(data)
	return output, nil
}

func EncodeLongUnsigned(data uint16) ([]byte, error) {
	var output [2]byte
	binary.BigEndian.PutUint16(output[:], data)
	return output[:], nil
}

func EncodeLong64(data int64) ([]byte, error) {
	var output [8]byte
	binary.BigEndian.PutUint64(output[:], uint64(data))
	return output[:], nil
}

func EncodeLong64Unsigned(data uint64) ([]byte, error) {
	var output [8]byte
	binary.BigEndian.PutUint64(output[:], data)
	return output[:], nil
}

func EncodeEnum(data uint8) ([]byte, error) {
	// must be in 0..255 data range
	output, err := EncodeUnsigned(data)
	return output, err
}

func EncodeFloat32(data float32) ([]byte, error) {
	var output [4]byte
	binary.BigEndian.PutUint32(output[:], math.Float32bits(data))
	return output[:], nil
}

func EncodeFloat64(data float64) ([]byte, error) {
	var output [8]byte
	binary.BigEndian.PutUint64(output[:], math.Float64bits(data))
	return output[:], nil
}

// Encodes a date of time object into 5 bytes data
// year highbyte,
// year lowbyte,
// month,
// day of month,
// day of week
func EncodeDate(data time.Time) ([]byte, error) {
	output := make([]byte, 5)
	yb := make([]byte, 2)

	binary.BigEndian.PutUint16(yb, uint16(data.Year()))
	output[0] = yb[0]
	output[1] = yb[1]
	output[2] = byte(int(data.Month()))
	output[3] = byte(data.Day())
	output[4] = byte(int(data.Weekday()))

	return output, nil
}

// Encodes a time of time object into 4 bytes data
// hour,
// minute,
// second,
// hundredths
func EncodeTime(data time.Time) ([]byte, error) {
	output := make([]byte, 4)

	output[0] = byte(data.Hour())
	output[1] = byte(data.Minute())
	output[2] = byte(data.Second())
	output[3] = byte(data.Nanosecond())

	return output, nil
}

// Encodes datetime of time object into 12 bytes data
// year highbyte,
// year lowbyte,
// month,
// day of month,
// day of week,
// hour,
// minute,
// second,
// hundredths of second,
// deviation highbyte, -- interpreted as long in minutes of local time of UTC
// deviation lowbyte,
// clock status -- 0x00 means ok, 0xFF means not specified
func EncodeDateTime(data time.Time) ([]byte, error) {
	output := make([]byte, 12)

	yb := make([]byte, 2)
	binary.BigEndian.PutUint16(yb, uint16(data.Year()))

	output[0] = yb[0]
	output[1] = yb[1]
	output[2] = byte(int(data.Month()))
	output[3] = byte(data.Day())
	output[4] = byte(int(data.Weekday()))
	output[5] = byte(data.Hour())
	output[6] = byte(data.Minute())
	output[7] = byte(data.Second())
	output[8] = byte(data.Nanosecond())
	output[9] = 0x00
	output[10] = 0x00
	output[11] = 0x00

	return output, nil
}
