package main

func main() {

	deck := newDeckFromFile("cards")
	deck.print()
}

func newCard() string {
	return "five of diamonds"
}
