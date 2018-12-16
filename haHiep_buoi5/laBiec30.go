package main

import (
	"../string"
	"fmt"
)

func main() {
	var chieuCaoX int
	fmt.Scan(&chieuCaoX)

	soDauCach1 := 0
	soDauCach2 := chieuCaoX - 2
	for i := 0; i < chieuCaoX/2; i++ {
		string.DoPrint(" ",soDauCach1)
		string.DoPrint("*",1)
		string.DoPrint(" ",soDauCach2)
		string.DoPrint("*",1)
		fmt.Println()
		soDauCach1++
		soDauCach2 = soDauCach2 - 2
	}
	string.DoPrint(" ",soDauCach1)
	string.DoPrint("*",1)
	fmt.Println()

	soDauCach1--
	soDauCach2 = 1
	for i := 0; i < chieuCaoX/2; i++ {
		string.DoPrint(" ",soDauCach1)
		string.DoPrint("*",1)
		string.DoPrint(" ",soDauCach2)
		string.DoPrint("*",1)
		fmt.Println()
		soDauCach1--
		soDauCach2 = soDauCach2 + 2
	}
}