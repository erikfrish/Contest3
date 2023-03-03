package main

import "fmt"

func main() {
	var a, b, c, A, B, C uint32

	// fmt.Scanln(&a, &b, &c)
	// fmt.Scanln(&A, &B, &C)

	// a, b, c = 1, 1, 1
	// A, B, C = 1, 0, 2 //10

	a, b, c = 1, 2, 3
	A, B, C = 3, 5, 4 //28

	fmt.Println(exchangevariants(a, b, c, A, B, C))

}

func exchangevariants(a, b, c, A, B, C uint32) uint32 {
	var a1, b1, c1, coef uint32

	coef = a * b * c
	a1 = coef / a
	b1 = coef / b
	c1 = coef / c

	var x, y, z, count uint32
	sum := a1*A + b1*B + c1*C
	// var arr [][3]uint32
	for x = 0; x <= sum; x++ {
		for y = 0; y <= sum; y++ {
			for z = 0; z <= sum; z++ {
				if (x+y+z == sum) && (x%a1 == 0) && (y%b1 == 0) && (z%c1 == 0) {
					count++
					// arr = append(arr, [3]uint32{x / a1, y / b1, z / c1})
				}
			}
		}
	}
	// fmt.Println(arr)
	return count

}
