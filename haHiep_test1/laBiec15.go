package main

import "fmt"

func main() {
	for i:=0; i<10; i++ {
		var n int
		fmt.Scan(&n)

		doPrint(" ",n-1)
		doPrint("*\n",1)

		soDauCach1 := n-2
		soDauCach2 := 1

		for i:=0; i<(n-2); i++ {
			doPrint(" ",soDauCach1)
			doPrint("*",1)
			doPrint(" ",soDauCach2)
			doPrint("*\n",1)
			soDauCach1--
			soDauCach2 = soDauCach2 + 2
		}

		doPrint("*",2*(n-1)+1)
		fmt.Printf("\n")
	}

}

func doPrint(c string, n int) {
	for i:=0; i<n; i++ {
		fmt.Printf("%s",c)
	}
}