package main

import "fmt"

type FlyableDuck struct{}

func (f *FlyableDuck) Fly() {
	fmt.Println("Voando")
}
