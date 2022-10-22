package main

import (
	sortSocks "github.com/christophermca/go-algo/sorting-socks"
	util "github.com/christophermca/go-algo/sorting-socks/util"
)

func main() {
	sortSocks.SortingSocks(util.GenerateSockPile(20))
}
