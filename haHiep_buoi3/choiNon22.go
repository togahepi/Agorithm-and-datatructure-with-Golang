package main

import "fmt"

func main() {

	var a,b,c int

	fmt.Scanf("%d %d %d",&a,&b,&c)

	fmt.Printf("%s\n",inHieu2SoBangSoConLai(a,b,c))
}

func inHieu2SoBangSoConLai(a,b,c int) string {
	if a - b == c {
		inHieu(a,b,c)
	}
	if b - a == c {
		inHieu(b,a,c)
	}
	if a - c == b {
		inHieu(a,c,b)
	}
	if c - a == b {
		inHieu(c,a,b)
	}
	if b - c == a {
		inHieu(b,c,a)
	}
	if c - b == a {
		inHieu(c,b,a)
	}
	return ""
}

func inHieu(a,b,c int) string {
	fmt.Printf("%d - %d = %d\n",a,b,c)
	return ""
}
