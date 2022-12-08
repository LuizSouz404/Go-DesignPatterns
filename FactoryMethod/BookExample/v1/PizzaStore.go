package main

type PizzaStore struct {
	Factory FactoryPizza
}

func (p *PizzaStore) OrderPizza(typePizza string) Pizza {
	pizza := p.Factory.CreatePizza(typePizza)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
