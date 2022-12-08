package main

type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggies() []Veggies
	CreatePepperoni() Pepperoni
	CreateClam() Clam
}
