package main

import "fmt"

func main() {

	for i:=0; i<4; i++{						//File input3.txt có 4 ví dụ nên cho nó chạy 4 lần :v Hoặc là mình để mỗi for {} thì ko phải run lại mà cứ nhập suốt thôi :v
		var n int
		fmt.Scan(&n)

		var number int
		tong:=0
		for i:=0; i<n; i++ {
			fmt.Scan(&number)
			tong = tong + number
		}

		fmt.Printf("%d\n",tong)
	}

}
