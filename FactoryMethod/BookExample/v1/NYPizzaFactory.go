package main

type NYPizzaFactory struct{}

func (s *NYPizzaFactory) CreatePizza(typePizza string) Pizza {
	var pizza Pizza

	if typePizza == "cheese" {
		pizza = &NYStyleCheesePizza{
			Name:              "NY style sauce and cheese pizza",
			IngredientFactory: NYPizzaIngredientFactory{},
		}
	} else if typePizza == "greek" {
		pizza = &NYStyleGreekPizza{
			Name:              "NY style mozzarella and chopped pizza",
			IngredientFactory: NYPizzaIngredientFactory{},
		}
	} else {
		pizza = &TraditionalPizza{}
	}

	return pizza
}
