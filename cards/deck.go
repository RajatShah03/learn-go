package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	
	cardSuites := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _,suite := range cardSuites {
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) show() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs,err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\n")
	
	return deck(s)
}

func (d deck) shuffle() {
	// Creating a random seed and using that to generate random index using time
	// Optional exercise but this version of rand.Intn equally generates randomized numbers
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		ridx := r.Intn(len(d) - 1)
		d[i], d[ridx] = d[ridx], d[i]
	}
}