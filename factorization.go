package factorization

import "math"

// Factorization 对正整数进行因式分解
func Factorization(n uint64) []uint64 {

	factorSlice := make([]uint64, 0)
	// 平方根之后就不可能了因此就没必要继续了
	factorEnd := uint64(math.Ceil(math.Sqrt(float64(n))))
	for n > 0 {

		// 启用素数表加速，跳着往后寻找约数，性能有明显上升
		// 仅当n小于这个数字时才启用素数表来加速，否则就还是常规的递增迭代，这是因为素数表的数据规模是有限的，要能够兼容数值超出素数表边界的情况
		factor, isFindSuccess, isCanFind := DefaultPrimeNumberTable.FindFirstFactor(n)

		if !isCanFind {
			// 普通的递增迭代判断，会有一些无效的判断，略有浪费
			// 采取效率慢点的方式，O(n^2)
			factor, isFindSuccess = findFirstFactorByEnd(factorEnd, n)
		}

		if isFindSuccess {
			factorSlice = append(factorSlice, factor)
			n /= factor
		} else {
			// 到达了边界值仍未寻找到，n已经没办法再约掉了
			break
		}

	}

	// 如果有未分解完的，则把它也追加到结果中
	if n > 1 {
		factorSlice = append(factorSlice, n)
	}

	return factorSlice
}

// 暴力搜索第一个因子
func findFirstFactorByEnd(factorEnd, n uint64) (firstFactor uint64, isFindSuccess bool) {
	for i := uint64(2); i <= factorEnd; i++ {
		if n%i == 0 {
			return i, true
		}
	}
	return 0, false
}
