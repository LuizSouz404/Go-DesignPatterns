package main

import "fmt"

type MotoRoute struct{}

func (m *MotoRoute) BuildRoute(A, B string) {
	fmt.Printf("Moto Route: Ponto de partida %s -> Ponto de chegada %s\n", A, B)
}
