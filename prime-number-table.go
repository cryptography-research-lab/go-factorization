package factorization

import (
	"github.com/cryptography-research-lab/go-factorization/data"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------

// PrimeNumberTable 素数表，用于封装查表相关的一些操作
type PrimeNumberTable struct {

	// 底层的素数表
	primeNumberTable []uint64

	// 当前的素数表中最大的数字，用于在使用判断数字是否超出了素数表的范畴
	max uint64
}

// NewPrimeNumberTable 创建一张素数表
func NewPrimeNumberTable() *PrimeNumberTable {
	return &PrimeNumberTable{
		primeNumberTable: make([]uint64, 0),
		max:              0,
	}
}

// Init 使用给定的素数表初始化，lines是每行一个素数，并且是从小到大排好序的
// 这个方法可以重复调用，每次重复调用的时候会把之后的素数表完全刷新掉
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

// FindFirstFactor 找到n的最小的一个因子
// return:
//
//	firstFactor: 找到的n的第一个因子
//	isFindSuccess: 是否找到了素数，仅在此参数为true的情况下firstFactor的值才有效
//	isCanFind: 当前素数表是否能够起到加速的作用，如果不能起到加速作用的话直接返回false，仅在此参数为true的前提下firstFactor和isFindSuccess才有效
func (x *PrimeNumberTable) FindFirstFactor(n uint64) (firstFactor uint64, isFindSuccess bool, isCanFind bool) {

	if n > x.max {
		return 0, false, false
	}

	// 按照素数表跳着往后寻找约数，性能有明显上升，越往后性能加速越快，因为省略掉了无效数字的取余运算
	for _, i := range x.primeNumberTable {
		if n%i == 0 {
			return i, true, true
		}
	}

	return 0, false, true
}

// ---------------------------------------------------------------------------------------------------------------------

// DefaultPrimeNumberTable 默认使用的素数表，如果觉得内置的素数表太小的话，可以使用自定义的素数表替换掉默认的素数表
var DefaultPrimeNumberTable *PrimeNumberTable

// 初始化默认的素数表
func init() {
	DefaultPrimeNumberTable = NewPrimeNumberTable()
	err := DefaultPrimeNumberTable.Init(data.PrimeNumberTable)
	if err != nil {
		panic(err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
