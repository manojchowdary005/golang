package main

import (
	"fmt"
)

func main() {
	var visitorName string
	fmt.Println("Please enter your first name")
	fmt.Scan(&visitorName)

	greet(visitorName)
	category()
	sayBye(visitorName)
}
