package product

import (
	"errors"
	"fmt"
)

var Products = []Product{
	{
		Id:          1,
		Name:        "Coca Cola",
		Price:       200,
		Description: "Gaseosa sabor cola",
		Category:    "gaseosas",
	},
	{
		Id:          2,
		Name:        "Blanca Flor",
		Price:       300.5,
		Description: "Harina 0000",
		Category:    "almacen",
	},
	{
		Id:          3,
		Name:        "Hellmans",
		Price:       700,
		Description: "Mayonesa clásica",
		Category:    "almacen",
	},
}

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() (result string) {

	sliceOriginal := Products
	nuevoSlice := append(Products, p)

	if len(nuevoSlice) > len(sliceOriginal) {
		result = "se agrego correctamente el producto"
	} else {
		result = "no se pudo agregar el nuevo producto"
	}
	return
}

func (p Product) GetAll() {
	fmt.Println("Lista de Productos:")
	for _, product := range Products {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Descripción: %s, Categoría: %s\n",
			product.Id, product.Name, product.Price, product.Description, product.Category)
	}
}

func (p Product) GetById(id int) (prod Product, err error) {

	for _, value := range Products {
		if value.Id == id {
			prod = value
		}
	}

	if prod.Id == 0 {
		err = errors.New("no se ecuentra el producto con id: %d ")
	}

	return
}
