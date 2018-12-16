package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%d",&x)

	for i:=0; i<x; i++ {
		for y:=0; y<10; y++ {
			fmt.Printf("%d ",i)
		}
		fmt.Printf("\n")
	}
}


