package appliance

import "fmt"

type Fridge struct {
	name string
}

func (fridge *Fridge) Start() {
	fridge.name = "Fridge"
}

func (fridge *Fridge) GetPurpose() {
	fmt.Println("Hi my name is ", fridge.name, "I cool stuff")
}
