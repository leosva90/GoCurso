package arreglos_slice

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 4, 90}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //Slice creado desde un vetor, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] //Slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

/*func Capacidad(){
nums := make([]int, 0, 0)
	for i:=0, i<100, i++ {
		nums = append(nums, i)
	}
func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Println("Largo %d, Capacidad  %d", len(elementos), cap(nums))
}

}*/
