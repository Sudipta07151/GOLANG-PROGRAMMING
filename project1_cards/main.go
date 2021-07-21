package main

import "fmt"



func main() {
	// var card string = "Ace of Spades"
	// card:="Ace of Spades"
	// card:=newCard()
	// fmt.Println(card)
	// cards := []string{newCard(), "Ace of Spades"}
	// cards := deck{newCard(), "Ace of Spades"}
	// cards = append(cards, "Six of Spades")
	
	cards:=newDeck()
	
	fmt.Println(cards)

	// for i,card:=range cards{
	// 	fmt.Println(i,card)
	// }
	hand,remain:=deal(cards,3)
	fmt.Println("ALL CARD VALUES");
	cards.print()
	fmt.Println("HAND VALUES");
	hand.print();
	fmt.Println("REAMINING VALUES");
	remain.print();

	fmt.Println([]byte("Hi there"))
	fmt.Println(cards.decktoString())
	cards.saveToFile("deckRes")
	readCards:=getDeckFromFile("deckRes")
	readCards.print()
	cards.shuffle()
	fmt.Println("\nAFTER SHUFFLE")
	cards.print()
}


// func newCard() string {
// 	return "Five of Diamonds"
// }


