package p7

import "fmt"

func Verano() {
	var meses [4]string
	for i := 0; i < 4; i++ {
		fmt.Printf("Escribe el mes #%d\n", i+1)
		fmt.Scan(&meses[i])
	}
	fmt.Printf("el verano se da en los meses de %s, %s, %s y parte de %s", meses[0], meses[1], meses[2], meses[3])
}