package main

import "fmt"

type CarRoute struct{}

func (c *CarRoute) BuildRoute(A, B string) {
	fmt.Printf("Car Route: Ponto de partida %s -> Ponto de chegada %s\n", A, B)
}
