package main

import (
	"agoGolang/lib"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var soLuong int
	for i := 0; i < n; i++ {

		fmt.Scan(&soLuong)
		slice := lib.InputSlice(soLuong)

		fmt.Println(danDau(slice))
	}
}

func danDau(slice []int)  string{

	for i:=0; i<len(slice)-1; i++ {
		if slice[i] == 0 {
			return "0"
		} else if slice[i]*slice[i+1] > 0 {
			return "0"
		}
	}
	return "1"
}
