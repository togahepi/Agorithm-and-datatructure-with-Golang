package main

import "fmt"

func main() {
	var a int;
	fmt.Printf("Nhap so nguyen bat ky: ");
	fmt.Scanf("%d",&a);

	var b, c int;
	b = a*a;
	c = a*a*a;

	fmt.Printf("Binh phuong cua %d la: %d\n", a, b);
	fmt.Printf("Lap phuong cua %d la: %d", a, c);
}
