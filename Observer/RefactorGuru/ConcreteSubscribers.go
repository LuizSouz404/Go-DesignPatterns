package main

import "fmt"

type ConcreteSubscriber struct {
	NameSubscriber string
}

func (c *ConcreteSubscriber) Update() {
	fmt.Println("Atualizando")
	fmt.Printf("%s Atualizado\n", c.NameSubscriber)
}
