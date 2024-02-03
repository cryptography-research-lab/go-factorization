package factorization

import (
	"github.com/cryptography-research-lab/go-factorization/data"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------

type PrimeNumberTable struct {
	primeNumberTable []uint64
	max              uint64
}

func NewPrimeNumberTable() *PrimeNumberTable {
	return &PrimeNumberTable{
		primeNumberTable: make([]uint64, 0),
		max:              0,
	}
}

func (x *PrimeNumberTable) Init(lines string) error {
	x.primeNumberTable = make([]uint64, 0)
	split := strings.Split(lines, "\n")
	for _, line := range split {

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		atoi, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		x.primeNumberTable = append(x.primeNumberTable, uint64(atoi))
	}

	if len(x.primeNumberTable) != 0 {
		n := x.primeNumberTable[len(x.primeNumberTable)-1]
		x.max = n * n
	}
	return nil
}

// findFirstFactorByEnd 找到第一个因子
func (x *PrimeNumberTable) FindFirstFactor(n uint64) (firstFactor uint64, isFindSuccess bool, isCanFind bool) {

	if n > x.max {
		return 0, false, false
	}

	// 跳着往前寻找约数，性能有明显上升
	for _, i := range x.primeNumberTable {
		if n%i == 0 {
			return i, true, true
		}
	}

	return 0, false, true
}

// ---------------------------------------------------------------------------------------------------------------------

var DefaultPrimeNumberTable *PrimeNumberTable

func init() {
	DefaultPrimeNumberTable = NewPrimeNumberTable()
	err := DefaultPrimeNumberTable.Init(data.PrimeNumberTable)
	if err != nil {
		panic(err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
