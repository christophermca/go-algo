package util

import "math/rand"

// GenerateSockPile return int[]
func GenerateSockPile(pileSize int) []int {
	pile := make([]int, pileSize)
	i := 0
	for i < pileSize {
		min := 1
		max := 4
		pile[i] = rand.Intn(max-min) + min

		i++
	}
	return pile
}
