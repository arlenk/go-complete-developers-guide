package main

import (
	"fmt" 
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

type deck []string
type card string

func newDeck() deck {
	suites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"one", "two", "three", "Ace", "King"}
	cards := deck{}
	for _, suite := range suites {
		for _, value := range values {
			cards = append(cards, value + " of " + suite)
			
		}
	}
	return cards;
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)
	
	nc := len(d)
	for i := range d {
		j := r.Intn(nc - i) + i
		fmt.Println("swapping ", i, j)
		d[i], d[j] = d[j], d[i]
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

func fromFile(filename string) deck {
	bstring, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str := string(bstring)
	s := strings.Split(str, ",")
	d := deck(s)
	return d
}

	
