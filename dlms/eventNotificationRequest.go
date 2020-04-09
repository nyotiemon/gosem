package cosem

import (
	"bytes"
	. "gosem/axdr"
	"time"
)

type EventNotificationRequest struct {
	Time           *time.Time
	AttributeInfo  AttributeDescriptor
	AttributeValue DlmsData
}

func (ev EventNotificationRequest) getTimebyValue() time.Time {
	return *ev.Time
}

func CreateEventNotificationRequest(tm *time.Time, attInfo AttributeDescriptor, attValue DlmsData) *EventNotificationRequest {
	return &EventNotificationRequest{
		Time:           tm,
		AttributeInfo:  attInfo,
		AttributeValue: attValue,
	}
}

func (ev EventNotificationRequest) Encode() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagEvenNotificationRequest))
	if ev.Time == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		tm, e := EncodeDateTime(ev.getTimebyValue())
		if e != nil {
			panic(e)
		}
		buf.Write(tm)
	}
	buf.Write(ev.AttributeInfo.Encode())
	buf.Write(ev.AttributeValue.Encode())

	return buf.Bytes()
}

func DecodeEventNotificationRequest(src *[]byte) (out EventNotificationRequest, err error) {
	if (*src)[0] != TagEvenNotificationRequest.Value() {
		err = ErrWrongTag(0, (*src)[0], byte(TagEvenNotificationRequest))
		return
	}

	haveTime := (*src)[1]
	(*src) = (*src)[2:]
	if haveTime == 0x0 {
		var nilTime *time.Time = nil
		out.Time = nilTime
	} else {
		_, time, e := DecodeDateTime(src)
		if e != nil {
			err = e
			return
		}
		out.Time = &time
	}

	out.AttributeInfo, err = DecodeAttributeDescriptor(src)
	if err != nil {
		return
	}

	decoder := NewDataDecoder(src)
	out.AttributeValue, err = decoder.Decode(src)

	return
}
