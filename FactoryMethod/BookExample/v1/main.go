package main

import "fmt"

func main() {
	NYPizzaStore := PizzaStore{
		Factory: &NYPizzaFactory{},
	}
	NYPizzaStore.OrderPizza("cheese")
	fmt.Println("----------------------------------------------")

	NYPizzaStore.OrderPizza("greek")
}
