// Code generated by "stringer -type=State"; DO NOT EDIT.

package staking

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StateBonding-0]
	_ = x[StateBonded-1]
	_ = x[StateUnbonded-2]
	_ = x[StateUnbonding-3]
	_ = x[StateUnregistered-4]
}

const _State_name = "StateBondingStateBondedStateUnbondedStateUnbondingStateUnregistered"

var _State_index = [...]uint8{0, 12, 23, 36, 50, 67}

func (i State) String() string {
	if i >= State(len(_State_index)-1) {
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}
