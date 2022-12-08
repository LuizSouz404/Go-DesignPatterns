package main

import "fmt"

type TruckTransport struct {
}

func (t *TruckTransport) Deliver(message string) {
	fmt.Println("Truck: " + message)
}
