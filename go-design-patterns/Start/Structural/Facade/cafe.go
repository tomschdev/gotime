package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n--------------------")
	// make an americano coffee using the coffeemachine API
	// determine beans amount to use - 5oz for every 8oz size
	americano := &CoffeeMachine{}
	amtBeans := (size / 8.0) * 5.0
	americano.startCoffee(amtBeans, 3)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()
	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n--------------------")
	// make a latte coffee using the coffeemachine API
	// determine beans amount to use - 5oz for every 8oz size
	// determine milk amount to use - 2oz for every 8oz size
	latte := &CoffeeMachine{}
	amtBeans := (size / 8.0) * 5.0
	latte.startCoffee(amtBeans, 3)
	latte.grindBeans()
	latte.useHotWater(size)
	amtMilk := (size / 8.0) * 2.0
	latte.useMilk(amtMilk)
	latte.doFoam(true)
	latte.endCoffee()
	fmt.Println("Latte is ready!")
}
