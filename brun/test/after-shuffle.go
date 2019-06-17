package main

import (
	"fmt"
	"resk-projects/src/hackfun.com/resk/infra/algo"
)

func main() {
	var (
		rev []int64
		sum int64
	)
	sum = 0
	rev = algo.AfterShuffle(int64(10), int64(10000))
	fmt.Println(rev)
	for _, item := range rev {
		sum += item
	}
	fmt.Println(sum)
}
