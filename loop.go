package main

import "fmt"

func main() {

	for i:=0;i<10;i++ {
		if i%2==0 {
			fmt.Printf("%d ", i)
		}
	}

	//-----------SLICES
	cards := []string{ "7 Diamond", "6 Spades" } 
	cards = append(cards, "1 Diamond")

	for index, data := cards {
		fmt.Printf(index, " ", data)
	}
	
}