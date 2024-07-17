package main

import (
	"fmt"
)

func category() {

	catalog := []string{"1. Action", "2. Adventure", "3. Isekai", "4. Hentai", "5. Romantic"}
	var selection int

	printCatalog(catalog)

	fmt.Printf("\n")
	fmt.Println("Please select any one of the categories ")
	fmt.Scan((&selection))
	fmt.Printf("\n")

	switch selection {
	case 1:
		action()
	case 2:
		adventure()
	case 3:
		isekai()
	case 4:
		hentai()
	case 5:
		romance()
	}
}

func printCatalog(catalog []string) {
	for i := 0; i < len(catalog); i++ {
		fmt.Println(catalog[i])
	}
}
