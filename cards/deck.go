package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'Deck' which is a slice of String
type deck []string

// Create new deck & initialize them with Cards
func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four" /*, "Five", "Six", "Seven", "Eight", "Nine", "Ten"*/}

	for _, s := range suits {
		for _, v := range values {
			cards = append(cards, v+" of "+s)
		}
	}

	return cards
}

// Print all cards in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return two new deck one with hand and another with remaining cards.
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// Return String value of deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saves deck to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read files and create a new deck
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffles the cards in deck
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
