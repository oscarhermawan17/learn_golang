package main

func main() {
	// tmp_cards := newDeck()
	cards := newDeckFromFile("my_cards")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())

	// tmp_cards.saveToFile("my_cards")
	// tmp_cards.print()
	// cards.print()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
