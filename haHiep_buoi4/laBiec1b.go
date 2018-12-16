package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d",&n)

	tong := 0

	for i:=1; i<=n; i++ {
		tong = tong + i
	}

	fmt.Printf("%d\n",tong)
}
