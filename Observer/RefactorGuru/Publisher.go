package main

import (
	"fmt"
	"reflect"
)

type Publisher struct {
	subcribers []Subscriber
	mainState  string
}

func (p *Publisher) Subscriber(s Subscriber) {
	p.subcribers = append(p.subcribers, s)
}

func (p *Publisher) Unsubscriber(s Subscriber) {
	var newArrayObserver []Subscriber
	fmt.Println("Unsubcriber")
	for i := 0; i < len(p.subcribers); i++ {
		if !reflect.DeepEqual(p.subcribers[i], s) {
			newArrayObserver = append(newArrayObserver, p.subcribers[i])
		}
	}

	p.subcribers = newArrayObserver
}

func (p *Publisher) NotifySubscribers() {
	fmt.Println("Atualizando dados")
	for i := 0; i < len(p.subcribers); i++ {
		p.subcribers[i].Update()
	}
}
