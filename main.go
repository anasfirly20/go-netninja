package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello there friends"


	// fmt.Println(strings.Contains(greeting, "Hello!"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	// // original value is unchanged
	// fmt.Println("original string value =", greeting)


	// // 
	// ages := []int{10, 12, 23, 40, 32, 11, 50, 44, 23}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 0)
	// fmt.Println(index)
	
	
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "peach"))
	
}