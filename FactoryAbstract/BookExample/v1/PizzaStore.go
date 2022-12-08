package main

type PizzaStore struct {
	Factory SimplePizzaFactory
}

func (p *PizzaStore) OrderPizza(typePizza string) IPizza {
	pizza := p.Factory.CreatePizza(typePizza)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
