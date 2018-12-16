package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var so int
	for i:=0; i<n; i++ {
		fmt.Scan(&so)
		fmt.Printf("%d\n",tongCacChuSo(so))
	}
}

func tongCacChuSo(so int) int{
	tong := 0
	for so > 0 {
		tong = tong + (so%10)
		so = so/10
	}
	return tong
}
