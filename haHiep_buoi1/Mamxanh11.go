package main

import "fmt"

func main() {
	var a,b,c,d int;

	fmt.Printf("Nhap 3 so A, B, C phan cach bang dau cach va <=100: ");

	fmt.Scanf("%d %d %d",&a,&b,&c);

	d = (a-b)*c;

	fmt.Printf("%d",d);
}
