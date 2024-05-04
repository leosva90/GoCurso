package main

import (
	"fmt"

	"github.com/leosva90/GoCurso/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1400)
	fmt.Println(estado)
	fmt.Println(texto)
}
