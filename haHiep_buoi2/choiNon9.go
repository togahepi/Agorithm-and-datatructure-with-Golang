package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d",&a)

	if a<0 {
		fmt.Printf("%d la so am\n",a)
	} else if a%2 == 1 {
		fmt.Printf("%d la so le\n",a)
	} else {
		fmt.Printf("%d la so chan\n",a)
	}
}
