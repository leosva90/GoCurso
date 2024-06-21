package ejer_interfaces

import (
	"fmt"

	"github.com/leosva90/GoCurso/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Println("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
