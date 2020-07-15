package main

type Deck []string

func (obj Deck) print(){

	for index, data := range obj{
		fmt.Println(index, data)
	}

}