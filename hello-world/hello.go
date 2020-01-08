package main

// Importerer kodepakken "fmt" (format)
import "fmt"

/* Funksjonen main. Alle funksjoner defineres med "func" etterfulgt av funksjonsnavnet.
Etter funksjonsnavnet kommer en liste med null eller flere "parametere" omringet av paranteser, en valgfri return type og en "body" omringet av {}.
Denne funksjonen har ingen parametere, returnerer ingenting og har bare et "statement". Main er spesielt, fordi det er funksjonen som blir sjekket når programmet kjøres.
*/
func main() {
	// formater.print linje("akkurat som det er skrevet her"). "Hallo, verden!" er en såkalt "string"
	fmt.Println("Hallo, verden!")
}
