package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := cards.deal(5)

	// hand.print()
	// fmt.Println()
	// remainingCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	newDeckFromFile("my_cards").print()
}
