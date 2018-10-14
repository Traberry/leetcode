package main

import (
	"fmt"
	"math"
)

func find1(a []int) int {
	m := make(map[float64]int)
	for i := 0; i <= len(a); i++ {
		p := math.Abs( float64(a[i]) )
		m[p] = i
	}
	return len(m)
}

func main() {
	test := []float64{-2, -1, 0, 1, 2}
	fmt.Println(find2(test))
}

func find2(a []float64) int {
	num1 := a[0]
	num2 := a[len(a) - 1]

	count := 0
	for i, j := 0, len(a)-1; i <= j; {
		num1 = math.Abs(a[i])
		num2 = math.Abs(a[j])

		if num1 > num2 {
			count++
			for i <= j && math.Abs(a[i]) == num1 {
				i++
			}
		}else if num1 < num2 {
			count++
			for i <= j && math.Abs(a[j]) == num2 {
				j--
			}
		}else {
			count++
			for i <= j && math.Abs(a[i]) == num1 {
				i++
			}
			for i <= j && math.Abs(a[j]) == num2 {
				j--
			}
		}
	}

	return count
}
