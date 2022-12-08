package main

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
