package main

type FactoryPizza interface {
	CreatePizza(typePizza string) Pizza
}
