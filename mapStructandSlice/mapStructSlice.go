package main

import "fmt"

//Customer of
type Customer struct {
	name    string
	address string
	age     int
	work    string
}

func (cm *Customer) setName(name string) {
	(*cm).name = name
}

func (cm *Customer) getName() string {
	return cm.name
}

func change(arr []int) {
	arr[0] = 10
}
func main() {
	cm := Customer{"Haja", "India", 23, "SD"}
	fmt.Println(cm.address)

	cst := Customer{
		name:    "HAJA",
		address: "TN",
	}
	fmt.Println(cst.address)
	cst.address = "KUM"
	fmt.Println(cst.address)
	cust := &cm
	fmt.Println(cust.name)
	cust.name = "Asmath"
	fmt.Println(cm.name)
	fmt.Println(cm.address)

	var crew []Customer

	crew = append(crew, cm, cst)

	for _, v := range crew {
		fmt.Println(v.address)
	}

	mymap := map[int]string{10: "hello"}

	mymap[20] = "hi"
	fmt.Println(mymap)
	var myMap = make(map[int]string)

	myMap[10] = "hello"

	fmt.Println(myMap)

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	_, present := mymap[10]
	fmt.Println(present)

	cm.setName("Arnold")

	fmt.Println(cm.name)
	fmt.Println(cm.getName())
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr[0])
	change(arr)
	fmt.Println(arr[0])
	slice := make([]int, 4, 5)
	slice[1] = 1
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	dst := make([]int, 5)
	copy(dst, arr)
	fmt.Println(dst)
	//s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//s2 := []int{6, 7, 8, 9, 10}
	// i := 0
	// j := 3
	// copy(s1[i:], s1[j:])
	// for k, n := len(s1)-j+i, len(s1); k < n; k++ {
	// 	s1[k] = 0
	// }
	// s1 = s1[:len(s1)-j+i]
	// fmt.Println(s1)

	testMap := make(map[int]int)
	for i := 0; i < 20; i++ {
		testMap[i] = 10 * i
	}

	for _, v := range testMap {
		fmt.Println(v)
	}

}
