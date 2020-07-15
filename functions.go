package main

import "fmt"

func main() {
	funOne()
	card := funTwo()
	fi, se := funThree()
	fmt.Println(card, fi, se)
}

func funOne() {
	fmt.Println("Function without return")
}

func funTwo() string {				//returns 1 value
	return "Returns a string"
}

func funThree() (string, int) {		//returns mutiple values
	return "First Return", 90
}