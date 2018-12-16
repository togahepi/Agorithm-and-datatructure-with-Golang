package main

import "fmt"

func main() {
	var a,b int
	var nghiem float32
	fmt.Scanf("%d %d",&a,&b)

	if a != 0 {
		nghiem=float32(-b)/float32(a)
		fmt.Printf("Phuong trinh da cho co 1 nghiem x = %f",nghiem)
	} else if b==0 {
		fmt.Printf("Phuong trinh da cho vo so nghiem\n")
	} else {
		fmt.Printf("Phuong trinh da cho vo nghiem\n")
	}
}
