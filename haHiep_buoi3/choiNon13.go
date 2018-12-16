package main

import "fmt"

func main() {

	var a,b,c,d int

	fmt.Scanf("%d %d %d %d",&a,&b,&c,&d)

	fmt.Printf("%d\n",soLonThu2(a,b,c,d))

}

func soLonThu2(a,b,c,d int) int {

	if a == max4(a,b,c,d) {
		return max3(b,c,d)
	} else if b == max4(a,b,c,d) {
		return max3(a,c,d)
	} else if c == max4(a,b,c,d) {
		return max3(b,a,d)
	} else {
		return max3(a,b,c)
	}
}

func max4(a,b,c,d int) int {

	if a>=b && a>=c && a>=d {
		return a
	}
	if b>=a && b>=c && b>=d {
		return b
	}
	if c>=b && c>=a && c>=d {
		return c
	}
	return d
}

func max3(a,b,c int) int {

	if a>=b && a>=c {
		return a
	}
	if b>=a && b>=c {
		return b
	}
	return c
}