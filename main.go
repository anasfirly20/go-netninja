package main

import "fmt"

func updateName(n string) {
	n = "wedge"
}

func main() {
	name := "tifa"

	fmt.Println("memory address of name is:", &name)
	
	m := &name

	fmt.Println("memory address:", m)
	fmt.Println("value at memory address", *m)

	updateName(name)

	fmt.Println("ori name >", name)
	
}



