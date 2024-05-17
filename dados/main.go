package main

import "fmt"

func main() {

	//dados
	//strings

	var name string = "Um"
	var nameDois = "tres"
	var nameTres string
	fmt.Println(name, nameDois, nameTres)

	name = "main"
	nameTres = "tres"

	fmt.Println(name, nameDois, nameTres)

	namerex := "rename"

	fmt.Println(namerex)

	// ints

	var numberUm int = 23
	var numberdois = 12
	number := 67
	fmt.Println(numberUm, numberdois, number)

	// bits & memory

	var num1 int8 = 2
	var num2 int16 = 23
	num3 := 28
	var num4 uint = 89
	fmt.Println(num1, num2, num3, num4)

	//Float

	var val1 float64 = 1.3
	var val2 float32 = -1.5

	fmt.Println(val1, val2)

}
