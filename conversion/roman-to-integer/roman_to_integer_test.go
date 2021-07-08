package conversion

import "testing"

type testData struct {
	name     string
	input    string
	expected int
}

var romanToIntegerTests = []testData{
	{input: "I", expected: 1, name: "I_1"},
	{input: "III", expected: 3, name: "III_3"},
	{input: "IV", expected: 4, name: "IV_4"},
	{input: "VII", expected: 7, name: "IV_7"},
	{input: "VII", expected: 7, name: "VII_7"},
	{input: "XI", expected: 11, name: "XI_11"},
	{input: "XL", expected: 40, name: "XL_40"},
	{input: "DCCLXXXIX", expected: 789, name: "DCCLXXXIX-789"},
}

func TestRomanToInteger(t *testing.T) {
	for _, test := range romanToIntegerTests {
		result := RomanToInteger(test.input)
		if result != test.expected {
			t.Error(test.name)
		}
	}
}
