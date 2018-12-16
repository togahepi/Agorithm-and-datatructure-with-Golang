package toan

func LaSoNguyenTo(so int) bool {
	if so == 0 || so == 1 {
		return false
	}
	for i:=2; i<=(so/2); i++ {
		if so%i == 0 {
			return false
		}
	}
	return true
}
