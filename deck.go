// this file should define our custom type 'deck' which is a slice of strings
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// prints contents of current deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// creates new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// cuts deck into separate decks of handSize and remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// type converter helper func
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// this function will save a list of cards to a file on the local machine
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// this func will create a list of cards from a file on the local machine
func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	strData := strings.Split(string(data), ",")
	if err != nil {
		fmt.Println("There was an error reading the file", err)
		os.Exit(1)
	}
	fmt.Println("Contents of the file:", string(data))
	return deck(strData)
}

// randomly shuffles the deck of cards by using the current time as random seed
func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d
}
