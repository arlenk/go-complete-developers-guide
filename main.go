package main

func main() {
	var cards deck = newDeck()
	cards.toFile("out.dat")
	cards = fromFile("out.dat")
	cards.print()
	cards.shuffle()
	cards.print()
	
}
