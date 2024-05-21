package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// The Standard Library (são um conjunto de funções que podem ser importadas na função "import"
	// site: https://pkg.go.dev/std onde podem ser observados

	//Import string

	var mensagem = "Oi pedro"

	// O pacote string me permite analisar as mensagens em uma string e a função chamada Contains estar analisando a presença dessa parte do texto na mensagem
	fmt.Println(strings.Contains(mensagem, "Oi"))

	//Agora um teste com a função ReplaceAll(Procura um termo na string e substitui) e ainda retona a mensagem modificada
	fmt.Println(strings.ReplaceAll(mensagem, "Oi", "Hello"))

	//A mensagem continua modificada após usar essa função
	// fmt.Println("A parti das modificações acima a mensagem foi =", mensagem)

	//Agora vamos usar uma função que deixa toda a mensagem maiuscula

	fmt.Println(strings.ToUpper(mensagem))

	//Função index procura um valor na variavel desejado e revela sua posição

	fmt.Println(strings.Index(mensagem, "p"))

	//Split é uma funcionalidade que retira um valor desejado da string

	fmt.Println(strings.Split(mensagem, " "))

	//Import Sort e função int = organiza os arrays por numero

	notas := []int{12, 24, 13, 45, 5, 6}

	//pre organização
	fmt.Println(notas)
	//pos organização
	sort.Ints(notas)
	fmt.Println(notas)

	//Função SeachInts procura um valor inteiro em um arrays e procura sua posição

	metodo := sort.SearchInts(notas, 24)

	fmt.Println(metodo)
	//Lembrando que as posições foram atualizadas anteriomente pelo metodo int

	//agora testando com o metodo String ele replica o mesmo que o sort.Int só que com strings
	
	names := []string{"Eduardo", "Pedro", "leon", "gabriel","felipe"}

	sort.Strings(names)
	fmt.Println(names)

	//Ultimo exemplo vamos testar o search string

fmt.Println(sort.SearchStrings(names, "Pedro"))

}
