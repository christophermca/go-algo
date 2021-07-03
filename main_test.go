package main

import (
	sortSocks "github.com/christophermca/go-algo/cmd/sortingSocks"
	"testing"
)

func TestMain(t *testing.T) {
	socks := []int{1, 2, 1, 1, 1, 2, 2}
	sortSocks.Run(socks)
}
