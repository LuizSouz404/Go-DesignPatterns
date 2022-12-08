package main

import "fmt"

type NYStyleGreekPizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (g *NYStyleGreekPizza) Prepare() {
	fmt.Println("Preparing " + g.Name)
	fmt.Println("Preparing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings...")
	for i := range g.Toppings {
		fmt.Printf(g.Toppings[i] + " ")
	}
}
func (g *NYStyleGreekPizza) Bake() {
	fmt.Println("\nBake for 25 minutes at 350º")
}
func (g *NYStyleGreekPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (g *NYStyleGreekPizza) Box() {
	fmt.Println("Place pizza " + g.Name + " in official PizzaStore box")
}
