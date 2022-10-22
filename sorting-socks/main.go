package sortingsocks

import (
	"fmt"
	"log"
	"sort"
)

var sizeOfPile int

func isEven(size int) bool {
	if size > 0 {
		return size%2 == 0
	}
	return false
}

func groupSocks(socks *[]int, compare func(a, b int) bool) *[][]int {
	sox := *socks
	group := [][]int{}
	sizeOfPile = len(sox)
	fmt.Println("sox:: ", sox)

	//group = append(group, sox)
	//binary search to find end of type.

	if isEven(sizeOfPile) {
		//half := sizeOfPile/2 - 1
	} else {
		//half := sizeOfPile / 2
	}

	end := cap(sox) - 1
	for len(sox) > 0 {
		log.Printf("sizeOfPile: %d \n {start: %d end: %d }\n___\n", sizeOfPile, 0, end)
		//lo
		for compare(sox[0], sox[end]) {
			log.Printf("{start: { idx: %d, value: %d},end: {idx:%d,  value: %d}}\n___\n", 0, sox[0], end, sox[end])
			end = end / 2
		}
		//hi
		// for compare(sox[end], sox[len(sox)-1]) {
		// 	log.Printf("{start: %d, end: %d}\n___\n", end, len(sox)-1)
		// 	end = half * 2
		// }

		sockType := sox[:end+1]
		log.Printf("sockType: %d", end)

		group = append(group, sockType)

		if len(sox) > end+1 {
			sox = sox[end:]
		} else {
			sox = sox[end+1:]
		}
		fmt.Println("new sox", sox)
	}
	break

	return &group
}

func isNewPair(a, b int) bool {
	compare := a != b

	log.Printf("isNewPair::: %d, %d, %t", a, b, compare)
	// update "anchor-index"

	/**
	 * if start == end  -> true // all same sockType
	 * if start != end -> binary search(end/2) check low range start - low |
	 * high range to array lenth for end of sockType
	 */

	return compare

	// if a == b {
	// 	return true
	// }

	// if a != b {
	// 	if a < b {
	// 		return false
	// 	}
	// 	return true
	// }
	// return true
}

// SortingSocks (sockData := []int{...n(1:3)})
func SortingSocks(sockData []int) []int {
	sort.Ints(sockData)
	grouped := groupSocks(&sockData, isNewPair)
	log.Println("sockData: :: ", *grouped)
	return sockData
}
