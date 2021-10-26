package dlms

import (
	"bytes"
)

type ExceptionResponse struct {
	StateError   uint8
	ServiceError uint8
}

func CreateExceptionResponse(stateError uint8, serviceError uint8) *ExceptionResponse {
	return &ExceptionResponse{
		StateError:   stateError,
		ServiceError: serviceError,
	}
}

func (er ExceptionResponse) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(byte(TagExceptionResponse))
	buf.WriteByte(byte(er.StateError))
	buf.WriteByte(byte(er.ServiceError))

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

	out.StateError = src[1]
	out.ServiceError = src[2]
	src = src[3:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
