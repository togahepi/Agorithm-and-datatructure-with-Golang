package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i:=6; i<n; i++ {
		if laSoHoanHao(i) {
			fmt.Printf("%d ",i)
		}
	}
}

func laSoHoanHao(so int)  bool{
	tongUoc := 0
	for i := so/2; i >= 1; i-- {
		if so%i == 0 {
			tongUoc = tongUoc + i
		}

		if tongUoc > so {
			return false
		}
	}

	if tongUoc < so {
		return false
	}

	return true
}

