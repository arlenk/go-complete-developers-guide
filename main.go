package main

import "fmt"

func main() {
	var cards deck = newDeck()
	fmt.Println(cards.toString())
	cards.print()
	cards.toFile("out.dat")
}
