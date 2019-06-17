package main

import (
	"fmt"
	"resk-projects/src/hackfun.com/resk/infra/algo"
)

func main() {
	//count, amount := int64(5), int64(100*100)
	//for i := int64(0); i < count; i++ {
	//	x := algo.SimpleRand(count-i, amount)
	//	amount -= x
	//	fmt.Print(float64(x)/float64(100), ",")
	//}
	count, amount := int64(5), int64(100)
	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count, amount*100)
		fmt.Print(float64(x)/float64(100), ",")
	}
	fmt.Println()
}
