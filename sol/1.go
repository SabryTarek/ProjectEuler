////////////////////////////////////////////////////////////////////////////////////////////////////
// #Problem:	1. [Multiples of 3 or 5](https://projecteuler.net/problem=1)			  //
// #Author:	SabryTarek									  //
// #Date:	11/2/2022									  //
////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func mul_3_or_5(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	input := 1000
	output := mul_3_or_5(input)

	fmt.Println(output)
}
