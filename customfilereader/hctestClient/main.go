package main

import (
	"customfilereader"
	"fmt"
)

type mystruct struct {
	TestString string  `name:"testString",json="testString"`
	TestBool   bool    `name:"testBool",json="testBool"`
	TestFloat  float64 `name:"testFloat",json="testFloat"`
	TestInt    int64   `name:"TestInt",json="TestInt"`
}

func main() {
	mystruct1 := new(mystruct)
	customfilereader.ReadFile("configfile.conf", mystruct1, "CUSTOM")
	fmt.Println(mystruct1)

	mystruct2 := new(mystruct)

	customfilereader.ReadFile("jsonConfig.json", mystruct2, "JSON")

	fmt.Println(mystruct2)
}
