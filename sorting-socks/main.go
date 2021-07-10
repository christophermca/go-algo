package sortingsocks

import (
	"fmt"
	"sort"
)

// SortingSocks (sockData := []int{...n(1:3)})
func SortingSocks(sockData []int) []int {
	if len(sockData) > 0 {
		sort.Ints(sockData)
		fmt.Println(sockData)
	}
	return sockData
}
