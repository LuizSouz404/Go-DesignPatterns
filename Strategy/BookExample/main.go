package main

func main() {
	noFlyBehavior := &FlyableDuck{}
	duck := &Duck{
		FlyBehavior: noFlyBehavior,
	}

	duck.Fly()
}
