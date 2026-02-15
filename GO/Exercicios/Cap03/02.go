package main

import "fmt"

/*
Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a
estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
 1. Identificador "x" deverá ter tipo int
 2. Identificador "y" deverá ter tipo string
 3. Identificador "z" deverá ter tipo bool

Na função main:
 1. Demonstre os valores de cada identificador
 2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
*/

var x00 int
var y00 string
var z00 bool

func exC302() {
	fmt.Println("O valor contido em x é:", x00)
	fmt.Println("O valor contido em y é:", y00)
	fmt.Println("O valor contido em x é:", z00)
}
