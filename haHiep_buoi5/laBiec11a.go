package main

import (
	"../string"
	"fmt"
)

func main() {
	var a,b int
	fmt.Scan(&a,&b)

	string.DoPrint("*",a)
	fmt.Println()
	for i := 0; i < (b-2); i++ {
		fmt.Print("*")
		string.DoPrint(" ",a-2)
		fmt.Print("*")
		fmt.Println()
	}
	string.DoPrint("*",a)

}
