package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	mybill := createBill()
	promptOptions(mybill)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n') 

	return strings.TrimSpace(input), error
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options (a -add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item price:", reader)
		fmt.Println(name, price)

	case "t":
		tip, _ := getInput("Enter tip amount ($):", reader)
		fmt.Println(tip)

	case "s":
		fmt.Println("You chose to save the bill")
	
	default: 
	fmt.Println("That was not a valid option..")
	promptOptions(b)
	}
}