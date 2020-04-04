// Nedprioritert til jeg finner ut hvordan golang funker
package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./site")))

	log.Fatal(http.ListenAndServe(":4040", nil))
}
