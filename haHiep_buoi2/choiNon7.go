package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	if a+b>c && a+c>b && b+c>a {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
