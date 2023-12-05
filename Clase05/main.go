package main

import (
	"demo/Clase05/archivos"
	"fmt"
)

func main() {

	path := "./archivo.txt"

	data, err := archivos.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	//myError.WriteTxt("archivo.txt", "Hola soy un texto de prueba")

}
