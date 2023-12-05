package main

import (
	"demo/Clase02/salario"
	"fmt"
)

// const (
// 	MINIMUN = "minimum"
// 	AVERAGE = "average"
// 	MAXIMUN = "maximum"
// )

func main() {
	// resultado := suma(4, 8)
	// fmt.Println("El resultado es:", resultado)
	//-------------------------------------- EJERCICIO 1∫ --------------------

	// result := salario.SalaryTax()

	// fmt.Printf("El impuesto correspondiente es de: %.2f\n", result)

	// //-------------------------------------- EJERCICIO 2 --------------------

	// notas := []int{9, 8, 8, 9, 1}

	// promedio, err := salario.Average(notas...)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("El promedio es: %.2f\n", promedio)
	// }

	// //-------------------------------------- EJERCICIO 3 --------------------

	// salarios, err := salario.Salary("r", 2354)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("El salario calculado es:$ %.2f\n", salarios)
	// }

	//-------------------------------------- EJERCICIO 4 --------------------
	// type calculationFunc func(...int) (float64, error)

	// operationType := salario.AVERAGE
	// numbers := []int{2, 3, 3, 4, 1, 2, 4, 5}

	// calcFunc, err := salario.Operation(operationType, numbers...)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// result, err := calcFunc(numbers...)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("El resultado de la operación %s es: %.2f\n", operationType, result)
	// }

	//-------------------------------------- EJERCICIO 5 --------------------

	animalType, err := salario.Animal(salario.Gato)

	if err != nil {
		panic(err)
	}

	result := animalType(10)

	fmt.Printf("La cantidad de comida necesaria para %s si son %d es de %.2f kg\n", salario.Gato, 10, result)
}

// func suma(a, b int) int {

// 	return a + b
// }
