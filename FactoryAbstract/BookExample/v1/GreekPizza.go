package main

import "fmt"

type GreekPizza struct{}

func (g *GreekPizza) Prepare() {

}
func (g *GreekPizza) Bake() {

}
func (g *GreekPizza) Cut() {

}
func (g *GreekPizza) Box() {
	fmt.Println("Pizza Grega Encaixotada :P")
}
