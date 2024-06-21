package main

import (
	"fmt"

	"github.com/leosva90/GoCurso/goroutines"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1400)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows,es", os)

	} else {
		fmt.Println("Esto es windows")

	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Println("%s \n", os)
	}

	numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()

	iteraciones.Iterar()*/

	//fmt.Println(ejercicios.TabledeMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()
	//files.LeoArchivo()

	//funciones.Exponencia(2)

	//arreglos_slice.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	//Pedro := new(modelos.Hombre)
	//e.HumanosRespirando(Pedro)

	//defer_panic.EjemploPanic()

	canal1 := make(chan bool)

	go goroutines.MiNombreLento("Crash Overide", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aca")

}
