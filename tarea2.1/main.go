package main

import (
	"fmt"
	"helloWord/tarea2.1/program1"
	"helloWord/tarea2.1/program2"
	"helloWord/tarea2.1/program3"
	"helloWord/tarea2.1/program4"
	"helloWord/tarea2.1/program5"
)

func main() {

	//Programa1
	fmt.Printf("\n*******Programa 1********\n\n")
		//Forma1
	program1.Forma1()
		//forma2
	program1.Forma2("Brandon Yoel\n")

	//Programa2
	fmt.Printf("\n*******Programa 2********\n\n")
		//Forma1
	program2.Forma1();
		//forma2
	program2.Forma2("Brandon Yoel", "Hernandez", "Ascencio")

	//Program3
	fmt.Printf("\n\n*******Programa 3********\n\n")
	fmt.Printf("%s\n", program3.Forma1)
	fmt.Printf("%s\n", program3.Forma2)

	//Program4
	fmt.Printf("\n*******Programa 4********\n\n")
	program4.Forma1()
	fmt.Printf("\n%s", program4.Forma2("Brandon", 17))

	//Program5
	fmt.Printf("\n*******Programa 5********\n\n")
	fmt.Printf("%s\n", program5.Forma1)
	program5.Forma2("Brandon", "Villa de Alvarez", 17)

}