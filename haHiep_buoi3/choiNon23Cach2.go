package main

import "fmt"

func main() {

	var ngay1,thang1,nam1,ngay2,thang2,nam2 int

	fmt.Scanf("%d %d %d\n%d %d %d\n",&ngay1,&thang1,&nam1,&ngay2,&thang2,&nam2)

	fmt.Printf("%d\n",laAnh(ngay1,thang1,nam1,ngay2,thang2,nam2))
}

func laAnh(ngay1,thang1,nam1,ngay2,thang2,nam2 int) int{

	x1 := ngay1 + thang1*1000 + nam1*1000000
	x2 := ngay2 + thang2*1000 + nam2*1000000

	if x1 > x2 {
		return 2
	}
	if x2 > x1 {
		return 1
	}
	return 0
}
