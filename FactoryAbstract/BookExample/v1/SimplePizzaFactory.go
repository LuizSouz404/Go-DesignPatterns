package main

type SimplePizzaFactory struct{}

func (s *SimplePizzaFactory) CreatePizza(typePizza string) IPizza {
	var pizza IPizza

	if typePizza == "cheese" {
		pizza = &CheesePizza{}
	} else if typePizza == "greek" {
		pizza = &GreekPizza{}
	} else {
		pizza = &TraditionalPizza{}
	}

	return pizza
}
