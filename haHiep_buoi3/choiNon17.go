package main

import "fmt"

func main() {

	var a,b int

	fmt.Scanf("%d %d",&a,&b)

	fmt.Printf("%d\n",chiaHetCho10(a*b))
}

func chiaHetCho10(number int) int {

	if number%10 == 0 {
		return 1
	}

	return 0
}