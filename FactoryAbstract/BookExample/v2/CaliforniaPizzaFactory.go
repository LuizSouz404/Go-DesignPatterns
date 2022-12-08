package main

type CaliforniaPizzaFactory struct{}

func (s *CaliforniaPizzaFactory) CreatePizza(typePizza string) Pizza {
	var pizza Pizza

	if typePizza == "cheese" {
		pizza = &CaliforniaStyleCheesePizza{
			Name:  "NY style sauce and cheese pizza",
			Dough: "Extra Thick Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Grated", "Reggiano", "Cheese",
			},
		}
	} else if typePizza == "greek" {
		pizza = &CaliforniaStyleGreekPizza{
			Name:  "California style mozzarella and chopped pizza",
			Dough: "Extra Thick Crust Dough",
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
