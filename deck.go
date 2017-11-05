package cards

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
type Deck []string

// NewDeck generates new deck of cards
func NewDeck() Deck {
	cards := Deck{}

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
func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Deal deals cards and returns hand and remaining cards
func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// ToString converts deck to string
func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

// SaveToFile saves deck to file
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// NewDeckFromFile reads deck from file
func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error:", err)
		os.Exit(1)
	}

	return Deck(strings.Split(string(bs), ","))
}

// Shuffle shuffles a deck of cards
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Random numbers using unixNano
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
