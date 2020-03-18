package axdr

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

type Decoder struct {
	tag dataTag
}

var (
	ErrLengthLess = errors.New("not enough byte length provided")

	lengthAfterTag = map[dataTag]bool{
		TagNull:               false,
		TagArray:              true,
		TagStructure:          true,
		TagBoolean:            false,
		TagBitString:          true,
		TagDoubleLong:         false,
		TagDoubleLongUnsigned: false,
		TagFloatingPoint:      false,
		TagOctetString:        true,
		TagVisibleString:      true,
		TagUTF8String:         true,
		TagBCD:                false,
		TagInteger:            false,
		TagLong:               false,
		TagUnsigned:           false,
		TagLongUnsigned:       false,
		TagCompactArray:       false,
		TagLong64:             false,
		TagLong64Unsigned:     false,
		TagEnum:               false,
		TagFloat32:            false,
		TagFloat64:            false,
		TagDateTime:           false,
		TagDate:               false,
		TagTime:               false,
		TagDontCare:           false,
	}

	mapToDataTag = map[uint8]dataTag{
		0:   TagNull,
		1:   TagArray,
		2:   TagStructure,
		3:   TagBoolean,
		4:   TagBitString,
		5:   TagDoubleLong,
		6:   TagDoubleLongUnsigned,
		7:   TagFloatingPoint,
		9:   TagOctetString,
		10:  TagVisibleString,
		12:  TagUTF8String,
		13:  TagBCD,
		15:  TagInteger,
		16:  TagLong,
		17:  TagUnsigned,
		18:  TagLongUnsigned,
		19:  TagCompactArray,
		20:  TagLong64,
		21:  TagLong64Unsigned,
		22:  TagEnum,
		23:  TagFloat32,
		24:  TagFloat64,
		25:  TagDateTime,
		26:  TagDate,
		27:  TagTime,
		255: TagDontCare,
	}
)

func CheckTag(src *[]byte) (t dataTag, err error) {
	if len((*src)) < 1 {
		err = ErrLengthLess
		return
	}
	t, _ = mapToDataTag[uint8((*src)[0])]
	(*src) = (*src)[1:]

	return
}

func NewDecoder(dt dataTag) *Decoder {
	return &Decoder{tag: dt}
}

func (dec *Decoder) Decode(src *[]byte) (r DlmsData, err error) {
	r.Tag = dec.tag
	haveLength, _ := lengthAfterTag[dec.tag]
	var lengthByte []byte
	var lengthInt uint64 = 0
	if haveLength {
		lengthByte, lengthInt, err = decodeLength(src)
		if err != nil {
			return
		}
		r.rawLength = lengthByte
	}

	var rawValue []byte
	var value interface{}
	switch dec.tag {
	case TagNull:
		panic("Not yet implemented")

	case TagArray:
		output := make([]DlmsData, lengthInt)
		for i := 0; i < int(lengthInt); i++ {
			thisDataTag, _ := CheckTag(src)
			thisDecoder := NewDecoder(thisDataTag)
			thisDlmsData, thisError := thisDecoder.Decode(src)
			if thisError != nil {
				err = thisError
				return
			}
			output[i] = thisDlmsData
		}
		value = output

	case TagStructure:
		// same same as structure
		output := make([]DlmsData, lengthInt)
		for i := 0; i < int(lengthInt); i++ {
			thisDataTag, _ := CheckTag(src)
			thisDecoder := NewDecoder(thisDataTag)
			thisDlmsData, thisError := thisDecoder.Decode(src)
			if thisError != nil {
				err = thisError
				return
			}
			output[i] = thisDlmsData
		}
		value = output

	case TagBoolean:
		rawValue, value, err = decodeBoolean(src)
	case TagBitString:
		rawValue, value, err = decodeBitString(src, lengthInt)
	case TagDoubleLong:
		rawValue, value, err = decodeDoubleLong(src)
	case TagDoubleLongUnsigned:
		rawValue, value, err = decodeDoubleLongUnsigned(src)
	case TagFloatingPoint:
		rawValue, value, err = decodeFloat32(src)
	case TagOctetString:
		rawValue, value, err = decodeOctetString(src, lengthInt)
	case TagVisibleString:
		rawValue, value, err = decodeVisibleString(src, lengthInt)
	case TagUTF8String:
		rawValue, value, err = decodeUTF8String(src, lengthInt)
	case TagBCD:
		rawValue, value, err = decodeBCD(src)
	case TagInteger:
		rawValue, value, err = decodeInteger(src)
	case TagLong:
		rawValue, value, err = decodeLong(src)
	case TagUnsigned:
		rawValue, value, err = decodeUnsigned(src)
	case TagLongUnsigned:
		rawValue, value, err = decodeLongUnsigned(src)
	case TagCompactArray:
		panic("Not yet implemented")
	case TagLong64:
		rawValue, value, err = decodeLong64(src)
	case TagLong64Unsigned:
		rawValue, value, err = decodeLong64Unsigned(src)
	case TagEnum:
		rawValue, value, err = decodeEnum(src)
	case TagFloat32:
		rawValue, value, err = decodeFloat32(src)
	case TagFloat64:
		rawValue, value, err = decodeFloat64(src)
	case TagDateTime:
		rawValue, value, err = decodeDateTime(src)
	case TagDate:
		rawValue, value, err = decodeDate(src)
	case TagTime:
		rawValue, value, err = decodeTime(src)
	case TagDontCare:
		panic("Not yet implemented")
	}

	if err != nil {
		return
	}

	r.rawValue = rawValue
	r.Value = value
	r.raw.WriteByte(byte(dec.tag))
	if haveLength {
		r.raw.Write(lengthByte)
	}
	r.raw.Write(rawValue)

	return
}

