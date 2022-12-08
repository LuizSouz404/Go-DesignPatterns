package main

func main() {
	pizzaStore := PizzaStore{
		Factory: SimplePizzaFactory{},
	}
	pizzaCheese := pizzaStore.OrderPizza("cheese")

	pizzaCheese.Box()

	pizzaGreek := pizzaStore.OrderPizza("greek")

	pizzaGreek.Box()
}
