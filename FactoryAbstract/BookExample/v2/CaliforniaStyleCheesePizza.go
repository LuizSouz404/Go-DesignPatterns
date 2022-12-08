package main

import "fmt"

type CaliforniaStyleCheesePizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (c *CaliforniaStyleCheesePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	fmt.Println("Preparing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings...")
	for i := range c.Toppings {
		fmt.Printf(c.Toppings[i] + " ")
	}
}
func (c *CaliforniaStyleCheesePizza) Bake() {
	fmt.Println("\nBake for 25 minutes at 350º")
}
func (c *CaliforniaStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (c *CaliforniaStyleCheesePizza) Box() {
	fmt.Println("Place pizza " + c.Name + " in official PizzaStore box")
}
