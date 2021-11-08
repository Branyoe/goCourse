package p5

import (
	"fmt"
)

func Cadena() string {
	var ciudad, dia, mes, año string
	fmt.Printf("ingresa tu ciudad:\n");
	fmt.Scan(&ciudad);
	fmt.Printf("ingresa un dia:\n");
	fmt.Scan(&dia);
	fmt.Printf("ingresa un mes:\n");
	fmt.Scan(&mes);
	fmt.Printf("ingresa un año:\n");
	fmt.Scan(&año);
	return ciudad + " a " + dia + " de " + mes + " de " + año 
}