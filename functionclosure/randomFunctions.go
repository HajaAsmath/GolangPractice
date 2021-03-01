package main

import "fmt"

func main() {
	addSubMulippp := addSubMul
	fmt.Println(addSubMulippp(12, 5))
	fmt.Println(addSubMulWithJustReturn(10, 10))
	greet := func(name string) string {
		return "Hello " + name + ",Have a nice day"
	}

	fmt.Println(greet("Haja"))

	fmt.Println(doSomething("1", greet))

	inc := incrementor()

	fmt.Println(inc(), inc())

	fmt.Println(inc())

	res := int(9 / 2)

	fmt.Println(res)
}

func addSubMul(a int, b int) (int, int, int) {
	add := a + b
	sub := a - b
	mul := a * b
	return add, sub, mul
}

func addSubMulWithJustReturn(a, b int) (add, sub, mul int) {
	add = a + b
	sub = a - b
	mul = a * b
	return
}

func doSomething(a string, fn func(name string) string) string {
	return a + "." + fn("Haja")
}

func incrementor() func() int {
	i := 0
	fmt.Println("Incrementer is called")
	return func() int {
		i++
		fmt.Println("Incremented")
		return i
	}
}
