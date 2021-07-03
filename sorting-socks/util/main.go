package util

import "math/rand"

// GenerateSockPile return int[]
func GenerateSockPile(pileSize int) []int {
	pile := make([]int, pileSize)
	//fmt.Printf("number of socks %s\n", len(pile))
	i := 0
	for i < pileSize {
		pile[i] = rand.IntN(3)

		i++

	}
	return pile
}
