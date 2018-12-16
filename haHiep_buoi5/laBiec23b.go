package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var so int
	for i := 0; i < n; i++ {
		fmt.Scan(&so)
		fmt.Printf("%d\n",chuSoLonNhat(so))
	}

}

func chuSoLonNhat(n int)  int{
	chuSoLonNhat := n%10
	n = n/10
	for n > 0 {
		if n%10 > chuSoLonNhat {
			chuSoLonNhat = n%10
		}
		n = n/10
	}

	return chuSoLonNhat
}