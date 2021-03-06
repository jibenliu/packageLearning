// Code generated by "stringer -type=Pill"; DO NOT EDIT.

package stringKiller

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Placebo-0]
	_ = x[Aspirin-1]
	_ = x[Ibuprofen-2]
	_ = x[Paracetamol-3]
	_ = x[Acetaminophen-4]
	_ = x[COVID19-5]
	_ = x[HIV-6]
	_ = x[OIV-7]
	_ = x[DIV-8]
}

const _Pill_name = "PlaceboAspirinIbuprofenParacetamolAcetaminophenCOVID19HIVOIVDIV"

var _Pill_index = [...]uint8{0, 7, 14, 23, 34, 47, 54, 57, 60, 63}

func (i Pill) String() string {
	if i < 0 || i >= Pill(len(_Pill_index)-1) {
		return "Pill(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
