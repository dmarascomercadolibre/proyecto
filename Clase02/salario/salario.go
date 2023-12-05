package salario

import (
	"errors"
	"strings"
)

const (
	MINIMUN   = "minimum"
	AVERAGE   = "average"
	MAXIMUN   = "maximum"
	Perro     = "perro"
	Gato      = "gato"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func SalaryTax() (impuesto float64) {

	var sueldo float64 = 200000.0

	switch {
	case sueldo > 50000 && sueldo > 150000:
		impuesto = calculateTax(27, sueldo)
		return
	case sueldo > 50000 && sueldo < 150000:
		impuesto = calculateTax(10, sueldo)
		return
	}
	return
}

func calculateTax(porcentaje, sueldo float64) (impuesto float64) {

	impuesto = (sueldo * porcentaje) / 100

	return
}

func Average(notas ...int) (promedio float64, err error) {

	sumatoria := 0
	contador := len(notas)

	for _, value := range notas {
		if value < 0 {
			err = errors.New("No se pueden incluir números negativos")
			return
		}
		sumatoria += value
	}

	promedio = float64(sumatoria) / float64(contador)

	return
}

func Salary(categoria string, minutosTrabajados int) (salario float64, err error) {

	categoria = strings.ToUpper(categoria)

	switch categoria {
	case "A":
		salario = CalculateSalary(3000, minutosTrabajados)
	case "B":
		salario = CalculateSalary(1500, minutosTrabajados)
	case "C":
		salario = CalculateSalary(1000, minutosTrabajados)
	default:
		err = errors.New("La categoría ingresada no existe")
	}

	return
}

func CalculateSalary(salarioPorHora int, minutosTrabajados int) (salario float64) {

	salario = float64(salarioPorHora) * (float64(minutosTrabajados) / 60)

	return
}

type calculationFunc func(...int) (float64, error)

func Operation(operationType string, notas ...int) (calculationFunc, error) {
	switch operationType {
	case MINIMUN:
		return func(...int) (float64, error) {
			minimo, err := minNoteFunc(notas...)
			return float64(minimo), err
		}, nil

	case AVERAGE:
		return func(...int) (float64, error) {
			promedio, err := averageNoteFunc(notas...)
			return promedio, err
		}, nil

	case MAXIMUN:
		return func(...int) (float64, error) {
			maximo, err := maxNoteFunc(notas...)
			return float64(maximo), err
		}, nil

	default:
		return nil, errors.New("Tipo de cálculo no válido.")
	}
}

func minNoteFunc(notas ...int) (minimo float64, err error) {

	if len(notas) == 0 {
		err = errors.New("No se puede realizar mínimo debido a que no se ingresaron notas")
		return
	}
	minimo = float64(notas[0])

	for _, value := range notas {
		if value < int(minimo) {
			minimo = float64(value)
		}
	}

	return

}

func maxNoteFunc(notas ...int) (maximo float64, err error) {

	if len(notas) == 0 {
		err = errors.New("No se puede realizar máximo debido a que no se ingresaron notas")
		return
	}
	maximo = float64(notas[0])

	for _, value := range notas {
		if value > int(maximo) {
			maximo = float64(value)
		}
	}

	return

}

func averageNoteFunc(notas ...int) (promedio float64, err error) {

	if len(notas) == 0 {
		err = errors.New("no se puede realizar promedios debido a que no se ingresaron notas")
		return
	}
	sumatoria := 0

	for _, value := range notas {
		sumatoria += value
	}

	promedio = float64(sumatoria) / float64(len(notas))

	return
}
func Animal(animal string) (funcion func(animalCantidad int) float64, err error) {

	switch animal {
	case Perro:
		funcion = DogFood
	case Gato:
		funcion = CatFood
	case Tarantula:
		funcion = SpiderFood
	case Hamster:
		funcion = HamsterFood
	default:
		err = errors.New("no existe el animal al cual quiere acceder")

	}

	return
}

func DogFood(cantidadAnimal int) (cantidadAlimento float64) {

	cantidadAlimento = float64(cantidadAnimal) * 10

	return
}
func SpiderFood(cantidadAnimal int) (cantidadAlimento float64) {

	cantidadAlimento = float64(cantidadAnimal) * 150

	return
}
func HamsterFood(cantidadAnimal int) (cantidadAlimento float64) {

	cantidadAlimento = float64(cantidadAnimal) * 250

	return
}
func CatFood(cantidadAnimal int) (cantidadAlimento float64) {

	cantidadAlimento = float64(cantidadAnimal) * 5

	return
}
