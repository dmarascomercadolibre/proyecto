package alumno

import (
	"fmt"
	"time"
)

type Alumno struct {
	Name     string
	Apellido string
	Dni      int
	Fecha    time.Time
}

func (a Alumno) Print() {
	fmt.Printf("Alumno:\nNombre: %s\nApellido: %s\nDni: %d\nFecha: %s\n", a.Name, a.Apellido, a.Dni, a.Fecha)
}
