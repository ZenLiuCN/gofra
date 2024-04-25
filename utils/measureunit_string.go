// Code generated by "stringer -type MeasureUnit"; DO NOT EDIT.

package utils

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NANOS-0]
	_ = x[MICROS-1]
	_ = x[MILLS-2]
	_ = x[SECONDS-3]
}

const _MeasureUnit_name = "NANOSMICROSMILLSSECONDS"

var _MeasureUnit_index = [...]uint8{0, 5, 11, 16, 23}

func (i MeasureUnit) String() string {
	if i < 0 || i >= MeasureUnit(len(_MeasureUnit_index)-1) {
		return "MeasureUnit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MeasureUnit_name[_MeasureUnit_index[i]:_MeasureUnit_index[i+1]]
}
