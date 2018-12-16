package main

import "fmt"

func main() {
	var soCot,soHang int
	fmt.Scanf("%d %d",&soCot,&soHang)

	for i:=0; i<soHang; i++ {
		for y:=0; y<soCot; y++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
