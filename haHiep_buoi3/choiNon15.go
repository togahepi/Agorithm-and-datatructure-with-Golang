package main

import "fmt"

func main() {

	var nam int

	fmt.Scanf("%d",&nam)

	fmt.Printf("%d\n",laNamNhuan(nam))
}

func laNamNhuan(nam int) int{

	if nam%400 == 0 {
		return 1
	}

	if nam%100 == 0 {
		return 0
	}

	if nam%4 == 0 {
		return 1
	}

	return 0
}