package main

import (
	"decorator/decorators"
	"fmt"
)

func main() {
	beverage := &decorators.Espresso{
		Size: "big",
	}
	EspressoMocha := &decorators.Mocha{
		Beverage: beverage,
	}
	EspressoMocha = &decorators.Mocha{
		Beverage: EspressoMocha,
	}
	EspressoDoubleMochaWhip := &decorators.Whip{
		Beverage: EspressoMocha,
	}

	fmt.Printf("Seu %s Tamanho %s ficou %.2f\n", EspressoDoubleMochaWhip.GetDescription(), EspressoDoubleMochaWhip.GetSize(), EspressoDoubleMochaWhip.Cost())
}
