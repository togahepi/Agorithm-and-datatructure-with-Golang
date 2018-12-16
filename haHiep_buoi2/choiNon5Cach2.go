package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	max:=a;
	if b>max {
		max = b
	}
	if c>max {
		max = c
	}
	fmt.Printf("%d la so lon nhat\n",max)
}
