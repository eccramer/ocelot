// Code generated by "stringer -type=TriggerType,Conditional"; DO NOT EDIT.

package runtime

import "strconv"

const _TriggerType_name = "TNoneBranchFilepathText"

var _TriggerType_index = [...]uint8{0, 5, 11, 19, 23}

func (i TriggerType) String() string {
	if i < 0 || i >= TriggerType(len(_TriggerType_index)-1) {
		return "TriggerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TriggerType_name[_TriggerType_index[i]:_TriggerType_index[i+1]]
}

const _Conditional_name = "CNoneOrAnd"

var _Conditional_index = [...]uint8{0, 5, 7, 10}

func (i Conditional) String() string {
	if i < 0 || i >= Conditional(len(_Conditional_index)-1) {
		return "Conditional(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Conditional_name[_Conditional_index[i]:_Conditional_index[i+1]]
}
