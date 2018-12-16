package main

import (
	"fmt"
	"math"
)

func main() {
	var r int;
	fmt.Printf("Nhap ban kinh hinh cau: ");
	fmt.Scanf("%d",&r);

	var s float32;
	var v float32;
	s=4.0*math.Pi*float32(r*r);
	v=(4.0/3)*math.Pi*float32(r*r*r);

	fmt.Printf("Dien tich hinh cau la: %f\n",s);
	fmt.Printf("The tich hinh cau la: %f",v);
}
