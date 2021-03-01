package main

import (
	"FactoryPattern/appliance"
	"fmt"
)

func main() {
	fmt.Println("Choose an appliance")
	fmt.Println("0: Microwave")
	fmt.Println("1: Stove")
	fmt.Println("2: Fridge")

	var app int
	fmt.Scan(&app)

	app1, err := appliance.GetAppliance(app)

	if err != nil {
		fmt.Println(err)
		return
	}
	app1.Start()
	app1.GetPurpose()
}
