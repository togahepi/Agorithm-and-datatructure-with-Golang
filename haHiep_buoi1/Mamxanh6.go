package main

import (
	"fmt"
)

func main() {
	var r int;

	fmt.Printf("Nhap ban kinh duong tron la 1 so nguyen duong: ");

	fmt.Scanf("%d",&r);

	var s float32;

	s = 3.14*float32(r*r);

	fmt.Printf("Dien tich hinh tron co ban kinh %d la: %f",r,s);

}
