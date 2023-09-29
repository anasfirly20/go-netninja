package main

import "fmt"

func main() {
	// age := 45

	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// } else if age < 20 {
	// 	fmt.Println("Age is less than 20")
	// } else {
	// 	fmt.Println("Age is less than 45")
	// }

	// 
	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}
}