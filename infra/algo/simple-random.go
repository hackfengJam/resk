package algo

import (
	"math/rand"
	"time"
)

// 最小一分
const min = int64(1)

func randInt63(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Int63n(max-min) + min
}

// 简单随机算法
// 红包数量，红包金额
// 金额单位为分，1元=100分 钱
func SimpleRand(count, amount int64) int64 {
	// 当红包金额剩余一个的时候，直接返回剩余金额
	if count == 1 {
		return amount
	}
	// 计算最大可调度金额
	max := amount - min*count
	x := randInt63(min, max)
	return x
}
