package main

import (
	"agoGolang/lib"
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var size int
		fmt.Scan(&size)

		slice := lib.InputSlice(size)

		fmt.Printf("%f\n",trungBinhCong(slice))
		//trungBinhNhan(slice)
		fmt.Printf("%f\n",trungBinhNhan(slice))
	}
}

func trungBinhCong(slice []int)  float64{
	var trungBinhCong float64

	tong:=0

	for i:=0; i<len(slice); i++ {
		tong = tong + slice[i]
	}

	trungBinhCong = float64(tong) / float64(len(slice))

	return trungBinhCong
}

func trungBinhNhan(slice []int) float64{
	var trungBinhNhan float64
	tich := 1

	for i:=0; i<len(slice); i++ {
		tich = tich*slice[i]
	}

	trungBinhNhan = math.Pow(float64(tich),1.0/float64(len(slice)))

	return trungBinhNhan
}
