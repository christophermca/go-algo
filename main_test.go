package main

import (
	sortSocks "github.com/christophermca/go-algo/sorting-socks"
	util "github.com/christophermca/go-algo/sorting-socks/util"
	"testing"
)

//this ends up being a smoke test i guess?
func TestSortingSocks(t *testing.T) {
	socks := util.GenerateSockPile(20)
	sortSocks.Run(socks)
}
