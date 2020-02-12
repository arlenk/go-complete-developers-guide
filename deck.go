package main

import (
	"fmt" 
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

type card struct {
	suite string
	value string
}
type deck []card

func newDeck() deck {
	suites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"one", "two", "three", "Ace", "King"}
	cards := deck{}
	for _, suite := range suites {
		for _, value := range values {
			c := card{suite: suite, value: value}
			cards = append(cards, c)
		}
	}
	return cards;
}


func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d %+v\n", i, card)
	}
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)
	
	nc := len(d)
	for i := range d {
		j := r.Intn(nc - i) + i
		d[i], d[j] = d[j], d[i]
	}
}

func (d deck) toString() string {
	var b strings.Builder
	b.Grow(10)
	for i, c := range d {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(c.suite + "|" + c.value)
		
	}

	return b.String()
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
	cards := deck{}
	str := string(bstring)
	s := strings.Split(str, ",")
	for _, t := range(s) {
		s_v := strings.Split(t, "|")
		fmt.Println(s_v)
		c := card{suite: s_v[0], value: s_v[1]}
		cards = append(cards, c)
	}

	return cards
}

	
