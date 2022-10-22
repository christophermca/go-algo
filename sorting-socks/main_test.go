package sortingsocks

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortingSocks(t *testing.T) {
	// t.Run("data_is_empty", testSortingSocks([]int{}))
	// t.Run("data_is_sorted", testSortingSocks([]int{1, 2, 3}))
	t.Run("data_is_grouped", testSortingSocks([]int{1, 2, 3, 1, 2, 3}))
	// t.Run("data_is_all_one_type", testSortingSocks([]int{1, 1, 1, 1, 1, 1}))
	// t.Run("data_is_unsorted", testSortingSocks([]int{1, 2, 1, 1, 1, 1, 3, 2, 2}))
}

func testSortingSocks(socks []int) func(*testing.T) {
	return func(t *testing.T) {
		fmt.Println(t.Name())
		result := SortingSocks(socks)
		var expect = []int{1, 1, 2, 2, 3, 3}
		fmt.Println(result, expect)

		if !sort.IntsAreSorted(result) {
			t.Error("socks are not sorted")
		}
	}
}
