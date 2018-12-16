package main

import (
	"fmt"
	"math"
)

func main() {

	var a,b,c int

	fmt.Scanf("%d %d %d",&a,&b,&c)

	if laTamGiac(a,b,c) {
		fmt.Printf("%f\n",dienTichTamGiac(a,b,c))
	} else {
		fmt.Printf("Khong phai la tam giac\n")
	}

}

func laTamGiac(a,b,c int) bool {

	if a+b>c && a+c>b && b+c>a {
		return true
	}

	return false
}

func dienTichTamGiac(a,b,c int) float64 {

	p:= (a+b+c)/2

	s:=math.Sqrt(float64(p)*(float64(p)-float64(a))*(float64(p)-float64(b))*(float64(p)-float64(c)))

	return s

}