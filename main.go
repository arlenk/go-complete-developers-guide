package main

import "fmt"

func main() {
	var cards deck = newDeck()
	fmt.Println(cards)
	cards.print()
}
