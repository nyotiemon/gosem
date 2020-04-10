/*
Package axdr provides functions to encode byte arrays
to A-XDR (Adjusted External Data Representation) encoding and back.
It is standardized by IEC 61334-6 standard [4] and used in DLMS APDUs.
*/

package axdr

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type dataTag int

const (
	TagNull               dataTag = 0
	TagArray              dataTag = 1
	TagStructure          dataTag = 2
	TagBoolean            dataTag = 3
	TagBitString          dataTag = 4
	TagDoubleLong         dataTag = 5
	TagDoubleLongUnsigned dataTag = 6
	TagFloatingPoint      dataTag = 7
	TagOctetString        dataTag = 9
	TagVisibleString      dataTag = 10
	TagUTF8String         dataTag = 12
	TagBCD                dataTag = 13
	TagInteger            dataTag = 15
	TagLong               dataTag = 16
	TagUnsigned           dataTag = 17
	TagLongUnsigned       dataTag = 18
	TagCompactArray       dataTag = 19
	TagLong64             dataTag = 20
	TagLong64Unsigned     dataTag = 21
	TagEnum               dataTag = 22
	TagFloat32            dataTag = 23
	TagFloat64            dataTag = 24
	TagDateTime           dataTag = 25
	TagDate               dataTag = 26
	TagTime               dataTag = 27
	TagDontCare           dataTag = 255
)

type DlmsData struct {
	Tag       dataTag
	Value     interface{}
	rawLength []byte
	rawValue  []byte
	raw       bytes.Buffer
}

func CreateAxdrArray(data []*DlmsData) *DlmsData {
	return &DlmsData{Tag: TagArray, Value: data}
}

func CreateAxdrStructure(data []*DlmsData) *DlmsData {
	return &DlmsData{Tag: TagStructure, Value: data}
}

func CreateAxdrBoolean(data bool) *DlmsData {
	return &DlmsData{Tag: TagBoolean, Value: data}
}

func CreateAxdrBitString(data string) *DlmsData {
	data = strings.ReplaceAll(data, " ", "")
	if len(strings.Trim(data, "01")) > 0 {
		panic("Data must be a string of binary, example: 11100000")
	}
	return &DlmsData{Tag: TagBitString, Value: data}
}

func CreateAxdrDoubleLong(data int32) *DlmsData {
	return &DlmsData{Tag: TagDoubleLong, Value: data}
}

func CreateAxdrDoubleLongUnsigned(data uint32) *DlmsData {
	return &DlmsData{Tag: TagDoubleLongUnsigned, Value: data}
}

func CreateAxdrFloatingPoint(data float32) *DlmsData {
	return &DlmsData{Tag: TagFloatingPoint, Value: data}
}

// expect Hex string as input
func CreateAxdrOctetString(data string) *DlmsData {
	return &DlmsData{Tag: TagOctetString, Value: data}
}

// expect ASCII strings as input
func CreateAxdrVisibleString(data string) *DlmsData {
	return &DlmsData{Tag: TagVisibleString, Value: data}
}

// expect UTF-8 strings as input
func CreateAxdrUTF8String(data string) *DlmsData {
	return &DlmsData{Tag: TagUTF8String, Value: data}
}

func CreateAxdrBCD(data int8) *DlmsData {
	return &DlmsData{Tag: TagBCD, Value: data}
}

func CreateAxdrInteger(data int8) *DlmsData {
	return &DlmsData{Tag: TagInteger, Value: data}
}

func CreateAxdrLong(data int16) *DlmsData {
	return &DlmsData{Tag: TagLong, Value: data}
}

func CreateAxdrUnsigned(data uint8) *DlmsData {
	return &DlmsData{Tag: TagUnsigned, Value: data}
}

func CreateAxdrLongUnsigned(data uint16) *DlmsData {
	return &DlmsData{Tag: TagLongUnsigned, Value: data}
}

func CreateAxdrLong64(data uint8) *DlmsData {
	return &DlmsData{Tag: TagLong64, Value: data}
}

func CreateAxdrLong64Unsigned(data uint8) *DlmsData {
	return &DlmsData{Tag: TagLong64Unsigned, Value: data}
}

func CreateAxdrEnum(data uint8) *DlmsData {
	return &DlmsData{Tag: TagEnum, Value: data}
}

func CreateAxdrFloat32(data float32) *DlmsData {
	return &DlmsData{Tag: TagFloat32, Value: data}
}

func CreateAxdrFloat64(data float64) *DlmsData {
	return &DlmsData{Tag: TagFloat64, Value: data}
}

func CreateAxdrDateTime(data time.Time) *DlmsData {
	return &DlmsData{Tag: TagDateTime, Value: data}
}

func CreateAxdrDate(data time.Time) *DlmsData {
	return &DlmsData{Tag: TagDate, Value: data}
}

func CreateAxdrTime(data time.Time) *DlmsData {
	return &DlmsData{Tag: TagTime, Value: data}
}

