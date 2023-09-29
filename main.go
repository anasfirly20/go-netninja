package main

import "fmt"


func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string){
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func main() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	cycleNames([]string{"cloud", "mario", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "mario", "barret"}, sayBye)
	
}