package person

import (
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	*Person
	Position string
}

func (e Employee) PrintEmployee(empl Employee) {
	fmt.Printf("Id: %d \nNombre: %s\nNacimiento: %v \nPosicion: %s \n", empl.Person.ID, empl.Name, empl.Person.DateOfBirth, empl.Position)
}
func (p Person) PrintEmployee(per Person) {
	fmt.Printf("Id: %d \nNombre: %s\nNacimiento: %v \n", per.ID, per.Name, per.DateOfBirth)
}
