package main

import "fmt"

func main() {

	var a int

	fmt.Scanf("%d",&a)

	fmt.Printf("%s",laChiaHetCho12(a))
}

func laChiaHetCho12(num int) string {
	if num%12 == 0 {
		return "Yes"
	}

	return "No"
}