package main

func main() {
	cards := newDeck()
	// hand, remainignCards := deal(cards, 5)
	cards.suffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
