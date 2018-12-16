package lib

import "fmt"

func InputSlice(size int) []int{
	var slice []int

	for i:=0; i<size; i++ {
		var value int
		fmt.Scan(&value)

		slice = append(slice,value)
	}

	return slice;
}

func OutPutSlice(slice []int)  {
	for i:=0; i<len(slice); i++ {
		fmt.Printf("%d ",slice[i])
	}
	fmt.Println()
}
