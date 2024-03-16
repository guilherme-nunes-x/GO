package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello wold")

	//strings = textos

	var name1 string = "jonh"
	var name2 = "fernando"

	fmt.Println(name1, name2)

	name1 = "fernando"
	name2 = "jonh"

	fmt.Println(name1, name2)

	name3 := "estevã"

	fmt.Println(name3)

	//Inteiros = int

	var number int = 23
	var number2 int = 24
	var soma = 23 + 24

	fmt.Println(number, number2, soma)

	// numeros tipo float

	var num1 = 23.4
	var num2 = 25.9
	var soma2 float64 = num1 + num2

	fmt.Printf("Os valores foram %v e %v, e a soma %v\n", num1, num2, soma2)

	// teste com arrays

	var nomes = [4]string{"luis", "joao", "jack", "estevão"}

	fmt.Println(nomes, len(nomes))

	nomes[0] = "leo"
	fmt.Println(nomes, len(nomes))

	//slices não possuem limites pre-estabelecidos ou seja posso adicionar mais elementos

	var nuns = []int{12, 23, 34}
	fmt.Println(nuns, len(nuns))
	nuns = append(nuns, 98, 34)
	fmt.Println(nuns, len(nuns))

	// slice para retirar valores

	primer1 := nuns[0:1]
	primer2 := nuns[0:2]
	primer3 := nuns[0:3]
	primer4 := nuns[0:4]

	fmt.Println(primer1, primer2, primer3, primer4)

}
