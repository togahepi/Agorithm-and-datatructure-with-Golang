package main

import (
	"../toan"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a,b int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d",&a,&b)
		tong,so := tongVaSoLuongCacSoNguyenTo(a,b)
		fmt.Printf("%d %d\n",tong,so)
	}
}

func tongVaSoLuongCacSoNguyenTo(a,b int)  (int,int) {
	tong := 0
	dem := 0
	for i := a; i <= b; i++ {
		if toan.LaSoNguyenTo(i) {
			tong = tong + i
			dem++
		}
	}
	return tong,dem
}