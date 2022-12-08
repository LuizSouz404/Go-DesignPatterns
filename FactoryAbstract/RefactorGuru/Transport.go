package main

type Transport interface {
	Deliver(message string)
}
