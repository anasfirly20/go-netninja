package main

import "fmt"

func main(){
	mybill := newBill("mario's bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("coffee", 3.25)
	
	mybill.updateTip(10.5)
	
	fmt.Println(mybill.format())
}