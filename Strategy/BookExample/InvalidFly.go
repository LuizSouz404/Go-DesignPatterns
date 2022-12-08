package main

import "fmt"

type InvalidFly struct{}

func (i *InvalidFly) Quack() {
	fmt.Println("QUACK")
}
