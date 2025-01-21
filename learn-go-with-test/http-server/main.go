package main

import (
	"net/http"

	"log"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	log.Fatal(http.ListenAndServe(":5001", server))
}
