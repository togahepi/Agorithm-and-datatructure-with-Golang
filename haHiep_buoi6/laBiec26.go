package main

import "fmt"

func main() {
	var slice []int;
	var value int;

	dem := make(map[int]int)

	for {
		status, _ := fmt.Scan(&value)
		if (status == 0) {
			break;
		}
		slice = append(slice,value)
	}

	for _,value := range slice {
		dem[value]++
	}

	fmt.Println(dem)
}
