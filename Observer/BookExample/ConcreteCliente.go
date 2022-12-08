package main

import "fmt"

type CupomStruct struct {
	IdCupom   string
	ValidDays int
}

type ConcreteCliente struct {
	NameCliente string
	Cupom       *CupomStruct
}

func (c *ConcreteCliente) Update(cupom CupomStruct) {
	fmt.Printf("Olá %s, você recebeu um cupom valido por %d dias\n", c.NameCliente, cupom.ValidDays)
	c.Cupom = &cupom
}

func (c *ConcreteCliente) UsingCupom() {
	if c.Cupom == nil {
		fmt.Println("Você não tem cupom")
		return
	}
	fmt.Printf("%s: Você usou o Cupom com id %s\n", c.NameCliente, c.Cupom.IdCupom)
	c.Cupom = nil
}
