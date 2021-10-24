package dlms

import (
	"bytes"
)

type stateErrorTag uint8

const (
	TagServiceNotAllowed stateErrorTag = 0x1
	TagServiceUnknown    stateErrorTag = 0x2
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s stateErrorTag) Value() uint8 {
	return uint8(s)
}

type serviceErrorTag uint8

const (
	TagOperationNotPossible serviceErrorTag = 0x1
	TagServiceNotSupported  serviceErrorTag = 0x2
	TagOtherReason          serviceErrorTag = 0x3
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s serviceErrorTag) Value() uint8 {
	return uint8(s)
}

type ExceptionResponse struct {
	StateError   stateErrorTag
	ServiceError serviceErrorTag
}

func CreateExceptionResponse(stateError stateErrorTag, serviceError serviceErrorTag) *ExceptionResponse {
	return &ExceptionResponse{
		StateError:   stateError,
		ServiceError: serviceError,
	}
}

func (er ExceptionResponse) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(TagExceptionResponse.Value())
	buf.WriteByte(er.StateError.Value())
	buf.WriteByte(er.ServiceError.Value())

	out = buf.Bytes()
	return
}

func DecodeExceptionResponse(ori *[]byte) (out ExceptionResponse, err error) {
	src := append([]byte(nil), (*ori)...)

	if len(src) < 3 {
		err = ErrWrongLength(len(src), 3)
		return
	}

	if src[0] != TagExceptionResponse.Value() {
		err = ErrWrongTag(0, src[0], byte(TagExceptionResponse))
		return
	}

	out.StateError = stateErrorTag(src[1])
	out.ServiceError = serviceErrorTag(src[2])
	src = src[3:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
