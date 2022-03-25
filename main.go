package main

import "fmt"

func main() {
	fmt.Println("Let's begin by newDeck()")
	cards := newDeck()
	// cards := newDeckFromFile("ms")
	cards.print()
	// cards.saveToFile("myCards")
	fmt.Println("Now we are going to shuffle the deck")
	cards.shuffle().print()
	fmt.Println("The above is the shuffled deck")
}
