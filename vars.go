package main

import "fmt"

func main() {
	var x int = 32		//WAY 1 to declare variable

	card := "7 Spade"	//:= only used for init		//WAY 2 to declare variable
	card = "6 Club"		//changes value
	//card := "9 Heart"	//error cannot reinit card

	fmt.Println(card, x)
}