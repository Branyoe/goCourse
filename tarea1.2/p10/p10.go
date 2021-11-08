package p10

import "fmt"

func Trabajo() {
	var dias [2]string
	for i := 0; i < 2; i++ {
		fmt.Printf("Escribe el dia #%d\n", i+1)
		fmt.Scan(&dias[i])
	}
	fmt.Printf("Trabajamos de %s a %s", dias[0], dias[1])
}