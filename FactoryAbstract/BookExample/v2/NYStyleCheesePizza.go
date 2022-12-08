package main

import "fmt"

type NYStyleCheesePizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (c *NYStyleCheesePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	fmt.Println("Preparing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings...")
	for i := range c.Toppings {
		fmt.Printf(c.Toppings[i] + " ")
	}
}
func (c *NYStyleCheesePizza) Bake() {
	fmt.Println("\nBake for 25 minutes at 350º")
}
func (c *NYStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (c *NYStyleCheesePizza) Box() {
	fmt.Println("Place pizza " + c.Name + " in official PizzaStore box")
}
