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

	var numberUm int = 2
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

	//calc

	div := numberdois / numberUm
	fmt.Println(div)
	ves := numberdois * numberUm
	fmt.Println(ves)
	soma := numberdois + numberUm
	fmt.Println(soma)
	sub := numberdois - numberUm
	fmt.Println(sub)

	//Arrays = armazenadores de varios dados

	var ages [5]int = [5]int{1, 2, 3, 4, 5}
	var age = [4]int{1, 2, 3, 4}
	names := [4]string{"guilherme", "eduardo", "gabriel", "luis"}

	//len é um comando que permite ver o tamanho de um arrays

	fmt.Println(age, len(age))
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	//substitui valor no array por posição
	ages[2] = 54
	fmt.Println(ages, len(ages))

	// Slices (usa fatias de um arrays)
	// existe a possibilidade não colocar um valor fixo no array
	var valores = []int{10, 20, 30, 40, 50}
	//substitui valor no array por posição
	valores[1] = 32
	//Append é capaz de adicionar novo valor no arrays
	valores = append(valores, 98)

	fmt.Println(valores, len(valores))

	// Ragens são intervalos de slice

	valor := valores[0:2]
	valor2 := names[2:]
	valor3 := names[:3]
	fmt.Println(valor, valor2, valor3)

	//uso de append dentro de um slice

	valor2 = append(valor2, "willian")

	fmt.Println(valor2)
}
