package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup //son complicados de mantener con el tiempo

	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg) //concurrencia
	wg.Wait()
	//time.Sleep(time.Second * 1) //esto no es eficiente, porque es contra producente
	//las goroutines es muy comun que se usen como funciones anonimas

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
