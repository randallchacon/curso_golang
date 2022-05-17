package main

import "fmt"

func say(text string, c chan<- string) { //indicar si los parametros son de entrada o de salida
	c <- text //simbolo de ingreso de datos al canal
}

func main() { //los channels siempre son la mejor opción
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c) // simbolo de salida de datos del canal
}
