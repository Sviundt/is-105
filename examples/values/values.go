package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	// Legger sammen to Strings "go" og "lang"

	fmt.Println("1+1 =", 1+1)
	// Legger sammen stringen "1+1 =" med integerne "`1+1`"

	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Del ett Floating-Point Number på ett annet og retuner verdien

	fmt.Println(true && false)
	// Returnerer verdien for "true and false", som er false

	fmt.Println(true || false)
	// Returner verdien for "true or false", som er true

	fmt.Println(!true)
	// Returner verdien som ikke er true (altså false)
}
