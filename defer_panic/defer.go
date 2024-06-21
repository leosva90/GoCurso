package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("este es un primer mensaje")
	defer fmt.Println("este es un mensaje final")

	fmt.Println("este es e segundo mensaje")
}

func EjemploPanic() {
	defer func ()  {
		reco := recover()
		if reco != nil {
          log.Fatalf("ocurrio un error que generó panic \n %v", reco)
		
	}()
	a :=
		1
	if a == 1 {
		panic("el valor de a es 1")
	}
}
