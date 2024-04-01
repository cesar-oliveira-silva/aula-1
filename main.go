package main

import "fmt"

func main(){


	// exercicio 1
	nome := "Cesar Silva"
	idade := 36

	fmt.Printf("Nome: %v \nIdade: %v \n", nome, idade)

	//exercicio 2
	temperatura :=23.5
	umidade := 65.2
	pressao := 44.1

	fmt.Printf("Temperatura: %v \nUmidade: %v \nPressão: %v \n", temperatura, umidade, pressao)
	fmt.Printf("Tipo de dado dos campos\n Temperatura: %T \nUmidade: %T \nPressao %T",temperatura, umidade, pressao)


//	//exercicio 3
//	//var 1nome string !nome da variavel nao pode comecar com numeral. correto: 
//	var nome1 string
//  var sobrenome string //correto

//  //var int idade !nome e tipo de dado estao invertidos. correto:
//  var idade int
//  //1sobrenome := 6  !nome da variavel nao pode comecar com numeral 
//  //! ja existe uma variavel com nome muito parecido. nao seria indicado declarar assim, apesar que corrigindo o nome, a declaracao esta correta
//  sobrenome1 := 6 
//  //var licenca_para_dirigir = true ! declaracao esta errada. utilizar := sem o var. ou adicionar o tipo e depois realizar a atribuicao
//  var licenca_para_dirigir bool
//  licenca_para_dirigir = true
//  //var estatura da pessoa int  ! nome da variavel no pode conter espacos. correto:
//  var estatura_da_pessoa int
//  quantidadeDeFilhos := 2

// exercicio 4

// var sobrenome string = "Silva"
//   var idade int = "25"
//   boolean := "false";
//   var salario string = 4585.90
//   var nome string = "Fellipe"

//resposta:
// var sobrenome string
// var idade int
// var salario string
// var nome string


// sobrenome = "Silva"
// idade = "25"
// boolean := "false"
// salario = 4585.90
// nome = "Fellipe"

	
}