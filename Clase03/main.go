package main

import (
	"demo/Clase03/alumno"
	"demo/Clase03/producto"

	// "demo/Clase03/person"
	// "demo/Clase03/product"
	"fmt"
	"time"
)

func main() {

	// //--------------------------- EJERCICIO 1 ------------------------
	// p := product.Product{
	// 	Id:          4,
	// 	Name:        "Sprite",
	// 	Price:       200,
	// 	Description: "Gaseosa sabor lima limon",
	// 	Category:    "gaseosas",
	// }

	// p.GetAll()
	// fmt.Println(product.Product.Save(p))

	// producto, err := product.Product.GetById(p, 1)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("el producto con id: %d es\nnombre: %s,\nprecio: %.2f\n", producto.Id, producto.Name, producto.Price)

	// //--------------------------- EJERCICIO 2 ------------------------

	// persona := person.Person{
	// 	ID:          1,
	// 	Name:        "Damian Marasco",
	// 	DateOfBirth: "20/02/1982",
	// }

	// employee := person.Employee{
	// 	Person: &person.Person{
	// 		ID:          2,
	// 		Name:        "John Doe",
	// 		DateOfBirth: "1987-01-12",
	// 	},
	// 	Position: "Software Engineer",
	// }

	// persona.PrintEmployee(persona)
	// employee.PrintEmployee(employee)

	//----------------------------
	alumno := alumno.Alumno{
		Name:     "Damian",
		Apellido: "Marasco",
		Dni:      29392118,
		Fecha:    time.Now(),
	}

	alumno.Print()

	//------------------------------

	productType := "Medium"
	productCost := 100.0

	// Crear un producto usando la función factory
	product := producto.CreateProduct(productType, productCost)

	// Verificar si el producto fue creado exitosamente
	if product != nil {
		// Llamar al método Price y mostrar el resultado
		totalPrice := product.Price()
		fmt.Printf("El precio total del producto %s es: $%.2f\n", productType, totalPrice)
	} else {
		fmt.Println("Tipo de producto no válido.")
	}
}
