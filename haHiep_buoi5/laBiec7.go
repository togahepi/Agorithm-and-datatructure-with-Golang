package main

import (
	"../string"
	"fmt"
)

func main() {
	var chieuCao int
	fmt.Scan(&chieuCao)

	soDauSao := 1
	soDauCach := chieuCao - 1
	for i := 0; i < chieuCao; i++ {
		string.DoPrint(" ",soDauCach)
		string.DoPrint("*",soDauSao)
		fmt.Printf("\n")
		soDauSao = soDauSao + 2
		soDauCach--
	}
}