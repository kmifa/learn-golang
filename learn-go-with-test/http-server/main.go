package main

import (
	"net/http"

	"log"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	log.Fatal(http.ListenAndServe(":5001", server))
}
