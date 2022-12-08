package main

import "fmt"

type CheesePizza struct{}

func (c *CheesePizza) Prepare() {

}
func (c *CheesePizza) Bake() {

}
func (c *CheesePizza) Cut() {

}
func (c *CheesePizza) Box() {
	fmt.Println("Pizza de Queijo Encaixotada :P")
}
