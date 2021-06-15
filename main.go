package main

func main() {
	cards := newDeck()

	hand, reaminingDeck := deal(cards, 5)

	hand.print()
	reaminingDeck.print()
}
