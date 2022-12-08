package main

type NYPizzaIngredientFactory struct{}

type Dough struct{}
type Sauce struct{}
type Cheese struct{}
type Veggies string
type Pepperoni struct{}
type Clam struct{}

func (n *NYPizzaIngredientFactory) CreateDough() Dough {
	return Dough{}
}
func (n *NYPizzaIngredientFactory) CreateSauce() Sauce {
	return Sauce{}
}
func (n *NYPizzaIngredientFactory) CreateCheese() Cheese {
	return Cheese{}
}
func (n *NYPizzaIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{"Tomatos", "Oregano", "Fried Garlic"}
}
func (n *NYPizzaIngredientFactory) CreatePepperoni() Pepperoni {
	return Pepperoni{}
}
func (n *NYPizzaIngredientFactory) CreateClam() Clam {
	return Clam{}
}
