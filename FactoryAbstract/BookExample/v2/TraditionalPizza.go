package main

import "fmt"

type TraditionalPizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (t *TraditionalPizza) Prepare() {
	fmt.Println("Preparing " + t.Name)
	fmt.Println("Preparing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings...")
	for _, v := range t.Toppings {
		fmt.Println(v)
	}
}
func (t *TraditionalPizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350º")
}
func (t *TraditionalPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (t *TraditionalPizza) Box() {
	fmt.Println("Place pizza " + t.Name + " in official PizzaStore box")
}
