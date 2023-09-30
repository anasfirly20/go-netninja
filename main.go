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
	
}
