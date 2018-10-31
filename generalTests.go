// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import (
	"fmt"
	"math"
)

func main() {

	pack := 32
	total := 135

	loops := math.Ceil(float64(total) / float64(pack))
	fmt.Println(loops)
	bufy := make([]byte, total)
	sent := 0
	for i := 0; i < int(loops); i++ {
		sent += pack
		if sent > total {
			rest := pack - (sent - total)
			sent -= pack
			sent += rest
			pack = rest
		}
		base := sent - pack
		top := sent
		fmt.Printf("%d:%d\n", base, top)

		newArr := bufy[base:top]
		_ = newArr

	}

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
