package main

import "fmt"

func main() {
	for {
		fmt.Println("Selecione o veiculo para calculo da rota: \n[1] Carro\n[2] Moto\n[3] Bike\n[4] Exit")
		var opt int
		fmt.Scanln(&opt)

		if opt == 4 {
			break
		}

		selected := SelectedVeic(opt)

		navigator := &Navigator{}
		navigator.SetRouteType(selected)

		navigator.BuildRoute("Santos", "São Vicente")
	}
}

func SelectedVeic(opt int) BuildRouteBehavior {
	switch opt {
	case 1:
		return &CarRoute{}
	case 2:
		return &MotoRoute{}
	case 3:
		return &BikeRoute{}
	}
	return nil
}
