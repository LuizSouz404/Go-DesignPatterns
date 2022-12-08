package main

func main() {
	logistic := Logistics{}

	transport, err := logistic.CreateTransport("truck")

	if err != nil {
		panic(err.Error())
	}

	transport.Deliver("Entregando")

	transport, err = logistic.CreateTransport("sea")

	if err != nil {
		panic(err.Error())
	}

	transport.Deliver("Entregando")
}
