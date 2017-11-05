# Card

Cards is Go package for creating, shuffling, dealing and saving decks of cards.

## Why?

I don't know. Just learning Go :)

## Import

First get and import cards package
~~~
go get "github.com/frasza/cards"
~~~
followed by
~~~
import "github.com/frasza/cards"
~~~

## Example usage

First new deck of cards has to be generated. New deck is generated with NewDeck function:
~~~
deck := cards.NewDeck()
~~~

Printing all the cards from deck is available from Print function:
~~~
deck.Print()
~~~

Shuffling the deck is available from Shuffle function:
~~~
deck.Shuffle()
~~~

Dealing card is available with function Deal which returns hand and remaining cards from deck (5)
~~~
Deal(deck, 5)
~~~

Saving the deck is available with function SaveToFile and needs filename argument
~~~
deck.SaveToFile("myDeck.txt")
~~~

Finally, reading deck from file is available with function NewDeckFromFile and needs filename argument
~~~
newDeck := NewDeckFromFile("myDeck.txt")
~~~

