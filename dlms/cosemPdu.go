package cosem

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
	TagGetRequest              cosemTag = 192
	TagSetRequest              cosemTag = 193
	TagEvenNotificationRequest cosemTag = 194
	TagActionRequest           cosemTag = 195
	TagGetResponse             cosemTag = 196
	TagSetResponse             cosemTag = 197
	TagActionResponse          cosemTag = 199
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

type CosemI interface {
	New() CosemPDU
}

type CosemPDU interface {
	Encode() []byte
}
