package main

func main() {
	//var card string = "Ace of spades"
	// var card = "Ace of spades"

	// card := "Ace of spades"
	// fmt.Println(card)
	// card = newCard()

	// cards := []string{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// 	fmt.Println(card)
	// }

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
