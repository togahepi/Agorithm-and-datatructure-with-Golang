package main

import "fmt"

func main() {
	var n int
	for {
		status,_ := fmt.Scanf("%d",&n)
		if n%10 == 0 || status ==0 {
			break;
		}

		fmt.Printf("%d ",n)
	}
}
