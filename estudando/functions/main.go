package main

import "fmt"

func main() {
	fmt.Println("Vamos trabalhar com loops")

	//ele vai repetir enquanto o indice for menor que 10 ou valor qualquer desejado
	//esse modelo é desorganizado
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("O indice é de valor igual", i)
	// 	i++
	// }

	// Modelo padrão e mais organizado
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("O indice é de valor", i)
	// }

	names := []string{"luis", "leo", "edgar", "estevã", "luiser"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("O indice é de valor", i, "E o nome na posição é igual a:", names[i])
	// }

	//nessa funcionalidade values pega todos os nomes do arrays pela funcionalidade range e o index recolhe o numeros e se torna o indice
	//  for index, values := range names {
	//  	fmt.Printf("O index é igual a %v e o valor igual %v\n", index, values)
	//  }

	//Podemos usar o '_' ao inves do indez se tornando mais facil mas impedindo de retirar o valor obtido pelo index
	for _, values := range names {
		fmt.Printf("O valor vai ser igual a %v\n", values)
		values = "leonardo"
	}
	//Não vai ocorrer modificações apesar da mudança realizada no for anterior
	fmt.Println(names)

	//booloneanos e condições

	idade := 15
//determinar se essas afirmações abaixo são verdadeiras
	fmt.Println(idade > 2)
	fmt.Println(idade < 21)
	fmt.Println(idade == 12)
	fmt.Println(idade != 16)

	if idade > 10 {
		fmt.Println("idade maior que 10")
	} else if idade < 10 {
		fmt.Println("idade menor que 10")
	}

	for index, values := range names {
		if index == 1 {
			fmt.Printf("O nome na posição 1 é %v\n", values)
			//mesmo se essa função for correspondente ela ira continuar
			continue 
		}
		if index > 2{
			fmt.Println("A quebra de loop é igual a index:", index)
			//quando essa condição for completa o loop encerra
			break
		}
		fmt.Printf("O index é igual a %v e o valor igual %v\n", index, values)
	}
}
