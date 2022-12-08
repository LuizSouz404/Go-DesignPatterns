package main

import "fmt"

type NYStyleCheesePizza struct {
	Name              string
	Dough             Dough
	Sauce             Sauce
	Cheese            Cheese
	Veggies           []Veggies
	Clam              Clam
	IngredientFactory NYPizzaIngredientFactory
}

func (c *NYStyleCheesePizza) Prepare() {
	fmt.Println("Preparing " + c.Name)
	fmt.Println("Preparing dough...")
	c.Dough = c.IngredientFactory.CreateDough()
	fmt.Println("Adding sauce...")
	c.Sauce = c.IngredientFactory.CreateSauce()
	fmt.Println("Adding toppings...")
	c.Cheese = c.IngredientFactory.CreateCheese()
	fmt.Println("Adding veggies...")
	c.Veggies = c.IngredientFactory.CreateVeggies()
	if len(c.Veggies) > 0 {
		for i := range c.Veggies {
			fmt.Printf("%s ", c.Veggies[i])
		}
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
