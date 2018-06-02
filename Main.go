package main

import (
	"fmt"
)

func main() {
	var _name string = "Melanie"
	_nameTwo := "Sharon"

	var (
		_nameThree string
	)
	_nameThree = "Veliz Ruiz"

	fmt.Println(_name, _nameTwo, _nameThree)
	fmt.Println("---------------------------")

	var (
		_valeInt     int
		_valueFloat  float64
		_valueString string
		_valueBool   bool
	)

	_valeInt = 20
	_valueFloat = 50.456
	_valueBool = false
	_valueString = "Leslie"

	fmt.Println(_valeInt, _valueFloat, _valueBool, _valueString)
	lista()
	fmt.Println(nombre())
}

func lista() {
	for index := 0; index < 10; index++ {
		fmt.Println("Numero", index)
	}
}

func nombre() string {
	return "Katty"

}
