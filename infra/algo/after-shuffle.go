package algo

import "math/rand"

// 后洗牌算法
func AfterShuffle(count, amount int64) [] int64 {
	inds := make([]int64, 0)
	// 计算最大可调用金额
	// 这种方式产生的数据没有1了。准确的做法应该是在 AfterShuffle 计算最大可调度金额时不需要减1，因为 SimpleRand 中已经减了1
	//max := amount - min*count
	max := amount
	remain := max
	// 随机生成初级红包序列
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		inds = append(inds, x)
	}
	// 洗牌，洗初级红包序列
	rand.Shuffle(len(inds), func(i, j int) {
		inds[i], inds[j] = inds[j], inds[i]
	})
	return inds
}
