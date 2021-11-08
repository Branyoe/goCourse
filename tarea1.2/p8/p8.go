package p8

import "fmt"

func Estaciones() {
	var estaciones [4]string
	for i := 0; i < 4; i++ {
		fmt.Printf("Escribe la estación #%d\n", i+1)
		fmt.Scan(&estaciones[i])
	}
	fmt.Printf("Las estaciones son %s, %s, %s e %s", estaciones[0], estaciones[1], estaciones[2], estaciones[3])
}