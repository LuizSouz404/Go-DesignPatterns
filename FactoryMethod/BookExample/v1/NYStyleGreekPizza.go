package main

import "fmt"

type NYStyleGreekPizza struct {
	Name              string
	Dough             Dough
	Sauce             Sauce
	Cheese            Cheese
	Veggies           []Veggies
	Clam              Clam
	IngredientFactory NYPizzaIngredientFactory
}

func (g *NYStyleGreekPizza) Prepare() {
	fmt.Println("Preparing " + g.Name)
	fmt.Println("Preparing dough...")
	g.Dough = g.IngredientFactory.CreateDough()
	fmt.Println("Adding sauce...")
	g.Sauce = g.IngredientFactory.CreateSauce()
	fmt.Println("Adding toppings...")
	g.Cheese = g.IngredientFactory.CreateCheese()
	fmt.Println("Adding veggies...")
	g.Veggies = g.IngredientFactory.CreateVeggies()
	if len(g.Veggies) > 0 {
		for i := range g.Veggies {
			fmt.Printf("%s ", g.Veggies[i])
		}
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
