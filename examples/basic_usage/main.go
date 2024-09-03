package main

import (
	"fmt"
	"github.com/cryptography-research-lab/go-factorization"
	"time"
)

func main() {
	beginTime := time.Now()
	// 参数是要进行因式分解的数，返回值是进行因式分解的结果
	factorization := factorization.Factorization(uint64(12343543765673436))
	cost := time.Now().Sub(beginTime)
	fmt.Println(factorization) // Output: [191749 6421127]
	fmt.Println(cost)          // Output: 132.167µs
}
