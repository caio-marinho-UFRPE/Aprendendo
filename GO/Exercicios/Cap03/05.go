package main

import "fmt"

/*Em package-level scope, utilizando a palavra-chave var, crie uma variável com
	o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo
	que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x"
			em seu tipo subjacente e, utilizando o operador "=", atribua o valor
			de "x" a "y"
            2. Demonstre o valor de "y"*/

type tipoB int

var x04 tipoB

var y03 int

func exC305() {
	fmt.Printf("X é uma variável de tipo: %T, de valor %v\n", x04, x04)

	x04 = 42

	fmt.Printf("O novo valor de X é %v\n", x04)

	y03 = int(x04)

	fmt.Printf("Y é uma variável de tipo: %T, de valor %v\n", y03, y03)
}
