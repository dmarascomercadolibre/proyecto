package main

import "fmt"

func main() {

	// WordRange()
	// Loan()
	// PrintMonth()
	AgeEmploye()
}

func WordRange() {
	// Inicializa la palabra
	palabra := "Damian"

	// Imprime la cantidad de letras en la palabra
	cantidadLetras := len(palabra)
	fmt.Printf("La palabra '%s' tiene %d letras.\n", palabra, cantidadLetras)

	// Imprime cada letra por separado
	fmt.Println("Letras:")
	for _, letra := range palabra {
		fmt.Printf("%c\n", letra)
	}
}

func Loan() {

	// Datos Cliente
	edad := 25
	empleado := false
	antiguedadTrabajo := 2
	sueldo := 90000

	if edad > 22 && empleado && antiguedadTrabajo > 1 {
		if sueldo > 100000 {
			fmt.Println("Ud es elegible para la acreditación del préstamo pero se le cobrará interes.")
		} else {
			fmt.Println("Ud es elegible para la acreditación del préstamo pero no se le cobrará interes.")
		}
	} else {
		fmt.Println("Ud no es elegible para la acreditación del préstamo ya que no cumple los requisitos.")
	}
}

func PrintMonth() {

	mes := 13

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("El número ingresado no corresponde a un mes del año")

	}

}

func AgeEmploye() {

	var employes = map[string]int{"Benjamin": 20, "Damián": 41, "Pedro": 25, "Joaquín": 34, "Raúl": 27}

	var mayores int = 0

	for nombre, edad := range employes {
		fmt.Printf("Nombre: %s, Edad: %d\n", nombre, edad)

		if edad > 21 {
			mayores++
		}
	}

	fmt.Printf("La cantidad de empleados mayores a 21 años son %d\n", mayores)

	employes["Federico"] = 30

	empleadoABorrar := "Pedro"

	delete(employes, empleadoABorrar)

	for nombre, edad := range employes {
		fmt.Printf("Nombre: %s, Edad: %d\n", nombre, edad)
	}
}
