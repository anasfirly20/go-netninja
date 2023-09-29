package main

import "fmt"

func main() {

	// // WHILE LOOP
	// x := 0
	// for x <= 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// // FOR LOOP
	// for i := 0; i <= 5 ; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// Iterating a slice
	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++{
		fmt.Println("NAME ->>", i, names[i])
	}
}