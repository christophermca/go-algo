package util

import (
	"reflect"
	"testing"
)

// GenerateSockPile return int[]
func TestGenerateSockPile(t *testing.T) {
	expect := GenerateSockPile(20)
	reflect.ValueOf(expect).Len()
}
