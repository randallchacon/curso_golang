package main //nombre de la carpeta donde esta guardado

import (
	"curso_golang/src/mypackage"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

type car struct {
	brand string
	year  int
}

//El reto es colocar todo el código normal en funciones

func main() { //go me obliga a escribir bien el codigo con buenas practicas

	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value1, value2 := doubleReturn(2) //me retorna dos valores
	fmt.Println("value1 y value2", value1, value2)

	value3, _ := doubleReturn(2) //me retorna uno de los dos
	fmt.Println("value1", value3)

	value := returnValue(2)
	fmt.Println(value)

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12          //los dos puntos indican q la variable no se ha utilizado antes, no se le agrega el tipo de dato
	var altura int = 14 //se declara variable y se le asigna el valor
	var area int

	//hasta no usar las variables el codigo de go no corre
	//evita gasta memoria que no se esta usando
	fmt.Println(base, altura, area)

	//zero values, valores por defecto
	var a int     //cero
	var b float64 //cero
	var c string  //espacio vacio
	var d bool    //false

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//División
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x //para saber si es par o impar
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Reto
	//Area rectangulo, trapecio y un circulo

	//Clase10
	helloMessage := "Hello"
	worldMessage := "World"

	//println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//printf agrega una funcion extra de valor de entrada
	nombre := "randall"
	cursos := 5
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos) //%s string, %d numero entero \n salto de linea
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) //si no sé el tipo de dato le coloco %v

	//sprintf genera un string pero no lo imprime en consola, solamente lo guarda como un string
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Conocer el Tipo dato de una variable con %T
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	/*
		counterForever := 0 //detener con CTRL+C
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/

	//Reto hacer un for en retroceso

	//operadores lógicos y de comparacion son iguales a C#

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("verdadero")
	}

	//convertir texto a num
	value, err := strconv.Atoi("53") //retorna el valor y si existe un error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)

	//Reto determinar si un numero es par o impar
	//Reto determinar si un usuario y un password son correctos

	//Switch para multiples condiciones anidadas
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Otra forma
	switch modulo1 := 4 % 2; modulo1 {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condicion
	value4 := 200
	switch {
	case value4 > 100:
		fmt.Println("Es mayor a 100")
	case value4 < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

	//Defer como buena practica se utiliza una sola vez, y sera lo último q se ejecute
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

	//Listas: arrays y slices
	//Arrays son inmutables, o bien no se les puede agregar más elementos de código, más eficientes en código
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array)) //cap me indica capacidad max de un array

	//slices no se le indica la cantidad de valores que va a tener
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4]) //el primer # es inclusivo el segundo # es exclusivo por eso no imprime el 4
	fmt.Println(slice[4:])

	//agregar elementos al slice porq son mutables
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//recorrido de arrays con range
	slice2 := []string{"Hola", "que", "hace"}

	for i, valor := range slice2 {
		fmt.Println(i, valor) //indice y valor
	}

	for _, valor := range slice2 {
		fmt.Println(valor) //omitir indice e imprimir valor
	}

	for i := range slice2 {
		fmt.Println(i) //imprimir solo el índice
	}

	//Palindromo palabra q se lee igual en ambos sentidos
	//ama

	isPalindromo("ama") //Reto: identificar el palindromo a pesar de ser mayuscula o minuscula

	var palabra string
	fmt.Scan(&palabra) //& se debe agregar para que el programa espere
	minuscula := strings.ToLower(palabra)
	isPalindromo(minuscula)

	//Maps  Llave-valor  en C# diccionarios, a diferencia de los arrays y slices, los Maps me permiten manejar concurrencia
	m := make(map[string]int)
	m["Jose"] = 14
	m["Randall"] = 20

	fmt.Println(m)

	//Recorrer map
	for indice, valor := range m { //se puede ejecutar en orden aleatorio, si se necesita estrictamente en orden es mejor usar arrays
		fmt.Println(indice, valor)
	}

	//encontrar un valor
	value5, ok := m["Josep"] //si no existe entonces imprime el zero value, el ok me indica si la llave existe o no
	fmt.Println(value5, ok)

	value6, ok := m["Jose"]
	fmt.Println(value6, ok)

	//Struct o bien clases
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Mercedes"
	fmt.Println(otherCar)

	var myCar2 mypackage.CarPublic
	myCar2.Brand = "Ferrari"
	fmt.Println(myCar2)

	mypackage.PrintMessage("Hola Randall")
}
