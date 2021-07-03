package conversion

var romanNums = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

// RomanToInteger ...
func RomanToInteger(roman string) int {
	var result int
	var i = 0

	for i < len(roman) {
		result += romanNums[string(roman[i])]
	}
	return result
}
