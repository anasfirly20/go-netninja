package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// Make new bill
func newBill(name string) bill {
	b := bill {
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	fmt.Println("ADD name >",name )
	fmt.Println("ADD PRICE >",price )
	b.items[name] = price
	fmt.Println("ADD CHECK b >",b )
}

// update tip
func (b *bill) updateTip(tip float64){
	(*b).tip = tip
}

// save bill
func (b *bill) save(){
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
