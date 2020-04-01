package cosem

import "fmt"

type accessResultTag uint8
type actionResultTag uint8

const (
	// DataAccessResult
	TagAccSuccess                 accessResultTag = 0
	TagAccHardwareFault           accessResultTag = 1
	TagAccTemporaryFailure        accessResultTag = 2
	TagAccReadWriteDenied         accessResultTag = 3
	TagAccObjectUndefined         accessResultTag = 4
	TagAccObjectClassInconsistent accessResultTag = 9
	TagAccObjectUnavailable       accessResultTag = 11
	TagAccTypeUnmatched           accessResultTag = 12
	TagAccScopeAccessViolated     accessResultTag = 13
	TagAccDataBlockUnavailable    accessResultTag = 14
	TagAccLongGetAborted          accessResultTag = 15
	TagAccNoLongGetInProgress     accessResultTag = 16
	TagAccLongSetAborted          accessResultTag = 17
	TagAccNoLongSetInProgress     accessResultTag = 18
	TagAccDataBlockNumberInvalid  accessResultTag = 19
	TagAccOtherReason             accessResultTag = 250

	// ActionResult
	TagActSuccess                 actionResultTag = 0
	TagActHardwareFault           actionResultTag = 1
	TagActTemporaryFailure        actionResultTag = 2
	TagActReadWriteDenied         actionResultTag = 3
	TagActObjectUndefined         actionResultTag = 4
	TagActObjectClassInconsistent actionResultTag = 9
	TagActObjectUnavailable       actionResultTag = 11
	TagActTypeUnmatched           actionResultTag = 12
	TagActScopeOfAccessViolated   actionResultTag = 13
	TagActDataBlockUnavailable    actionResultTag = 14
	TagActLongActionAborted       actionResultTag = 15
	TagActNoLongActionInProgress  actionResultTag = 16
	TagActOtherReason             actionResultTag = 250
)

func (tag accessResultTag) String() string {
	switch tag {
	case TagAccSuccess:
		return "success"
	case TagAccHardwareFault:
		return "hardware-fault"
	case TagAccTemporaryFailure:
		return "temporary-failure"
	case TagAccReadWriteDenied:
		return "read-write-denied"
	case TagAccObjectUndefined:
		return "object-undefined"
	case TagAccObjectClassInconsistent:
		return "object-class-inconsistent"
	case TagAccObjectUnavailable:
		return "object-unavailable"
	case TagAccTypeUnmatched:
		return "type-unmatched"
	case TagAccScopeAccessViolated:
		return "scope-of-access-violated"
	case TagAccDataBlockUnavailable:
		return "data-block-unavailable"
	case TagAccLongGetAborted:
		return "long-get-aborted"
	case TagAccNoLongGetInProgress:
		return "no-long-get-in-progress"
	case TagAccLongSetAborted:
		return "long-set-aborted"
	case TagAccNoLongSetInProgress:
		return "no-long-set-in-progress"
	case TagAccDataBlockNumberInvalid:
		return "data-block-number-invalid"
	case TagAccOtherReason:
		return "other-reason"
	default:
		return ""
	}
}

func (tag actionResultTag) String() string {
	switch tag {
	case TagActSuccess:
		return "success"
	case TagActHardwareFault:
		return "hardware-fault"
	case TagActTemporaryFailure:
		return "temporary-failure"
	case TagActReadWriteDenied:
		return "read-write-denied"
	case TagActObjectUndefined:
		return "object-undefined"
	case TagActObjectClassInconsistent:
		return "object-class-inconsistent"
	case TagActObjectUnavailable:
		return "object-unavailable"
	case TagActTypeUnmatched:
		return "type-unmatched"
	case TagActScopeOfAccessViolated:
		return "scope-of-access-violated"
	case TagActDataBlockUnavailable:
		return "data-block-unavailable"
	case TagActLongActionAborted:
		return "long-action-aborted"
	case TagActNoLongActionInProgress:
		return "no-long-action-in-progress"
	case TagActOtherReason:
		return "other-reason"
	default:
		return ""
	}
}

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s accessResultTag) Value() uint8 {
	return uint8(s)
}

// Value will return primitive value of the target.
// This is used for comparing with non custom typed object
func (s actionResultTag) Value() uint8 {
	return uint8(s)
}

func GetAccessTag(in uint8) (out accessResultTag, err error) {
	switch in {
	case 0:
		out = TagAccSuccess
	case 1:
		out = TagAccHardwareFault
	case 2:
		out = TagAccTemporaryFailure
	case 3:
		out = TagAccReadWriteDenied
	case 4:
		out = TagAccObjectUndefined
	case 9:
		out = TagAccObjectClassInconsistent
	case 11:
		out = TagAccObjectUnavailable
	case 12:
		out = TagAccTypeUnmatched
	case 13:
		out = TagAccScopeAccessViolated
	case 14:
		out = TagAccDataBlockUnavailable
	case 15:
		out = TagAccLongGetAborted
	case 16:
		out = TagAccNoLongGetInProgress
	case 17:
		out = TagAccLongSetAborted
	case 18:
		out = TagAccNoLongSetInProgress
	case 19:
		out = TagAccDataBlockNumberInvalid
	case 250:
		out = TagAccOtherReason
	default:
		err = fmt.Errorf("Value not recognized.")
	}

	return
}

func GetActionTag(in uint8) (out actionResultTag, err error) {
	switch in {
	case 0:
		out = TagActSuccess
	case 1:
		out = TagActHardwareFault
	case 2:
		out = TagActTemporaryFailure
	case 3:
		out = TagActReadWriteDenied
	case 4:
		out = TagActObjectUndefined
	case 9:
		out = TagActObjectClassInconsistent
	case 11:
		out = TagActObjectUnavailable
	case 12:
		out = TagActTypeUnmatched
	case 13:
		out = TagActScopeOfAccessViolated
	case 14:
		out = TagActDataBlockUnavailable
	case 15:
		out = TagActLongActionAborted
	case 16:
		out = TagActNoLongActionInProgress
	case 250:
		out = TagActOtherReason
	default:
		err = fmt.Errorf("Value not recognized")
	}

	return
}
