package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d",&n)

	var so int
	for i:=0; i<n; i++ {
		fmt.Scanf("%d",&so)
		if laSoNguyenTo(so) {
			fmt.Printf("%d la so nguyen to\n",so)
		} else {
			fmt.Printf("%d khong la so nguyen to\n",so)
		}
	}
}

func laSoNguyenTo(so int) bool {
	if so == 0 || so == 1 {
		return false
	}
	for i:=2; i<=(so/2); i++ {
		if so%i == 0 {
			return false
		}
	}
	return true
}