func decodeLength(src *[]byte) (outByte []byte, outVal uint64, err error) {
	if (*src)[0] > byte(128) {
		lOfLength := int((*src)[0]) - 128     // L-of-length part
		realLength := (*src)[1 : lOfLength+1] // real length part
		if len(realLength) < lOfLength {
			err = ErrLengthLess
			return
		}
		outByte = (*src)[0 : lOfLength+1] // L-of-length and length

		buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}

		if len(realLength) > 8 {
			err = fmt.Errorf("Length value is bigger than uint64 max value. This decoder is limited to uint64")
		} else {
			bufStart := 7
			outStart := len(realLength) - 1
			for outStart >= 0 {
				buf[bufStart] = realLength[outStart]
				outStart--
				bufStart--
			}
		}
		outVal = binary.BigEndian.Uint64(buf[:])
		(*src) = (*src)[1+len(realLength):]

	} else {
		outByte = append(outByte, (*src)[0])
		outVal = uint64((*src)[0])
		(*src) = (*src)[1:]
	}

	return
}

func decodeBoolean(src *[]byte) (outByte []byte, outVal bool, err error) {
	if len(*src) < 1 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:1]
	if outByte[0] == 0xFF {
		outVal = true
	} else {
		outVal = false
	}
	(*src) = (*src)[1:]
	return
}

func decodeBitString(src *[]byte, length uint64) (outByte []byte, outVal string, err error) {
	byteLength := int(math.Ceil(float64(length) / 8))
	if len(*src) < byteLength {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:byteLength]

	var r strings.Builder
	for _, b := range outByte {
		r.WriteString(fmt.Sprintf("%08b", b))
	}
	outVal = (r.String())[:length]
	(*src) = (*src)[byteLength:]

	return
}

func decodeDoubleLong(src *[]byte) (outByte []byte, outVal int32, err error) {
	if len(*src) < 4 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:4]
	outVal |= int32(outByte[0]) << 24
	outVal |= int32(outByte[1]) << 16
	outVal |= int32(outByte[2]) << 8
	outVal |= int32(outByte[3])

	// buf := bytes.NewBuffer(outByte)
	// binary.Read(buf, binary.BigEndian, &outVal)
	(*src) = (*src)[4:]
	return
}

func decodeDoubleLongUnsigned(src *[]byte) (outByte []byte, outVal uint32, err error) {
	if len(*src) < 4 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:4]
	outVal |= uint32(outByte[0]) << 24
	outVal |= uint32(outByte[1]) << 16
	outVal |= uint32(outByte[2]) << 8
	outVal |= uint32(outByte[3])
	(*src) = (*src)[4:]
	return
}

func decodeOctetString(src *[]byte, length uint64) (outByte []byte, outVal string, err error) {
	if uint64(len(*src)) < length {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:length]
	outVal = string(outByte)
	(*src) = (*src)[length:]
	return
}

func decodeVisibleString(src *[]byte, length uint64) (outByte []byte, outVal string, err error) {
	if uint64(len(*src)) < length {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:length]
	outVal = string(outByte)
	(*src) = (*src)[length:]
	return
}

func decodeUTF8String(src *[]byte, length uint64) (outByte []byte, outVal string, err error) {
	if uint64(len(*src)) < length {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:length]

	var sb strings.Builder
	for sb.Len() < len(outByte) {
		r, _ := utf8.DecodeRune(outByte[sb.Len():])
		if r == utf8.RuneError {
			err = fmt.Errorf("Byte slice contain invalid UTF-8 runes")
			return
		}
		sb.WriteRune(r)
	}

	outVal = sb.String()
	(*src) = (*src)[length:]
	return
}

func decodeBCD(src *[]byte) (outByte []byte, outVal int8, err error) {
	outByte = (*src)[:1]
	outVal = int8(outByte[0])
	(*src) = (*src)[1:]
	return
}

