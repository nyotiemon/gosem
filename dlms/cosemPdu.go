package cosem

import "fmt"

type cosemTag uint8

var (
	// ---- standardized DLMS APDUs
	TagInitiateRequest          cosemTag = 1
	TagReadRequest              cosemTag = 5
	TagWriteRequest             cosemTag = 6
	TagInitiateResponse         cosemTag = 8
	TagReadResponse             cosemTag = 1
	TagWriteResponse            cosemTag = 4
	TagUnconfirmedWriteRequest  cosemTag = 2
	TagInformationReportRequest cosemTag = 2
	// --- APDUs used for data communication services
	TagGetRequest               cosemTag = 192
	TagSetRequest               cosemTag = 193
	TagEventNotificationRequest cosemTag = 194
	TagActionRequest            cosemTag = 195
	TagGetResponse              cosemTag = 196
	TagSetResponse              cosemTag = 197
	TagActionResponse           cosemTag = 199
	// --- global ciphered pdus
	TagGloGetRequest               cosemTag = 200
	TagGloSetRequest               cosemTag = 201
	TagGloEventNotificationRequest cosemTag = 202
	TagGloActionRequest            cosemTag = 203
	TagGloGetResponse              cosemTag = 204
	TagGloSetResponse              cosemTag = 205
	TagGloActionResponse           cosemTag = 207
	// --- dedicated ciphered pdus
	TagDedGetRequest               cosemTag = 208
	TagDedSetRequest               cosemTag = 209
	TagDedEventNotificationRequest cosemTag = 210
	TagDedActionRequest            cosemTag = 211
	TagDedGetResponse              cosemTag = 212
	TagDedSetResponse              cosemTag = 213
	TagDedActionResponse           cosemTag = 215
	TagExceptionResponse           cosemTag = 216
)

func ErrWrongTag(idx int, get byte, correct byte) error {
	return fmt.Errorf("wrong data tag on index %v, expecting %v instead of %v", idx, correct, get)
}

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s cosemTag) Value() uint8 {
	return uint8(s)
}

func (s cosemTag) isExist(bt byte) bool {
	switch bt {
	case
		TagGetRequest.Value(),
		TagSetRequest.Value(),
		TagEventNotificationRequest.Value(),
		TagActionRequest.Value(),
		TagGetResponse.Value(),
		TagSetResponse.Value(),
		TagActionResponse.Value():
		return true
	}

	return false
}

type CosemI interface {
	New() (out CosemPDU, err error)
	Decode() (out CosemPDU, err error)
}

type CosemPDU interface {
	Encode() ([]byte, error)
}

// DecodeCosem is a global function to decode payload based on implemented DLMS/COSEM APDU en/decoder
func DecodeCosem(src *[]byte) (out CosemPDU, err error) {
	var t cosemTag
	if !t.isExist((*src)[0]) {
		err = fmt.Errorf("byte idx 0 (%v) is not recognized, or relevant DLMS/COSEM is not yet implemented.", (*src)[0])
		return
	}

	switch (*src)[0] {
	case TagGetRequest.Value():
		var decoder GetRequest
		out, err = decoder.Decode(src)
	case TagSetRequest.Value():
		var decoder SetRequest
		out, err = decoder.Decode(src)
	case TagActionRequest.Value():
		var decoder ActionRequest
		out, err = decoder.Decode(src)
	case TagGetResponse.Value():
		var decoder GetResponse
		out, err = decoder.Decode(src)
	case TagSetResponse.Value():
		var decoder SetResponse
		out, err = decoder.Decode(src)
	case TagActionResponse.Value():
		var decoder ActionResponse
		out, err = decoder.Decode(src)
	case TagEventNotificationRequest.Value():
		out, err = DecodeEventNotificationRequest(src)
	}

	return
}
