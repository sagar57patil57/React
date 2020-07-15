package main

import "fmt"

func main() {

	var x int8 = 23
	var y uint32 = 49908
	var flag bool = true
	str := "place"
	fmt.Println(x, y, flag, len(str))

	//--------------------------------------------------------------------------
	//declarations
	var p string	//works
	var q = "Hello"	//works
	//var r			//fails
	//i.e. atleast mention value or either data type

	//--------------------------------------------------------------------------
	var x3 int 		//initiallizes with 0
	fmt.Println(p, q, x3)

	//--------------------------------------------------------------------------
	a, b, c := 1,2,"yup"
	fmt.Println(a,b,c)

	//--------------------------------------------------------------------------
	const PIE = 21
	fmt.Println(PIE)
}