package main

import (
	"fmt"
)

//Animal interface
type Animal interface {
	eat(i string)
	sleep(i string)
}

//Dog is Animal
type Dog interface {
	eat(i string)
	sleep(time string)
}

//Rat is Animal
type Rat struct {
	name string
}

type con struct {
	name string
}

func printAnimal(animal Animal) {
	animal.eat("Buscuit")
}

func (con *con) eat(dish string) {
	fmt.Println(con.name, "ate ", dish)
}

func (con *con) sleep(time string) {
	fmt.Println(con.name, " at ", time)
}

func (rat *Rat) eat(dish string) {
	fmt.Println(rat.name, " ate ", dish)
}

func (rat *Rat) sleep(time string) {
	fmt.Println(rat.name, " slept at ", time)
}

var i interface{} = "Hello"

func main() {
	dog := con{
		name: "Tom",
	}
	rat := Rat{
		name: "Jerry",
	}

	printAnimal(&dog)
	printAnimal(&rat)
	fmt.Printf("%T", i)
}
