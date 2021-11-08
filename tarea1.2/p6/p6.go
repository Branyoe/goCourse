package p6

import (
	"fmt"
)

func Citys() {
	var citys [4]string 
	for i := 0; i < 4; i++ {
		fmt.Printf("Escribe la ciudad #%d\n", i+1);
		fmt.Scan(&citys[i])
	}
	fmt.Printf("Las ciudades son %s, %s, %s y %s", citys[0], citys[1], citys[2], citys[3])
}