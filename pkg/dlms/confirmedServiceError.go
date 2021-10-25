package dlms

import "bytes"

type confirmedServiceErrorTag uint8

const (
	TagErrInitiateError confirmedServiceErrorTag = 1
	TagErrRead          confirmedServiceErrorTag = 5
	TagErrWrite         confirmedServiceErrorTag = 6
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s confirmedServiceErrorTag) Value() uint8 {
	return uint8(s)
}

type serviceErrorTag uint8

const (
	TagErrApplicationReference serviceErrorTag = 0
	TagErrHardwareResource     serviceErrorTag = 1
	TagErrVdeStateError        serviceErrorTag = 2
	TagErrService              serviceErrorTag = 3
	TagErrDefinition           serviceErrorTag = 4
	TagErrAccess               serviceErrorTag = 5
	TagErrInitiate             serviceErrorTag = 6
	TagErrLoadDataSet          serviceErrorTag = 7
	TagErrTask                 serviceErrorTag = 9
	TagErrOtherError           serviceErrorTag = 10
)

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s serviceErrorTag) Value() uint8 {
	return uint8(s)
}

type ConfirmedServiceError struct {
	ConfirmedServiceError confirmedServiceErrorTag
	ServiceError          serviceErrorTag
	Value                 uint8
}

func CreateConfirmedServiceError(service confirmedServiceErrorTag, serviceError serviceErrorTag, value uint8) *ConfirmedServiceError {
	return &ConfirmedServiceError{
		ConfirmedServiceError: service,
		ServiceError:          serviceError,
		Value:                 value,
	}
}

func (cse ConfirmedServiceError) Encode() (out []byte, err error) {
	var buf bytes.Buffer
	buf.WriteByte(TagConfirmedServiceError.Value())
	buf.WriteByte(cse.ConfirmedServiceError.Value())
	buf.WriteByte(cse.ServiceError.Value())
	buf.WriteByte(cse.Value)

	out = buf.Bytes()
	return
}

func DecodeConfirmedServiceError(ori *[]byte) (out ConfirmedServiceError, err error) {
	src := append([]byte(nil), (*ori)...)

	if len(src) < 4 {
		err = ErrWrongLength(len(src), 4)
		return
	}

	if src[0] != TagConfirmedServiceError.Value() {
		err = ErrWrongTag(0, src[0], byte(TagConfirmedServiceError))
		return
	}

	out.ConfirmedServiceError = confirmedServiceErrorTag(src[1])
	out.ServiceError = serviceErrorTag(src[2])
	out.Value = src[3]
	src = src[4:]

	(*ori) = (*ori)[len((*ori))-len(src):]
	return
}
