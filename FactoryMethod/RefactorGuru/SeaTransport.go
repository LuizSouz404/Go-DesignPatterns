package main

import "fmt"

type SeaTransport struct {
}

func (s *SeaTransport) Deliver(message string) {
	fmt.Println("Sea: " + message)
}
