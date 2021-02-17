package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.saveToFile("new_deck")
	newCards := newDeckFromFile("new_deck")
	newCards.print()
	hand, remainingDeck := deal(newCards, 5)
	hand.print()
	remainingDeck.print()
}
