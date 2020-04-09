package cosem

import (
	"bytes"
	. "gosem/axdr"
	"testing"
	"time"
)

func TestNew_EventNotificationRequest(t *testing.T) {
	tm := time.Date(1500, time.January, 1, 0, 0, 0, 255, time.UTC)
	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var attrVal DlmsData = *CreateAxdrBoolean(true)

	var a EventNotificationRequest = *CreateEventNotificationRequest(&tm, attrDesc, attrVal)
	t1 := a.Encode()
	result := []byte{194, 1, 12, 5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0, 0, 1, 1, 0, 0, 3, 0, 255, 2, 3, 255}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}

	// --- with nil time
	var nilTime *time.Time = nil
	a = *CreateEventNotificationRequest(nilTime, attrDesc, attrVal)
	t1 = a.Encode()
	result = []byte{194, 0, 0, 1, 1, 0, 0, 3, 0, 255, 2, 3, 255}

	res = bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t2 Failed. get: %d, should:%v", t1, result)
	}
}

func TestDecode_EventNotificationRequest(t *testing.T) {
	src := []byte{194, 1, 12, 5, 220, 1, 1, 1, 0, 0, 0, 255, 0, 0, 0, 0, 1, 1, 0, 0, 3, 0, 255, 2, 3, 255}
	a, err := DecodeEventNotificationRequest(&src)
	if err != nil {
		t.Errorf("t1 failed on DecodeEventNotificationRequest. Err: %v", err)
	}

	tm := time.Date(1500, time.January, 1, 0, 0, 0, 255, time.UTC)
	var attrDesc AttributeDescriptor = *CreateAttributeDescriptor(1, "1.0.0.3.0.255", 2)
	var attrVal DlmsData = *CreateAxdrBoolean(true)
	var b EventNotificationRequest = *CreateEventNotificationRequest(&tm, attrDesc, attrVal)

	if *a.Time != *b.Time {
		t.Errorf("t1 err Time. get: %v, should: %v", a.Time, b.Time)
	}

	if a.AttributeInfo.ClassId != b.AttributeInfo.ClassId {
		t.Errorf("t1 Failed. AttributeInfo.ClassId get: %v, should:%v", a.AttributeInfo.ClassId, b.AttributeInfo.ClassId)
	}
	res := bytes.Compare(a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	if res != 0 {
		t.Errorf("t1 Failed. AttributeInfo.InstanceId get: %v, should:%v", a.AttributeInfo.InstanceId.Bytes(), b.AttributeInfo.InstanceId.Bytes())
	}
	if a.AttributeInfo.AttributeId != b.AttributeInfo.AttributeId {
		t.Errorf("t1 Failed. AttributeInfo.AttributeId get: %v, should:%v", a.AttributeInfo.AttributeId, b.AttributeInfo.AttributeId)
	}

	if a.AttributeValue.Tag != b.AttributeValue.Tag {
		t.Errorf("t1 Failed. AttributeValue.Tag get: %v, should:%v", a.AttributeValue.Tag, b.AttributeValue.Tag)
	}
	if a.AttributeValue.Value != b.AttributeValue.Value {
		t.Errorf("t1 Failed. AttributeValue.Value get: %v, should:%v", a.AttributeValue.Value, b.AttributeValue.Value)
	}

	// --- with nil time
	src = []byte{194, 0, 0, 1, 1, 0, 0, 3, 0, 255, 2, 3, 255}
	a, err = DecodeEventNotificationRequest(&src)
	if err != nil {
		t.Errorf("t2 failed on DecodeEventNotificationRequest. Err: %v", err)
	}

	if a.Time != nil {
		t.Errorf("t2 err Time should be nil. get: %v", a.Time)
	}
}
