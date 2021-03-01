package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type mystruct struct {
	field1 int
	Field2 string `desc:"gopher"`
	field3 float32
}

func (mystr mystruct) TestFunction(i int) {
	mystr.field1 = i
	fmt.Println(mystr.field1)
}
func main() {

	mys := &mystruct{
		field1: 1234,
		Field2: "haja",
		field3: 1.234,
	}
	mysRtype1 := reflect.TypeOf(mys)
	mysRvalue1 := reflect.ValueOf(mys)
	fmt.Println(mysRtype1.Kind())

	val := mysRvalue1.MethodByName("TestFunction")

	vv := 12341234

	val.Call([]reflect.Value{reflect.ValueOf(vv)})
	fmt.Println(mysRtype1)
	fmt.Println(mysRvalue1)

	mysRvalue := mysRvalue1.Elem()
	mysRtype := mysRvalue.Type()
	mysRvalue.Field(1).Set(reflect.ValueOf("Asmath"))
	for i := 0; i < mysRtype.NumField(); i++ {
		mysRfieldType := mysRtype.Field(i)
		mysRfieldValue := mysRvalue.Field(i)
		fmt.Println(mysRfieldType)
		fmt.Println(mysRfieldValue)
	}

	var r io.Reader
	r = strings.NewReader("ajfhaj")
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.ValueOf(r))
	tty, _ := os.Open("configfile.conf")
	fmt.Println(reflect.TypeOf(tty))
	fmt.Println(reflect.ValueOf(tty))
	r = tty
	var w io.Writer
	w = r.(io.Writer)
	fmt.Println(reflect.TypeOf(w))
	fmt.Println(reflect.ValueOf(w))
}
