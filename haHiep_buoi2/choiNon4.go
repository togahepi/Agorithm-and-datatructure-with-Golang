package main

import "fmt"

func main() {
	var diem float32
	fmt.Scanf("%f",&diem)

	if diem>=9 {
		fmt.Printf("Xuat sac\n")
	} else if diem>=8 {
		fmt.Printf("Gioi\n")
	} else if diem>=7 {
		fmt.Printf("Kha\n")
	} else if diem>=6 {
		fmt.Printf("TBKha\n")
	} else if diem>=5 {
		fmt.Printf("Tbinh\n")
	} else {
		fmt.Printf("Yeu\n")
	}
}
