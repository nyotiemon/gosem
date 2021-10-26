package dlms

import (
	"bytes"
	"gosem/pkg/axdr"
	"time"
)

type accesSelector uint8

const (
	AccessSelectorRange accesSelector = 0x1
	AccessSelectorEntry accesSelector = 0x2
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s accesSelector) Value() uint8 {
	return uint8(s)
}

type SelectiveAccessDescriptor struct {
	AccessSelector  accesSelector
	AccessParameter axdr.DlmsData
}

func CreateSelectiveAccessDescriptor(as accesSelector, ap interface{}) *SelectiveAccessDescriptor {
	if as == AccessSelectorRange {
		// make sure AccessParameter is a [2]time.Time
		ranges := ap.([]time.Time)
		// selector range should be of:
		// structure { structure {classid, obis, attributeid, dataidx}, range-start, range-end, selected val }
		var classId axdr.DlmsData = *axdr.CreateAxdrLongUnsigned(8)
		var obisCode axdr.DlmsData = *axdr.CreateAxdrOctetString("0.0.1.0.0.255") // obis of clock
		var attributeId axdr.DlmsData = *axdr.CreateAxdrInteger(2)
		var dataIdx axdr.DlmsData = *axdr.CreateAxdrLongUnsigned(0)
		var rangeStart axdr.DlmsData = *axdr.CreateAxdrDateTime(ranges[0])
		var rangeEnd axdr.DlmsData = *axdr.CreateAxdrDateTime(ranges[1])
		var selectedValue axdr.DlmsData = *axdr.CreateAxdrArray([]*axdr.DlmsData{})

		var restrictingObject axdr.DlmsData = *axdr.CreateAxdrStructure([]*axdr.DlmsData{&classId, &obisCode, &attributeId, &dataIdx})
		var rangeDescriptor axdr.DlmsData = *axdr.CreateAxdrStructure([]*axdr.DlmsData{&restrictingObject, &rangeStart, &rangeEnd, &selectedValue})

		return &SelectiveAccessDescriptor{AccessSelector: as, AccessParameter: rangeDescriptor}
	} else {
		// make sure AccessParameter is a [2]uint32
		entries := ap.([]uint32)
		// selector enty should be of:
		// structure {fromEntry, toEntry, fromSelectedValue, toSelectedValue}
		var fromEntry axdr.DlmsData = *axdr.CreateAxdrDoubleLongUnsigned(entries[0])
		var toEntry axdr.DlmsData = *axdr.CreateAxdrDoubleLongUnsigned(entries[1])
		var fromSelectedValue axdr.DlmsData = *axdr.CreateAxdrLongUnsigned(0)
		var toSelectedValue axdr.DlmsData = *axdr.CreateAxdrLongUnsigned(0)

		var entryDescriptor axdr.DlmsData = *axdr.CreateAxdrStructure([]*axdr.DlmsData{&fromEntry, &toEntry, &fromSelectedValue, &toSelectedValue})
		return &SelectiveAccessDescriptor{AccessSelector: as, AccessParameter: entryDescriptor}
	}
}

func (s SelectiveAccessDescriptor) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(s.AccessSelector.Value()))
	val, e := s.AccessParameter.Encode()
	if e != nil {
		err = e
		return
	}
	buf.Write(val)

	out = buf.Bytes()
	return
}

func DecodeSelectiveAccessDescriptor(ori *[]byte) (out SelectiveAccessDescriptor, err error) {
	src := append([]byte(nil), (*ori)...)

	if src[0] == AccessSelectorRange.Value() {
		out.AccessSelector = AccessSelectorRange
	} else {
		out.AccessSelector = AccessSelectorEntry
	}
	src = src[1:] // remove access-selector byte

	var axdrDecoder axdr.Decoder = *axdr.NewDataDecoder(&src)
	out.AccessParameter, err = axdrDecoder.Decode(&src)
	if err != nil {
		return
	}

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
