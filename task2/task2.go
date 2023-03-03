package main

import "fmt"

func main() {
	var a, b, c, A, B, C uint32

	fmt.Scanln(&a, &b, &c)
	fmt.Scanln(&A, &B, &C)

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
	for x = 0; x <= sum; x++ {
		for y = 0; y <= sum; y++ {
			for z = 0; z <= sum; z++ {
				if (x+y+z == sum) && (x%a1 == 0) && (y%b1 == 0) && (z%c1 == 0) {
					count++
				}
			}
		}
	}
	return count

}
