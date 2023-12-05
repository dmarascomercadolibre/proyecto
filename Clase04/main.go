package main

import (
	"demo/Clase04/myError"
	"fmt"
)

func main() {

	// var salary int = 1500

	// err := myError.IsPayTax(salary)

	// if err != nil {
	// 	_, ok := err.(*myError.ErrTax)
	// 	if !ok {
	// 		fmt.Println("Error: ", err)
	// 		return
	// 	}
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Must pay tax")
	// }

	// //-----------------------

	// err2 := myError.IsLess(salary)

	// if errors.Is(err2, myError.ErrorLess) {
	// 	fmt.Println("Error:", err2)
	// }

	// //-----------------

	// salaryMonth, err3 := myError.CalculateSalaryLessTax(79, 2000)

	// if err3 != nil {
	// 	if errors.Is(err3, myError.MinimumHoursError) {
	// 		fmt.Println("Error: ", err3)
	// 		return
	// 	}

	// }

	// fmt.Printf("The corresponding salary for the hours worked is: %.2f\n", salaryMonth)

	// defer myError.Adios()
	// fmt.Println("sigo vivo")
	// panic("muelte")

	// err := myError.WriteTxt("test.txt", "Hola mundo")
	// if err != nil {
	// 	panic(err)
	// }

	// myError.IsPair(3)
	// fmt.Println("Fin de la ejecución")

	myError.WriteTxt("customer.txt", "Pedro Suarez")
	fmt.Println("Fin de la ejecución")
}
