package main

import "fmt"

func main() {
	var gio,luongGio,tienLuong int
	fmt.Scanf("%d %d",&gio,&luongGio)

	if gio<=40 {
		tienLuong = gio*luongGio
	} else {
		tienLuong = 40*luongGio + (gio-40)*2*luongGio
	}

	fmt.Printf("%d\n",tienLuong)
}
