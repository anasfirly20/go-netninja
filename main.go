package main

import "fmt"

func main() {
	age := 45

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 20 {
		fmt.Println("Age is less than 20")
	} else {
		fmt.Println("Age is less than 45")
	}
}