package main

import "fmt"

func main() {

	var ngay1,thang1,nam1,ngay2,thang2,nam2 int

	fmt.Scanf("%d %d %d\n%d %d %d\n",&ngay1,&thang1,&nam1,&ngay2,&thang2,&nam2)

	fmt.Printf("%d\n",laAnh(ngay1,thang1,nam1,ngay2,thang2,nam2))
}

func laAnh(ngay1,thang1,nam1,ngay2,thang2,nam2 int) int{

	if nam1 > nam2 {
		return 2
	}
	if nam2 > nam1 {
		return 1
	}
	if thang1 > thang2 {
		return 2
	}
	if thang2 > thang1 {
		return 1
	}
	if ngay1 > ngay2 {
		return 2
	}
	if ngay2 > ngay1 {
		return 1
	}
	return 0
}

