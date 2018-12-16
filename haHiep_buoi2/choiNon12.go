package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	var delta,x1,x2 float64

	delta = float64(b*b - 4*a*c)

	if delta < 0 {
		fmt.Printf("Phuong trinh da cho vo nghiem\n")
	} else if delta==0 {
		x1 = float64(-b)/float64(2*a)
		fmt.Printf("Phuong trinh da cho co nghiem kep x1 = x2 = %f64\n",x1)
	} else {
		x1 = (float64(-b) + math.Sqrt(delta)) / float64(2*a)
		x2 = (float64(-b) - math.Sqrt(delta)) / float64(2*a)
		fmt.Printf("Phuong trinh da cho co 2 phan biet x1 = %f64 và x2 = %f64\n",x1,x2)
	}
}
