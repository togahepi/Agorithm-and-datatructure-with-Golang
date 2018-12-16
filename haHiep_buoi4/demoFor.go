package main

import "fmt"

func main() {
	x := 0;

	for i:=1; i<5; i++ {
		x = x*2 + i
		x++
		fmt.Printf("%d ",x)
	}
}
