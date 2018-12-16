package main

import (
	"../string"
	"fmt"
)

func main() {
	var chieuCaoX int
	fmt.Scan(&chieuCaoX)

	for i := 0; i < chieuCaoX; i++ {
		for j := 0; j < chieuCaoX; j++ {
			if i == j || i + j == chieuCaoX -1 {
				string.DoPrint("*",1)
			} else {
				string.DoPrint(" ",1)
			}
		}
		fmt.Println()
	}
}
