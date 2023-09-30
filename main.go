package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup" : 4.99,
		"pie" : 7.99,
		"salad" : 6.99,
		"toffee pudding" : 3.55,

	}

	// fmt.Println(menu)
	// fmt.Println(menu["soup"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k,"-",v)
	}
	

	// ints as key type
	phonebook := map[int]string{
		12312312: "mario",
		34343434: "luigi",
		56565656: "peach",
	}
	
	fmt.Println(phonebook)
	// fmt.Println(phonebook[12312312])

	phonebook[12312312] = "anas"
	fmt.Println(phonebook)
	
	phonebook[56565656] = "tomi"
	fmt.Println(phonebook)

}
