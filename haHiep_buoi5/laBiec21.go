package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var x,so,min,dem int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		fmt.Scan(&so)
		min = so
		dem = 1
		for y := 1; y < x; y++ {
			fmt.Scan(&so)
			if so < min {
				min = so
				dem = 1
			} else if so == min {
				dem++
			}
		}
		fmt.Printf("%d %d",min,dem)
	}
}
