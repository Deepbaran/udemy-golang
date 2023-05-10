package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	//We are not making it as a function with receiver.
	//As making it function with receiver will seem like,
	//we are trying to modify the value of the existing variable by calling "cards.deal(5)"
	//and now cards will have 5 less cards in the slice. Which is not the case.
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		panic(err)
	}

	s := strings.Split(string(bs), ",")
	return deck(s) // This conversion is possible, only because deck is an extension of []string
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //time.Now().UnixNano() this seed value makes sure that the seed passed to the source for rand stays random
	r := rand.New(source)                           //Creating a new rand with our custom seed

	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1) //Our randomly seeded rand is used here

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
