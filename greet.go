package main

import (
	"fmt"
)

func greet(visitorName string) {
	fmt.Printf("%52s %s\n", "Konnichiwa ", visitorName)
	fmt.Printf("%72s\n", " Welcome to Otaku World ")
	fmt.Println()
}

func sayBye(visitorName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Println(" Oniiicha - Visit again for the new feed! ", visitorName)
}
