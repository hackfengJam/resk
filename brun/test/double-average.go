package main

import (
	"fmt"
	"resk-projects/src/hackfun.com/resk/infra/algo"
)

func main() {
	var (
		sum    int64
		remain int64
	)

	count, amount := int64(10), int64(100*100)
	sum = 0
	remain = amount
	for i := int64(0); i < count; i++ {
		x := algo.DoubleAverage(count-i, remain)
		remain -= x
		fmt.Print(x, ", ")
		sum += x
	}
	fmt.Println()
	fmt.Println(sum)
}
