// Code generated by "stringer -type=respirationState"; DO NOT EDIT

package xethru

import "fmt"

const _respirationState_name = "breathingmovementtrackingnoMovementinitializingstateReservedstateUnknownsomeotherState"

var _respirationState_index = [...]uint8{0, 9, 17, 25, 35, 47, 60, 72, 86}

func (i respirationState) String() string {
	if i >= respirationState(len(_respirationState_index)-1) {
		return fmt.Sprintf("respirationState(%d)", i)
	}
	return _respirationState_name[_respirationState_index[i]:_respirationState_index[i+1]]
}