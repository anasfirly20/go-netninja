package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	fmt.Println("N >",n)

	s := strings.ToUpper(n)
	fmt.Println("S >",s)
	
	names := strings.Split(s, " ")
	fmt.Println("Names >",names)

	var initials []string
	for _, v := range names {
		fmt.Println("CHECK>>>", v[:1])
		initials = append(initials, v[:1])
	}
	
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
} 

func main() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)
	
}