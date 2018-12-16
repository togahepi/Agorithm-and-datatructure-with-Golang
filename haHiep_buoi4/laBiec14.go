package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0{
		fmt.Printf("%d\n",(n/2)-1)
	} else {
		fmt.Printf("%d\n",(n/2))
	}

	for i:=n-1; i>=(n/2+1); i-- {
		fmt.Printf("%d %d\n",i,(n-i))
	}
}
