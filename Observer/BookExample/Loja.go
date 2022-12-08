package main

type Loja struct {
	clientes []Cliente
}

func (l *Loja) RegisterClient(cliente Cliente) {
	l.clientes = append(l.clientes, cliente)
}

func (l *Loja) UnregisterClient(cliente Cliente) {
	var newArrayClientes []Cliente

	for _, oldCliente := range l.clientes {
		if cliente != oldCliente {
			newArrayClientes = append(newArrayClientes, cliente)
		}
	}
	l.clientes = newArrayClientes
}

func (l *Loja) NotifyPromotion(cupom CupomStruct) {
	for _, v := range l.clientes {
		v.Update(cupom)
	}
}
