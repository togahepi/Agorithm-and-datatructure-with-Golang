package main

import (
	"../toan"
	"fmt"
)

func main() {
	var a,b int
	fmt.Scan(&a,&b)

	uc := uocSoChungLonNhat(a,b)
	boiSoChungNhoNhat := a * b / uc

	fmt.Printf("%d %d",uc,boiSoChungNhoNhat)
}

func uocSoChungLonNhat(a,b int)  int{
	var uc int

	if a > b && a%b == 0 {
		return b
	}

	if b > a && b%a == 0 {
		return a
	}

	for i := toan.Min2so(a,b)/2; i >= 1 ; i--{
		if a%i == 0 && b%i == 0 {
			uc = i
			break;
		}
	}

	return uc
}

