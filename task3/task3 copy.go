package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	// N = 2
	a, b := NOKofSum(N)
	fmt.Printf("%d %d\n", a, b)

	// for i := 1; i < 150; i++ {
	// 	a, b = NOKofSum(i)
	// 	fmt.Printf("%s %d %s %d %d\n", "N ==", i, "| A, B ==", a, b)

	// }
}

func NOKofSum(N int) (int, int) {
	NOK := 1000000000000000
	var A, B int
	for a := 1; a < N/2+1; a++ {
		b := N - a

		// fmt.Println("НОК для ", a, b)
		// fmt.Println(findNOK(a, b))

		if findNOK(a, b) < NOK {
			A, B = a, b
			NOK = findNOK(A, B)
		}
	}

	// fmt.Println("A, B ==", A, " ", B)
	return A, B
}

func findNOK(a int, b int) int {

	var deliteliA, deliteliB []int
	res := 1

	deliteliA = prostieDeliteli(a)
	deliteliB = prostieDeliteli(b)

	// fmt.Println("Простые делители для ", a, b)
	// fmt.Println(deliteliA)
	// fmt.Println(deliteliB)

	var mensheDeliteley, bolsheDeliteley *[]int

	if len(deliteliA) >= len(deliteliB) {
		mensheDeliteley = &deliteliB
		bolsheDeliteley = &deliteliA
	} else {
		mensheDeliteley = &deliteliA
		bolsheDeliteley = &deliteliB
	}

	m := make(map[int]uint8)

	for _, v := range *mensheDeliteley {
		m[v]++
	}

	for _, v := range *bolsheDeliteley {
		res *= v
		if m[v] != 0 {
			m[v]--
		}
	}

	for i, v := range m {
		if v != 0 {
			res *= i * int(v)
		}
	}

	return res
}

func prostieDeliteli(N int) []int {
	var deliteli []int
	for naimenshiyDelitel(N) != N {
		deliteli = append(deliteli, naimenshiyDelitel(N))
		N /= naimenshiyDelitel(N)
	}
	deliteli = append(deliteli, naimenshiyDelitel(N))
	return deliteli
}

func naimenshiyDelitel(N int) int {
	if N == 1 {
		return 1
	}
	res := 0
	for i := 2; i <= N; i++ {
		if N%i == 0 {
			res = i
			break
		}
	}
	return res
}
