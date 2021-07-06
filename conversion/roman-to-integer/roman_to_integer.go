package conversion

import "fmt"

var romanNums = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

// RomanToInteger ...
func RomanToInteger(roman string) int {
	result := 0
	i := 0

	for i < len(roman) {
		current := romanNums[string(roman[i])]
		if i+1 < len(roman) && (current < romanNums[string(roman[i+1])]) {
			result += (romanNums[string(roman[i+1])] - current)
			i += 2
		} else {
			result += current
			i++
		}
	}
	var x = fmt.Sprintf("result %d", result)
	fmt.Println(x)
	return result
}
