package main

import "fmt"

// Usando o operador curto de declaração dê os seguintes valores às variáveis:
// x = 42, y = "James Bond", z = true e depois imprima com uma única declaração
// de print e depois múltiplas declarações
func ex01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("O ator do novo filme do", y, "tem", x, "anos de idade? É sério?", z)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
