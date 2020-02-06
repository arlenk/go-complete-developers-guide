package main

import "fmt"

type deck []string
type card string

func newDeck() deck {
	suites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"one", "two", "three", "Ace", "King"}
	cards := deck{}
	for _, suite := range suites {
		for _, value := range values {
			cards = append(cards, suite + value)
			
		}
	}
	return cards;
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
