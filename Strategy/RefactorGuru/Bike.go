package main

import "fmt"

type BikeRoute struct{}

func (b *BikeRoute) BuildRoute(A, B string) {
	fmt.Printf("Bike Route: Ponto de partida %s -> Ponto de chegada %s\n", A, B)
}
