package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scanf("%d",&a)

	fmt.Printf("%s\n",laSoChinhPhuong(a))
}

func laSoChinhPhuong(a int) string {
	if a == 0 {
		return "false"
	}

	b := math.Sqrt(float64(a))
	bNguyen := int(b)

	if a == bNguyen*bNguyen {
		return "true"
	}
	return "false"
}
