package main

import (
		"fmt"
		"strings"
		"io/ioutil"
		"os"
		"math/rand"
)

type Deck []string

func (this Deck) print(){	//reciever	this is obj of Deck which calls print()

	for index, data := range this{
		fmt.Println(index, data)
	}

}


func createNewDeck() Deck {

	cardSuits := []string{ "Spade", "Diamond", "Club", "Heart" }
	cardValues := []string{ "Ace", "One", "Two", "Three", "Four" }
	cards := Deck{}

	for _, suit := range cardSuits{		//_ means that we know there is a var to be used there but we are not going to use it
		for _, x := range cardValues{	//eg. index replaced by _ bcoz index will not be used and this will avoid error
			var temp = suit + " of " + x
			cards = append(cards, temp)
		}
	}

	return cards

}


func deal(this Deck, size int) (Deck, Deck){

	return this[:size], this[size:]

}

func (this Deck) deckToString() string {		//reciever to convert deck to string

	temp := []string(this)		//[]string(this) converts this/cards to array of string
	str := strings.Join(temp,",")
	return str

}

func (this Deck) saveToFile(name string) error {	//reciever to save file

	return ioutil.WriteFile(name, []byte(this.deckToString()), 0666)	//filename,byte slice of data to write,permission(r/w is 0666)

}	//return used bcoz if err then error is returned

func readFromFile(filename string) (Deck) {
	byteData, err := ioutil.ReadFile(filename)	//returns data in byte[] and error
	if err!=nil {		//i.e. if err then exits
		fmt.Println("error", err)
		os.Exit(1)
	}
	temp :=  string(byteData)	//byte[] to string
	sliceOfString := strings.Split(temp, ",")	//string to slice of str
	return Deck(sliceOfString) //slice of str to Deck
}

func (this Deck) shuffle(){

	for i := range this{

		randpos = rand.Intn(len(this)-1)	//get rand int
		this[randpos], this[i] = this[i], this[randpos]

	}

}

func main() {
	
	cards := createNewDeck()

	/*cards.print()

	handsOfCard, remainingCards := deal(cards, 5)	//returns first 5 cards and remaisning cards

	fmt.Println("First 5 cards are")
	handsOfCard.print()	//bcoz handsOfCard is of Deck type

	fmt.Println("Apart from first 5 cards remaining cards are")
	remainingCards.print()	//bcoz remainingCards is of Deck type*/

	fmt.Println(cards.deckToString())
	cards.saveToFile("Tp")
	New_Cards := readFromFile("Tp")
	New_Cards.print()

	cards.shuffle()

}