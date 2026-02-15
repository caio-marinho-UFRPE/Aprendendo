package main

import "fmt"

/*Crie um tipo. O tipo subjacente deve ser int.
Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"*/

type tipoA int

var x03 tipoA

func exC304() {
	fmt.Printf("X é uma variável de tipo: %T, de valor %v\n", x03, x03)

	x03 = 42

	fmt.Printf("O novo valor de X é %v", x03)
}
