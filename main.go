package main

import "fmt"

func main()  {
	var nombre string
	var edad int
	//fmt.Println("calando Git")
	fmt.Printf("cual es tu nombre: ")
	fmt.Scanf("%s\n",&nombre)
	fmt.Printf("cual es tu edad: ")
	fmt.Scanf("%d\n",&edad)

	fmt.Println("Que onda",nombre,"\ntu edad es",edad)
}