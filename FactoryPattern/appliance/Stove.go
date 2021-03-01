package appliance

import "fmt"

type Stove struct {
	name string
}

func (stove *Stove) Start() {
	stove.name = "Stove"
}

func (stove *Stove) GetPurpose() {
	fmt.Println("Hi my name is ", stove.name, "I cook")
}
