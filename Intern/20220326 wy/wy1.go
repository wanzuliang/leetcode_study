func main() {
	// a, b, x, y := 0, 0, 0, 0
	str := ""
	for {
		// n, _ := fmt.Scan(&a, &b, &x, &y)
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		} else {
			// test1(a, b, x, y)
			test1(str)
		}
	}
}

func test1(a, b, x, y int) {
	// n, m := 0, 0
	res := 20
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 40; j++ {
			for k := 0; k <= 40; k++ {
				if a <= j*x+i*y && b <= k*x+i*y {
					if i+j+k < res {
						res = i + j + k
					}
				}
			}
		}
	}
	fmt.Println(res)
}
