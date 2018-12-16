package main

import "fmt"

func main() {

	var a,b,c,d int

	fmt.Scanf("%d %d %d %d",&a,&b,&c,&d)

	fmt.Printf("%d\n",soLonThu2(a,b,c,d))
}

func soLonThu2(a,b,c,d int) int {

	max1 := a
	max2 := b

	if c > a {
		max2 = max1
		max1 = c
	} else if c < a && c > b {
		max2 = c
	}

	if d > a {
		max2 = max1
		max1 = d
	} else if d < a && d > b {
		max2 = d
	}

	return max2
}
