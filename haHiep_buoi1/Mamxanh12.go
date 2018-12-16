package main

import "fmt"

func main() {
	fmt.Printf("Nhap 2 so A, B phan cach boi dau cach va <=1000: ");
	var a,b,c int;
	fmt.Scanf("%d %d",&a,&b);

	c = a + b + a*b;

	fmt.Printf("%d",c);
}
