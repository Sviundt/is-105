package main

import "fmt"

func main() {
	fmt.Println(len("Hallo, verden!"))
	// Finner lengden av "Hallo, verden!"

	fmt.Println("Hallo, verden!"[1])
	// Finner den andre verdien i "Hallo, verden!" siden vi teller fra 0.

	fmt.Println("Hallo, " + "verden!")
	// Legger sammen "Hallo, " og "verden!" til en streng.
}
