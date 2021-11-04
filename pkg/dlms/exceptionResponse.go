package dlms

import (
	"bytes"
)

type exceptionStateErrorTag uint8

const (
	TagExcServiceNotAllowed exceptionStateErrorTag = 1
	TagExcServiceUnknown    exceptionStateErrorTag = 2
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s exceptionStateErrorTag) Value() uint8 {
	return uint8(s)
}

type exceptionServiceErrorTag uint8

const (
	TagExcOperationNotPossible exceptionServiceErrorTag = 1
	TagExcServiceNotSupported  exceptionServiceErrorTag = 2
	TagExcOtherReason          exceptionServiceErrorTag = 3
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s exceptionServiceErrorTag) Value() uint8 {
	return uint8(s)
}

type ExceptionResponse struct {
	StateError   exceptionStateErrorTag
	ServiceError exceptionServiceErrorTag
}

func CreateExceptionResponse(stateError exceptionStateErrorTag, serviceError exceptionServiceErrorTag) *ExceptionResponse {
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

	out.StateError = exceptionStateErrorTag(src[1])
	out.ServiceError = exceptionServiceErrorTag(src[2])
	src = src[3:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
