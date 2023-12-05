package myError

import (
	"fmt"
	"io"
	"os"
)

func WriteTxt(path string, text string) (err error) {

	//Abrir el archivo
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// _, err = os.Open(path)
	// if err != nil {
	// 	panic("The indicated file vas not found or is damaged")
	// }

	file, err := os.Create(path)
	if err != nil {
		return
	}

	_, err = io.WriteString(file, text)
	if err != nil {
		return
	}

	return
}
func IsPair(num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if num%2 != 0 {
		panic("Número impar\n")
	}
	fmt.Print("Número par")
}
