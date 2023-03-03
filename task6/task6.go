package main

import "fmt"

func main() {
	var q, cur, maxXOR uint32
	S := make(map[uint32]bool)
	var res []uint32
	fmt.Scanln(&q)

	for i := 0; uint32(i) < q; i++ {
		fmt.Scanln(&cur)
		if S[cur] {
			res = append(res, maxXOR)
			continue
		}
		S[cur] = true
		for m := range S {
			for n := range S {
				if maxXOR < m^n {
					maxXOR = m ^ n
				}
			}
		}
		res = append(res, maxXOR)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
