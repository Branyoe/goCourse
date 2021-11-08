package main

import (
	"fmt"
	"helloWord/tarea1.2/p1"
	"helloWord/tarea1.2/p10"
	"helloWord/tarea1.2/p11"
	"helloWord/tarea1.2/p2"
	"helloWord/tarea1.2/p3"
	"helloWord/tarea1.2/p4"
	"helloWord/tarea1.2/p5"
	"helloWord/tarea1.2/p6"
	"helloWord/tarea1.2/p7"
	"helloWord/tarea1.2/p8"
	"helloWord/tarea1.2/p9"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("%s\n", p1.Ciudades()[i]);
	}
	fmt.Printf("*********************\n")
	p2.Animales()
	fmt.Printf("*********************\n")
	p3.Suma(3, 5)
	fmt.Printf("\n*********************\n")
	fmt.Print(p4.Pow(3));
	fmt.Printf("\n*********************\n")
	fmt.Print(p5.Cadena());
	fmt.Printf("\n*********************\n")
	p6.Citys()
	fmt.Printf("\n*********************\n")
	p7.Verano()
	fmt.Printf("\n*********************\n")
	p8.Estaciones()
	fmt.Printf("\n*********************\n")
	p9.Descanso()
	fmt.Printf("\n*********************\n")
	p10.Trabajo()
	fmt.Printf("\n*********************\n")
	fmt.Printf("%s %s %s %s", p11.MyName()[0], p11.MyName()[1], p11.MyName()[2], p11.MyName()[3])
}