package main

import "design-pattern/decorator/dress"

func main() {
	canwinter := dress.PersonConstructor("canwinter")

	tx := dress.TShirtConstructor(canwinter)
	xz := dress.SuitConstructor(tx)
	ld := dress.TieConstructor(xz)
	xie := dress.ShoesConstructor(ld)

	xie.Show()
}
