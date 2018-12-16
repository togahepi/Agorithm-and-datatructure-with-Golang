package main

import "fmt"

func main() {

	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	if laTamGiacVuong(a,b,c) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

func laTamGiacVuong(a,b,c int) bool{

	A:=a*a
	B:=b*b
	C:=c*c

	if laTamGiac(a,b,c) && ( A+B == C || A+C == B || B+C==A) {
		return true
	}

	return false
}

func laTamGiac(a,b,c int) bool{

	if a+b > c && a+c > b && b+c > a {
		return true
	}

	return false
}


