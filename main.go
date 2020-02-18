package main

import "fmt"


func main() {
	var cards deck = newDeck()
	cards.toFile("out.dat")
	cards = fromFile("out.dat")
	//cards.print()
	cards.shuffle()
	//cards.print()

	p := Player{"Albert"}
	fmt.Println(p)
	p.setName("John")
	fmt.Println(p)

	
}

