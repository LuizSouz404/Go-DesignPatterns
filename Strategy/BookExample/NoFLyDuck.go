package main

import "fmt"

type NoFlyDuck struct{}

func (n *NoFlyDuck) Fly() {
	fmt.Println("Não voo")
}
