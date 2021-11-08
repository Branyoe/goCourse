package p9

import "fmt"

func Descanso() {
	var dias [2]string
	for i := 0; i < 2; i++ {
		fmt.Printf("Escribe el dia #%d\n", i+1)
		fmt.Scan(&dias[i])
	}
	fmt.Printf("Descansamos %s y %s", dias[0], dias[1])
}