func decodeInteger(src *[]byte) (outByte []byte, outVal int8, err error) {
	outByte = (*src)[:1]
	outVal = int8(outByte[0])
	(*src) = (*src)[1:]
	return
}

func decodeLong(src *[]byte) (outByte []byte, outVal int16, err error) {
	if len(*src) < 2 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:2]
	outVal |= int16(outByte[0]) << 8
	outVal |= int16(outByte[1])
	(*src) = (*src)[2:]
	return
}

func decodeUnsigned(src *[]byte) (outByte []byte, outVal uint8, err error) {
	outByte = (*src)[:1]
	outVal = uint8(outByte[0])
	(*src) = (*src)[1:]
	return
}

func decodeLongUnsigned(src *[]byte) (outByte []byte, outVal uint16, err error) {
	if len(*src) < 2 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:2]
	outVal |= uint16(outByte[0]) << 8
	outVal |= uint16(outByte[1])
	(*src) = (*src)[2:]
	return
}

func decodeLong64(src *[]byte) (outByte []byte, outVal int64, err error) {
	if len(*src) < 8 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:8]
	outVal |= int64(outByte[0]) << 56
	outVal |= int64(outByte[1]) << 48
	outVal |= int64(outByte[2]) << 40
	outVal |= int64(outByte[3]) << 32
	outVal |= int64(outByte[4]) << 24
	outVal |= int64(outByte[5]) << 16
	outVal |= int64(outByte[6]) << 8
	outVal |= int64(outByte[7])
	(*src) = (*src)[8:]
	return
}

func decodeLong64Unsigned(src *[]byte) (outByte []byte, outVal uint64, err error) {
	if len(*src) < 8 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:8]
	outVal |= uint64(outByte[0]) << 56
	outVal |= uint64(outByte[1]) << 48
	outVal |= uint64(outByte[2]) << 40
	outVal |= uint64(outByte[3]) << 32
	outVal |= uint64(outByte[4]) << 24
	outVal |= uint64(outByte[5]) << 16
	outVal |= uint64(outByte[6]) << 8
	outVal |= uint64(outByte[7])
	(*src) = (*src)[8:]
	return
}

func decodeEnum(src *[]byte) (outByte []byte, outVal uint8, err error) {
	outByte = (*src)[:1]
	outVal = uint8(outByte[0])
	(*src) = (*src)[1:]
	return
}

func decodeFloat32(src *[]byte) (outByte []byte, outVal float32, err error) {
	if len(*src) < 4 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:4]
	outVal = math.Float32frombits(binary.BigEndian.Uint32(outByte))
	(*src) = (*src)[4:]
	return
}

func decodeFloat64(src *[]byte) (outByte []byte, outVal float64, err error) {
	if len(*src) < 8 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:8]
	outVal = math.Float64frombits(binary.BigEndian.Uint64(outByte))
	(*src) = (*src)[8:]
	return
}

// Decode 5 bytes data into time.Time object
// year highbyte,
// year lowbyte,
// month,
// day of month,
// day of week
func decodeDate(src *[]byte) (outByte []byte, outVal time.Time, err error) {
	if len(*src) < 5 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:5]

	year := int(binary.BigEndian.Uint16(outByte[0:2]))
	month := int(outByte[2])
	day := int(outByte[3])
	// weekday := int(outByte[4])

	outVal = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	(*src) = (*src)[5:]
	return
}

// Decode 4 bytes data into time.Time object
// hour,
// minute,
// second,
// hundredths
func decodeTime(src *[]byte) (outByte []byte, outVal time.Time, err error) {
	if len(*src) < 4 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:4]

	hour := int(outByte[0])
	minute := int(outByte[1])
	second := int(outByte[2])
	hundredths := int(outByte[3])

	outVal = time.Date(0, time.Month(1), 1, hour, minute, second, hundredths, time.UTC)

	(*src) = (*src)[4:]
	return
}

// Decode 12 bytes data into time.Time object
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
func decodeDateTime(src *[]byte) (outByte []byte, outVal time.Time, err error) {
	if len(*src) < 12 {
		err = ErrLengthLess
		return
	}
	outByte = (*src)[:12]

	if outByte[11] == 0xff {
		err = fmt.Errorf("Clock status value(%v) not OK(0x00)", outByte[11])
		return
	}

	year := int(binary.BigEndian.Uint16(outByte[0:2]))
	month := int(outByte[2])
	day := int(outByte[3])
	// weekday := int(outByte[4])
	hour := int(outByte[5])
	minute := int(outByte[6])
	second := int(outByte[7])
	hundredths := int(outByte[8])
	// deviation := int(binary.BigEndian.Uint16(outByte[9:11]))

	outVal = time.Date(year, time.Month(month), day, hour, minute, second, hundredths, time.UTC)
	(*src) = (*src)[12:]
	return
}
