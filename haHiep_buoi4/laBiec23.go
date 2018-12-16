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
	tong:=0
	for i:=10000; i>0; i = i/10 {
		tong = tong + so/i
		so = so - i*(so/i)
	}
	return tong
}
