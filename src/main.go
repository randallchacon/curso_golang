package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c)) //capacidad máxima de dos

	//Range y close
	close(c) //cerrar los canales una vez que ya no se usan
	//c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case em1 := <-email1:
			fmt.Println("Email recibido de email 1", em1)
		case em2 := <-email2:
			fmt.Println("Email recibido de email 2", em2)
		}
	}
}
