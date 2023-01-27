package factorization

import (
	"fmt"
	"testing"
	"time"
)

func TestFactorization(t *testing.T) {
	beginTime := time.Now()
	factorization := Factorization(uint64(1231244681123))
	cost := time.Now().Sub(beginTime)
	fmt.Println(factorization) // Output: [191749 6421127]
	fmt.Println(cost)          // Output: 9.1188ms
}
