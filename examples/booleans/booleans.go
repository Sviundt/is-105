package main

import "fmt"

func main() {
	fmt.Println(true && true) // Ignorer feilmeldinger i Golang her, ettersom dette er et eksempel.
	// True og True vil returnere True

	fmt.Println(true && false)
	// True og False vil returnere False

	fmt.Println(true || true) // Ignorer feilmeldinger i Golang her, ettersom dette er et eksempel.
	// True eller True vil returnere True

	fmt.Println(true || false)
	// True eller False vil returnere True

	fmt.Println(!true)
	// !True vil returnere False
}
