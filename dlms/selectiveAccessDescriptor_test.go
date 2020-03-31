package cosem

import (
	"bytes"
	"testing"
	"time"
)

func TestSelectiveAccessDescriptor_Encode(t *testing.T) {
	var a SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorEntry, []uint32{0, 5})
	t1 := a.Encode()
	result := []byte{2, 2, 4, 6, 0, 0, 0, 0, 6, 0, 0, 0, 5, 18, 0, 0, 18, 0, 0}

	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("Test AccessSelectorEntry failed. get: %d, should:%v", t1, result)
	}

	timeStart := time.Date(2020, time.January, 1, 10, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2020, time.January, 1, 11, 0, 0, 0, time.UTC)
	var b SelectiveAccessDescriptor = *CreateSelectiveAccessDescriptor(AccessSelectorRange, []time.Time{timeStart, timeEnd})
	t2 := b.Encode()
	result = []byte{1, 2, 4, 2, 4, 18, 0, 8, 9, 6, 0, 0, 1, 0, 0, 255, 15, 2, 18, 0, 0, 25, 7, 228, 1, 1, 3, 10, 0, 0, 0, 0, 0, 0, 25, 7, 228, 1, 1, 3, 11, 0, 0, 0, 0, 0, 0, 1, 0}

	res = bytes.Compare(t2, result)
	if res != 0 {
		t.Errorf("Test AccessSelectorRange failed. get: %d, should:%v", t2, result)
	}
}
