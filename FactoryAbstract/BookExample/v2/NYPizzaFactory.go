package main

type NYPizzaFactory struct{}

func (s *NYPizzaFactory) CreatePizza(typePizza string) Pizza {
	var pizza Pizza

	if typePizza == "cheese" {
		pizza = &NYStyleCheesePizza{
			Name:  "NY style sauce and cheese pizza",
			Dough: "Thin Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Grated", "Reggiano", "Cheese",
			},
		}
	} else if typePizza == "greek" {
		pizza = &NYStyleGreekPizza{
			Name:  "NY style mozzarella and chopped pizza",
			Dough: "Thin Crust Dough",
			Sauce: "tomatoes halved",
			Toppings: []string{
				"Oregano", "Basil", "Thyme",
			},
		}
	} else {
		pizza = &TraditionalPizza{}
	}

	return pizza
}
