package sortingsocks

import (
	"sort"
	"testing"
)

func TestSortingSocks(t *testing.T) {
	t.Run("{1,2,1,1,3,2,2}", testSortingSocks([]int{1, 2, 1, 1, 3, 2, 2}))
}

func testSortingSocks(socks []int) func(*testing.T) {
	return func(t *testing.T) {
		result := SortingSocks(socks)
		if !sort.IntsAreSorted(result) {
			t.Error("socks are not sorted")
		}
	}
}
