package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var so int
	for i := 0; i < n; i++ {
		fmt.Scan(&so)
		fmt.Printf("%d\n",chuSoLonThu2(so))
	}
}

func chuSoLonThu2(so int)  int{
	if so < 10 {
		return so
	}

	var max1, max2 int

	max1 = so%10		//Lấy ra chữ số đầu tiên
	so = so/10
	max2 = so%10

	if max1 < max2 {	//So sánh 2 số đó để đặt đúng max1 > max2
		x := max1
		max1 = max2
		max2 = x
	}

	if max1 == max2 {	//Trường hợp bằng nhau thì ta đặt max2 = -1 để tiện sắp xếp về sau
		max2 = -1
	}

	for so >= 10 {									//Ta sẽ xét tới các chữ số thứ 2 của số đó sẽ được max1 và max2
		if so%10 > max1 {
			max2 = max1
			max1 = so%10
		} else if so%10 > max2 && so%10 < max1 {
			max2 = so%10
		}
		so = so/10
	}

	if max2 == -1 && so < max1 {					//2 trường hợp đầu là nếu qua cả vòng lặp kia mà max2 = -1 là do tất cả các số đều bằng nhau
		return so
	} else if max2 == -1 {
		return max1
	} else if so > max2 && so < max1{				//2 trường hợp sau là để xét nốt số cuối so với max1, max2 có từ vòng lặp for thôi ko có gì đặc biệt
		return so
	} else if so > max1 {
		return max1
	}

	return max2										//Nếu ko có gì đặc biệt thì ta return max2 thui
}
