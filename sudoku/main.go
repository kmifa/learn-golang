package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	for key, value := range "hello" {
		fmt.Println(key, value)
	}
}
