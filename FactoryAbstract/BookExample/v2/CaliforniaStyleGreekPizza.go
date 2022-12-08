package main

import "fmt"

type CaliforniaStyleGreekPizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (g *CaliforniaStyleGreekPizza) Prepare() {
	fmt.Println("Preparing " + g.Name)
	fmt.Println("Preparing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings...")
	for i := range g.Toppings {
		fmt.Printf(g.Toppings[i] + " ")
	}
}
func (g *CaliforniaStyleGreekPizza) Bake() {
	fmt.Println("\nBake for 25 minutes at 350º")
}
func (g *CaliforniaStyleGreekPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (g *CaliforniaStyleGreekPizza) Box() {
	fmt.Println("Place pizza " + g.Name + " in official PizzaStore box")
}
