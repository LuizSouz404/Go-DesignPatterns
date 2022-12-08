package main

func main() {
	LojaSul := &Loja{}

	cliente := &ConcreteCliente{
		NameCliente: "Luiz",
	}
	LojaSul.RegisterClient(cliente)

	cliente1 := &ConcreteCliente{
		NameCliente: "Brunna",
	}
	LojaSul.RegisterClient(cliente1)

	cliente2 := &ConcreteCliente{
		NameCliente: "Pedro",
	}
	LojaSul.RegisterClient(cliente2)

	cliente3 := &ConcreteCliente{
		NameCliente: "João",
	}
	LojaSul.RegisterClient(cliente3)

	cliente4 := &ConcreteCliente{
		NameCliente: "Carol",
	}
	LojaSul.RegisterClient(cliente4)

	cupom := CupomStruct{
		IdCupom:   "123456",
		ValidDays: 4,
	}

	LojaSul.NotifyPromotion(cupom)

	cliente.UsingCupom()
}
