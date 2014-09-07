package main

import "fmt"

func readX (N int) {
	if N != 0 {
		var X int
		fmt.Scan(&X)
		res := readY(X)
		fmt.Println(res)
		readX(N-1)
	}
}

func readY (X int) int {
	if X != 0 {
		var Y int
		var temp int
		fmt.Scan(&Y)
		if Y > 0 {
			temp = Y * Y
		} else {
			temp = 0
		}
		return temp + readY(X-1)
	} 
	return 0
}

func main() {
    var N int
	fmt.Scan(&N)
	readX(N)
}