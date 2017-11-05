package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck type
type deck []string

// newDeck generates new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearths", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cardName := value + " of " + suit
			cards = append(cards, cardName)
		}
	}

	return cards
}

// Print prints all cards in deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Deals cards and returns hand and remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save deck to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read deck from file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// Shuffles a deck of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Random numbers using unixNano
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
