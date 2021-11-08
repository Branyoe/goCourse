package program4

import (
	"fmt"
	"strconv"
)

func Forma1() {
	var name string = "Brandon"
	var edad int = 17

	fmt.Printf("%s %d", name, edad)
}

func Forma2 (name string, edad int) string{
	i1 := edad
	srt1 := strconv.Itoa(i1)
	return name + " " + srt1
}