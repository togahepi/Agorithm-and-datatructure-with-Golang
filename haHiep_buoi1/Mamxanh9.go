package main

import "fmt"

func main() {
	var gio,phut,giay,soGiay int;

	fmt.Printf("Nhap so giay can tinh trong khoang 0 den 86399: ");
	fmt.Scanf("%d",&soGiay);

	gio = soGiay/3600;
	phut = soGiay%3600/60;
	giay = soGiay%60;

	fmt.Printf("%d:%d:%d",gio,phut,giay);
}