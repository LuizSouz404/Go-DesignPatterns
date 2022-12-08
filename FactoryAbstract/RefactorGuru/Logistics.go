package main

import "errors"

type Logistics struct{}

func (l *Logistics) PlanDelivery() {

}

func (l *Logistics) CreateTransport(typeTransport string) (Transport, error) {
	if typeTransport == "truck" {
		return &TruckTransport{}, nil
	} else if typeTransport == "sea" {
		return &SeaTransport{}, nil
	}
	return nil, errors.New("Wrong trnasport type passed")
}
