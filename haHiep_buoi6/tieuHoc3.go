package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var soTien int
		fmt.Scan(&soTien)
		menhGia := []int{1,5,10,20,50}
		ketQua := doiTien(soTien,menhGia)
		inKetQua(ketQua,menhGia)
	}
}

func doiTien(soTien int, menhGia []int)  []int{
	ketQua := []int{0,0,0,0,0}
	for i:=len(menhGia)-1; i>=0; i-- {
		ketQua[i] = soTien/menhGia[i]
		soTien = soTien%menhGia[i]
	}
	return ketQua
}

func inKetQua(ketQua,menhGia []int)  {
	for index,_ := range ketQua {
		fmt.Printf("(%d) %d ",menhGia[index],ketQua[index])
	}
}
