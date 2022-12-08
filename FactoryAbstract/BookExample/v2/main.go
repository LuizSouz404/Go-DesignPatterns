package main

import "fmt"

func main() {
	NYPizzaStore := PizzaStore{
		Factory: &NYPizzaFactory{},
	}
	NYPizzaStore.OrderPizza("cheese")
	NYPizzaStore.OrderPizza("greek")

	fmt.Println("----------------------------------------------")

	CaliforniaPizzaStore := PizzaStore{
		Factory: &CaliforniaPizzaFactory{},
	}
	CaliforniaPizzaStore.OrderPizza("cheese")
	CaliforniaPizzaStore.OrderPizza("greek")
}
