package Util

import "strconv"

// Max Int选大的
func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// Min Int选小的
func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// StrToIntMax Str转Int选大的
func StrToIntMax(i, j string) string {
	iv, err := strconv.Atoi(i)
	jv, err := strconv.Atoi(j)
	if err != nil {
		return "NAN"
	}
	if iv > jv {
		return i
	} else {
		return j
	}
}

// StrToIntMin Str转Int选小的
func StrToIntMin(i, j string) string {
	iv, err := strconv.Atoi(i)
	jv, err := strconv.Atoi(j)
	if err != nil {
		return "NAN"
	}
	if iv < jv {
		return i
	} else {
		return j
	}
}

// StrToFloatMax Str转Float选大的
func StrToFloatMax(i, j string) string {
	fi, err := strconv.ParseFloat(i, 64)
	fj, err := strconv.ParseFloat(j, 64)
	if err != nil {
		return "NAN"
	}
	if fi > fj {
		return i
	} else {
		return j
	}
}

// StrToFloatMin Str转Float选小的
func StrToFloatMin(i, j string) string {
	fi, err := strconv.ParseFloat(i, 64)
	fj, err := strconv.ParseFloat(j, 64)
	if err != nil {
		return "NAN"
	}
	if fi < fj {
		return i
	} else {
		return j
	}
}
