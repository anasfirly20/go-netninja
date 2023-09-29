package main

import "fmt"

func main() {

	// ARRAYS
	// var ages [3]int = [3]int{35, 40, 45}
	var ages = [3]int{35, 40, 45}
	fmt.Println(ages, len(ages))
	
	names := [4]string{"Muhammad", "Thomi", "Anas", "Firly"}
	fmt.Println(names)
	

	// SLICES
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// SLICE RANGES
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	
	fmt.Println("rangeOne -->" , rangeOne)
	fmt.Println("rangeTwo -->",rangeTwo)
	fmt.Println("rangeThree -->",rangeThree)

}


