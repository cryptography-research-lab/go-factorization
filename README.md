#  Golang因式分解库

# 一、这是什么？

这是一个做因式分解的库，优势：

- 采用了素数表来加速，效率极高！

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/go-factorization
```

# 三、API Example

## 3.1 基本使用

```golang
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
```

## 3.2 自定义素数表

```golang
package main

import (
	"fmt"
	"github.com/cryptography-research-lab/go-factorization"
	"github.com/cryptography-research-lab/go-factorization/data"
	"time"
)

func main() {

	// 自定义素数表
	table := factorization.NewPrimeNumberTable()
	// 在这个地方可以使用自定义的素数表来初始化
	// 格式是每行一个素数，按照从小到大排好序
	err := table.Init(data.PrimeNumberTable)
	if err != nil {
		panic(err)
	}
	factorization.DefaultPrimeNumberTable = table

	beginTime := time.Now()
	// 参数是要进行因式分解的数，返回值是进行因式分解的结果
	factorization := factorization.Factorization(uint64(12343543765673436))
	cost := time.Now().Sub(beginTime)
	fmt.Println(factorization) // Output: [191749 6421127]
	fmt.Println(cost)          // Output: 132.167µs
}
```







