package main

import "fmt"

func main() {
	var n,m int
	var user, password string;
	fmt.Scan(&n)

	db := make(map[string]string)

	for i:=0; i<n; i++ {
		fmt.Scan(&user,&password)
		db[user] = password
	}

	fmt.Scan(&m)

	for i:=0; i<m; i++ {
		fmt.Scan(&user,&password)

		if password == db[user] {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
