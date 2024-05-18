package main

import "fmt"

func main() {
	var name = "jonathan"
	var number = 12

	//Print  = não pula linha existe necessidade de \n

	fmt.Print("hello")
	fmt.Print(" world\n")
	//Usando o \n conseguimos separar essas duas mensagens
	fmt.Print("Hello")
	fmt.Print(" world")

	//Println pula de linha automatica

	fmt.Println("Command line")
	fmt.Println("Pula de linha automatica")
	// Teste de exposição de variaveis em uma mensagem
	fmt.Println("Meu nome é", name, "E minha idade é", number)

	//Printf = organizar strings e utilizar %_ formatação expecifica e não ocorre a quebra de linha

	name = "ricardo"
	number = 23
	number1 := "23"

	//espor variavel normalmente
	fmt.Printf("Meu nome é %v, e minha idade é %v\n", name, number)
	//expor variavel string entre ""
	fmt.Printf("Meu nome é %q, e minha idade é %q\n", name, number1)
	//Expor tipo de variavel
	fmt.Printf("Meu nome é %T\n", name)
	//expor valor sem variavel
	fmt.Printf("Eu tenho %f reais\n", 34.5)
	//limitar casas decimais de float
	fmt.Printf("Eu tenho %0.2f reais", 34.5)

	//Sprintf = salvar formatação de string(transformação em variavel e possivel reutilização)

	var mensagem = fmt.Sprintf("Hello my name is %v\n", name)
	fmt.Println("Eu sei ingles agora : ", mensagem)
}
