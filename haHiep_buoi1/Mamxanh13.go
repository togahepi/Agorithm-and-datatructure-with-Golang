package main

import "fmt"

func main() {
	var gio,phut,giay,tongGiay int;

	fmt.Printf("Nhap thoi gian dang Gio:Phut:Giay \n");

	fmt.Scanf("%d:%d:%d",&gio,&phut,&giay);

	tongGiay = gio*3600 + phut*60 + giay;

	fmt.Printf("Nhap so X <=1000 \n");

	var x,gio2,phut2,giay2 int;

	fmt.Scanf("%d",&x);

	tongGiay = tongGiay + x;

	gio2=tongGiay/3600%24;
	phut2=tongGiay%3600/60;
	giay2=tongGiay%60;

	fmt.Printf("Sau %d giay ke thu thoi diem %d:%d:%d thi thoi gian la: %d:%d:%d.",x,gio,phut,giay,gio2,phut2,giay2);
}
