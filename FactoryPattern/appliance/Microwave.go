package appliance

import "fmt"

type Microwave struct {
	name string
}

func (microwave *Microwave) Start() {
	microwave.name = "Microwave"
}

func (microwave *Microwave) GetPurpose() {
	fmt.Println("Hi my name is ", microwave.name, "I heat stuff")
}
