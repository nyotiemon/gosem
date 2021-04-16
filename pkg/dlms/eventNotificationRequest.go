package dlms

import (
	"bytes"
	"time"
	"gosem/pkg/axdr"
)

type EventNotificationRequest struct {
	Time           *time.Time
	AttributeInfo  AttributeDescriptor
	AttributeValue axdr.DlmsData
}

func (ev EventNotificationRequest) getTimebyValue() time.Time {
	return *ev.Time
}

func CreateEventNotificationRequest(tm *time.Time, attInfo AttributeDescriptor, attValue axdr.DlmsData) *EventNotificationRequest {
	return &EventNotificationRequest{
		Time:           tm,
		AttributeInfo:  attInfo,
		AttributeValue: attValue,
	}
}

func (ev EventNotificationRequest) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagEventNotificationRequest))
	if ev.Time == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		tm, e := axdr.EncodeDateTime(ev.getTimebyValue())
		if e != nil {
			err = e
			return
		}
		buf.WriteByte(uint8(len(tm)))
		buf.Write(tm)
	}
	attInfo, eInfo := ev.AttributeInfo.Encode()
	if eInfo != nil {
		err = eInfo
		return
	}
	buf.Write(attInfo)
	attValue, eValue := ev.AttributeValue.Encode()
	if eValue != nil {
		err = eValue
		return
	}
	buf.Write(attValue)

	out = buf.Bytes()
	return
}

func DecodeEventNotificationRequest(ori *[]byte) (out EventNotificationRequest, err error) {
	var src []byte = append((*ori)[:0:0], (*ori)...)

	if src[0] != TagEventNotificationRequest.Value() {
		err = ErrWrongTag(0, src[0], byte(TagEventNotificationRequest))
		return
	}

	haveTime := src[1]
	src = src[2:]

	if haveTime == 0x0 {
		var nilTime *time.Time = nil
		out.Time = nilTime
	} else {
		src = src[1:] // length of time
		_, time, e := axdr.DecodeDateTime(&src)
		if e != nil {
			err = e
			return
		}
		out.Time = &time
	}

	out.AttributeInfo, err = DecodeAttributeDescriptor(&src)
	if err != nil {
		return
	}

	decoder := axdr.NewDataDecoder(&src)
	out.AttributeValue, err = decoder.Decode(&src)

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
