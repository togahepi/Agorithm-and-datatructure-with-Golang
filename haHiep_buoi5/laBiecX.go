package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var so int
	for i := 0; i < n; i++ {
		fmt.Scan(&so)
		fmt.Printf("%d\n",soChuKy(so))
	}
}

func soChuKy(so int)  int{
	dem := 0
	if so == 0 {
		return -1
	}
	for so > 1 {
		if so%2 == 0 {
			so = so / 2
		} else {
			so = so*3 + 1
		}
		dem++
	}

	return dem
}
