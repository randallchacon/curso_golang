package mypackage

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct { //para hacer esta clase publica basta con las mayúsculas
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrintMessahe imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
