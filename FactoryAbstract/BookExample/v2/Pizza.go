package main

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
