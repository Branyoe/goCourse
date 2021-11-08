package program5

import (
	"fmt"
	"strconv"
)

//Forma 1
var name string  = "Brandon"
var ciudad string = "Villa de Alvarez"
var edad string = "17"

var Forma1 = "Hola, soy " + name + ", vivo en " + ciudad + " y tengo " + edad + " años"

//Forma 2

func Forma2(name string, ciudad string, edad int)  {
	i1 := edad
	srt1 := strconv.Itoa(i1)
	fmt.Printf("Hola, soy %s, vivo en %s y tengo %s años\n\n", name, ciudad, srt1)
}