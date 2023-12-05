package archivos

import (
	"io"
	"os"
)

func ReadFile(path string) (texto string, err error) {

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)

	texto = string(data)
	// buf := make([]byte, 2)
	// for {
	// 	var n int
	// 	n, err = f.Read(buf)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			err = nil
	// 			return
	// 		}
	// 		return
	// 	}
	// 	fmt.Println(string(buf[n]))
	// 	data += string(buf[n])
	// }
	return
}
