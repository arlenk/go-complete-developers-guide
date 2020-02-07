package main

import (
	"fmt" 
	"strings"
	"io/ioutil"
)

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

func (d deck) toString() string {
	s := strings.Join([]string (d), ",")
	return s
}

func (d deck) toFile(filename string) error {
	res := ioutil.WriteFile(filename,
		[]byte(d.toString()),
		0666)
	return res
}