// Encodes Value of DlmsData object according to the Tag
// It will panic if Value is nil, data type does not match
// the Tag or if failed happen in encoding length/value level.
func (d *DlmsData) Encode() (out []byte, err error) {
	if d.Value == nil {
		err = fmt.Errorf("Value to encode cannot be nil")
		return
	}

	errDataType := fmt.Errorf("cannot encode value %v with tag %v", d.Value, d.Tag)
	var dataLength []byte

	switch d.Tag {
	case TagNull:
		d.rawValue = []byte{0}

	case TagArray:
		data, ok := d.Value.([]*DlmsData)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeArray(data)
		d.rawValue = rawValue

		if dl, errLength := EncodeLength(len(data)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagStructure:
		// what's the difference between array & structure?
		data, ok := d.Value.([]*DlmsData)
		if !ok {
			err = errDataType
		}
		rawValue, _ := EncodeStructure(data)
		d.rawValue = rawValue

		if dl, errLength := EncodeLength(len(data)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagBoolean:
		data, ok := d.Value.(bool)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeBoolean(data)
		d.rawValue = rawValue

	case TagBitString:
		data, ok := d.Value.(string)
		if !ok {
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeBitString(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

		// length of bitstring is count by bits, not bytes
		// length of "1110" is 4, not 1
		if dl, errLength := EncodeLength(len(data)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagDoubleLong:
		data, ok := d.Value.(int32)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeDoubleLong(data)
		d.rawValue = rawValue

	case TagDoubleLongUnsigned:
		data, ok := d.Value.(uint32)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeDoubleLongUnsigned(data)
		d.rawValue = rawValue

	case TagFloatingPoint:
		data, ok := d.Value.(float32)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeFloat32(data)
		d.rawValue = rawValue

	case TagOctetString:
		data, ok := d.Value.(string)
		if !ok {
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeOctetString(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

		if dl, errLength := EncodeLength(len(d.rawValue)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagVisibleString:
		data, ok := d.Value.(string)
		if !ok {
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeVisibleString(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

		if dl, errLength := EncodeLength(len(d.rawValue)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagUTF8String:
		data, ok := d.Value.(string)
		if !ok {
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeUTF8String(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

		if dl, errLength := EncodeLength(len(d.rawValue)); errLength != nil {
			err = errLength
			return
		} else {
			dataLength = dl
		}

	case TagBCD:
		data, ok := d.Value.(int8)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeBCD(data)
		d.rawValue = rawValue

	case TagInteger:
		data, ok := d.Value.(int8)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeInteger(data)
		d.rawValue = rawValue

	case TagLong:
		data, ok := d.Value.(int16)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeLong(data)
		d.rawValue = rawValue

	case TagUnsigned:
		data, ok := d.Value.(uint8)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeUnsigned(data)
		d.rawValue = rawValue

	case TagLongUnsigned:
		data, ok := d.Value.(uint16)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeLongUnsigned(data)
		d.rawValue = rawValue

	case TagCompactArray:
		err = fmt.Errorf("not yet implemented")
		return

	case TagLong64:
		data, ok := d.Value.(int64)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeLong64(data)
		d.rawValue = rawValue

	case TagLong64Unsigned:
		data, ok := d.Value.(uint64)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeLong64Unsigned(data)
		d.rawValue = rawValue

	case TagEnum:
		data, ok := d.Value.(uint8)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeEnum(data)
		d.rawValue = rawValue

	case TagFloat32:
		data, ok := d.Value.(float32)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeFloat32(data)
		d.rawValue = rawValue

	case TagFloat64:
		data, ok := d.Value.(float64)
		if !ok {
			err = errDataType
			return
		}
		rawValue, _ := EncodeFloat64(data)
		d.rawValue = rawValue

	case TagDateTime:
		var data time.Time
		switch d.Value.(type) {
		case time.Time:
			data, _ = d.Value.(time.Time)
		case string:
			// max year value using parse string is 9999, over will give year 0000
			v, _ := d.Value.(string)
			data, _ = time.Parse("2006-01-02 15:04:05", v)
		default:
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeDateTime(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

	case TagDate:
		var data time.Time
		switch d.Value.(type) {
		case time.Time:
			data, _ = d.Value.(time.Time)
		case string:
			v, _ := d.Value.(string)
			data, _ = time.Parse("2006-01-02", v)
		default:
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeDate(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

	case TagTime:
		var data time.Time
		switch d.Value.(type) {
		case time.Time:
			data, _ = d.Value.(time.Time)
		case string:
			v, _ := d.Value.(string)
			data, _ = time.Parse("15:04:05", v)
		default:
			err = errDataType
			return
		}

		if rv, errEncoding := EncodeTime(data); errEncoding != nil {
			err = errEncoding
			return
		} else {
			d.rawValue = rv
		}

	case TagDontCare:
		d.rawValue = []byte{0}

	}

	d.raw.Reset()
	d.raw.WriteByte(byte(d.Tag))
	if len(dataLength) > 0 {
		d.rawLength = dataLength
		d.raw.Write(dataLength)
	}
	d.raw.Write(d.rawValue)
	return d.raw.Bytes(), nil
}

// Return bytes of raw data if Encode() has been called before
// raw data is combination of Tag, Length(if any), and Value
func (d *DlmsData) Raw() []byte {
	return d.raw.Bytes()
}

// Return bytes of raw value if Encode() has been called before
// raw value does not include Tag and Length
func (d *DlmsData) RawValue() []byte {
	return d.rawValue
}

// Return bytes of raw length if Encode() has been called before
// raw length does not include Tag and Value
func (d *DlmsData) RawLength() []byte {
	return d.rawLength
}
