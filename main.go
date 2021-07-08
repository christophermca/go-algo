package main

import (
	sortSocks "github.com/christophermca/go-algo/sorting-socks"
	util "github.com/christophermca/go-algo/sorting-socks/util"
)

func main() {
	sortSocks.Run(util.GenerateSockPile(20))
}
