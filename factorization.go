package factorization

import "math"

// Factorization 对正整数进行因式分解
func Factorization(n int) []int {
	if n <= 0 {
		return nil
	}
	// 先采取效率慢点的方式，O(n^2)，后面再看看怎么优化
	factorSlice := make([]int, 0)
	// 平方根之后就不可能了因此就没必要继续了
	end := int(math.Ceil(math.Sqrt(float64(n))))
	for n > 0 {
		isFindSuccess := false
		for i := 2; i <= end; i++ {
			if n%i == 0 {
				isFindSuccess = true
				factorSlice = append(factorSlice, i)
				n /= i
				break
			}
		}
		if !isFindSuccess {
			if n != 1 || len(factorSlice) == 0 {
				factorSlice = append(factorSlice, n)
			}
			break
		}
	}
	return factorSlice
}
