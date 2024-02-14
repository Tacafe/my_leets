package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	pow2 := 0
	y := x
	for y != 1 {
		y = int(y / 2)
		pow2 += 1
	}

	ans := int(math.Pow(2, float64(int(pow2/2))))

	for {
		test := int(math.Pow(float64(ans), 2.0))
		if test == x {
			return ans
		} else if test > x {
			return ans - 1
		} else {
			ans++
		}
	}
}

func main() {
	x := 10
	fmt.Println(mySqrt(x))

}
