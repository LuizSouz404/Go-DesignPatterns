package main

import "fmt"

type TraditionalPizza struct{}

func (t *TraditionalPizza) Prepare() {

}
func (t *TraditionalPizza) Bake() {

}
func (t *TraditionalPizza) Cut() {

}
func (t *TraditionalPizza) Box() {
	fmt.Println("Pizza Tradicional Encaixotada :P")
}
