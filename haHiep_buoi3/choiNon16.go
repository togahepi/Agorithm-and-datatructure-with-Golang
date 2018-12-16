package main

import (
	"fmt"
)

func main() {

	var a,b,c int

	fmt.Scanf("%d %d %d",&a,&b,&c)

	fmt.Printf("%d\n",soLonNhat(a,b,c))

	fmt.Printf("%d\n",soBeNhat(a,b,c))
}

func max(a,b int) int {

	if a>b {
		return a
	}
	return b
}

func min(a,b int) int {

	if a<b {
		return a
	}
	return b
}

func soBeNhat(a,b,c int) int {

	soMin := min(a,b)

	soMin = min(soMin,c)

	return soMin
}

func soLonNhat(a,b,c int) int {

	soMax := max(a,b)

	soMax = max(soMax,c)

	return soMax
}
