package main

import ("fmt")

type Dict []string

func createDict() (Dict) {	//returns list of strings

	list := Dict{}
	defaultWords := []string{"man", "pan", "lamp", "fan", "gum"}
	for _, data := range defaultWords{	//copied contents
		list = append(list, data)
	}
	return list

}

func (this Dict) print(){	//reciever of Dict

	for _,data := range this{	//copied contents
		fmt.Println(data)
	}

}

func main() {
	
	d := createDict()
	d.print()

}